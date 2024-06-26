
# Monitoring and tracing

The monitoring stack for this project is based on the following components:

```
OpenTelemetry
Prometheus
Fluentbit
Grafana Loki
Grafana
Jaeger
```

Here are the instructions for installing the monitoring stack.  

# Install the basic components in your Kubernetes cluster

Here's how to quickly install Flux on a Kuberenetes cluster: 

```
flux check --pre
flux install
flux check
```

Install cert-manager to add certificates and certificate issuers as resource types in Kubernetes clusters. This tool simplifies the process of obtaining, renewing, and using certificates.  

```
flux create source helm cert-manager --url https://charts.jetstack.io
flux create helmrelease cert-manager \
  --chart cert-manager \
  --source HelmRepository/cert-manager.flux-system \
  --release-name cert-manager \
  --target-namespace cert-manager \
  --create-target-namespace \
  --values monitoring/cert-manager/values.yaml
```

Create a namespace for deploying the monitoring stack:  

```
kubectl create namespace monitoring
```

## Monitoring

Install the OpenTelemetry collector first:

```
flux create source helm opentelemetry --url https://open-telemetry.github.io/opentelemetry-helm-charts
flux create helmrelease otel-collector \
  --chart opentelemetry-collector \
  --source HelmRepository/opentelemetry.flux-system \
  --release-name otel-collector \
  --target-namespace monitoring \
  --create-target-namespace \
  --values monitoring/otel-collector/values.yaml
```

Install Prometheus:
```
flux create source helm prometheus --url https://prometheus-community.github.io/helm-charts
flux create helmrelease prometheus \
  --chart prometheus \
  --source HelmRepository/prometheus.flux-system \
  --release-name prometheus \
  --target-namespace monitoring \
  --create-target-namespace \
  --values monitoring/prometheus/values.yaml
```

Setup access to Prometheus UI
```
export PR_POD_NAME=$(kubectl get pods --namespace monitoring -l "app.kubernetes.io/name=prometheus,app.kubernetes.io/instance=prometheus" -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace monitoring port-forward $PR_POD_NAME 9090
```

You can then navigate to http://127.0.0.1:3000 to access the Prometheus UI:
```
http://127.0.0.1:9090
```

Then install the rest of the monitoring tools:

```
helm repo add fluent https://fluent.github.io/helm-charts
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm install fluent-bit fluent/fluent-bit -n monitoring --values fluent-bit/values.yaml 
helm install grafana grafana/grafana -n monitoring --values grafana/values.yaml
helm install loki grafana/loki -n monitoring --values  loki/values.yaml
```

Get admin password for Grafana access:
```
kubectl get secret --namespace monitoring grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
```

Set up access to the Grafana UI: 

```
export GR_POD_NAME=$(kubectl get pods --namespace monitoring -l "app.kubernetes.io/name=grafana,app.kubernetes.io/instance=grafana" -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace monitoring port-forward $GR_POD_NAME 3000
```

You can then navigate to http://127.0.0.1:3000 to access the Grafana UI.  

Demo screenshots of the Grafana UI with dashboards: [grafana1](/images/grafana1.png), [grafana2](/images/grafana2.png) and [grafana3](/images/grafana3.png).

## Tracing

Jaeger is an open-source distributed tracing and observability platform that is very popular among developers.

Install Jaeger on a Kubernetes cluster:

```
helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
helm repo update

helm install jaeger-operator jaegertracing/jaeger-operator -n monitoring --values jaeger-operator/values.yaml
helm install jaeger jaegertracing/jaeger  -n monitoring --values jaeger/values.yaml
```

Set up access to the Jaeger UI: 

```
export JG_POD_NAME=$(kubectl --namespace monitoring get pods  -l "app.kubernetes.io/instance=jaeger,app.kubernetes.io/component=query" -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace monitoring port-forward  $JG_POD_NAME 16686:16686
```

You can then navigate to http://127.0.0.1:16686 to access the Jaeger UI.

Tracing requires adding special functions and calls to an application. 
Here is an example of how to do it in a Golang app: https://github.com/open-telemetry/opentelemetry-go/blob/main/example/otel-collector/main.go.
For the monitoring stack described here, the endpoint for tracing is `otel-collector-opentelemetry-collector.monitoring.svc.cluster.local:4317`.

Views of the Jaeger UI with example tracing data from a chatbot app: [tracing1](/images/tracing1.png), [tracing2](/images/tracing2.png) and [tracing3](/images/tracing3.png).

