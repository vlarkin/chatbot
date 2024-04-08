FROM --platform=$BUILDPLATFORM quay.io/projectquay/golang:1.20 AS build
WORKDIR /go/src/app
COPY . .
ARG TARGETOS TARGETARCH
RUN make build

FROM alpine
WORKDIR /
COPY --from=build /go/src/app/chatbot .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./chatbot"]
