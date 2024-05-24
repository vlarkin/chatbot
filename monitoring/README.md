
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

Install cert-manager to add certificates and certificate issuers as resource types in Kubernetes clusters. This tool simplifies the process of obtaining, renewing, and using certificates.  

```
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.14.5/cert-manager.yaml
```

Create a namespace for deploying the monitoring stack:  

```
kubectl create namespace monitoring
```

## Monitoring

Install the OpenTelemetry collector first:

```
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm install otel-collector open-telemetry/opentelemetry-collector -n monitoring --set image.repository="otel/opentelemetry-collector-k8s" --values otel-collector/values.yaml
```

Then install the rest of the monitoring tools:

```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add fluent https://fluent.github.io/helm-charts
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/prometheus  -n monitoring --values prometheus/values.yaml
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

