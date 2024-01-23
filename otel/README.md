I am not in time to implement the middle/senior/principal, but I know what and how to do that along with TraceID
I'm not so proficient with kuberntes and monitoring, so head a little exploding:) But it is very interesting and challenging.

All would be done with Helm even for middle(without FluxCD). Though it can be done directly with manifests, 
like `kubectl apply -f .../releases/latest/download/opentelemetry-operator.yaml` 
I but want to do it with helm to adapt it later for Flux.

The hardest part for me is to make correct configs and labels/tags in configs for Kubernetes.
So I am in the process now.


My plan:

- new plan:

	- install grafana/loki-stack with helm
		- Fluent Bit + Loki + Grafana(datasources)
		- promtail, logstash and filebeat are disabled
		- some pseudo-code:
			- helm repo add grafana https://grafana.github.io/helm-charts
			- helm repo update
			- change values
			- helm install loki grafana/loki-stack -n loki-stack --create-namespace -f loki-stack-values.yaml

	- install OpenTelemetry operator helm with a self-signed certificate and secret
		- run Collector and Instrumentation

	- install prometheus helm
		- disable grafana
		- disable prometheus-node-exporter


- old plan(deprecated):

	- install separate Fluent Bit helm
		- with var FLUENT_LOKI_URL in values
		- some pseudo-code:
			- helm repo add fluent https://fluent.github.io/helm-charts
			- helm repo update
			- set config: in values.yaml		
			- helm install fluent-bit fluent/fluent-bit -f values.yaml


	- install otel operator helm with a self-signed certificate and secret
		- some pseudo-code:
			- helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
			- helm repo update
			- helm install my-opentelemetry-operator open-telemetry/opentelemetry-operator ...			 
			- add OpenTelemetryCollector with spec config


	- install separate loki with helm
		- with promtail disabled
		- some pseudo-code:
			...
		
	- install prometheus operator with helm
		- some pseudo-code:
			- helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
			- helm repo update
			- helm show values prometheus-community/kube-prometheus-stack > prometheus-values.yaml
			- change values to add scrape targets
			- disable grafana as it would be installed in grafana stack		
			- disable prometheus-node-exporter as Fluent Bit would do that
			- helm install prometheus-community/kube-prometheus-stack  -f prometheus-values.yaml
			- helm list

	- check all with kubectl


- As for distributed tracing:
	- if Kbot go app:
		- add tracing to go app with already used lib
			- "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
			- "go.opentelemetry.io/otel/sdk/trace"
		- create spans, create and send traceIds with trace.NewTracerProvider and traceExporter
		- when for example we receive payload from telegram API

	- implement passing traceID with labels/tags in otel collector, Fluent Bit, Loki, Prometheus and Grafana with Tempo
		- not sure about details yet



