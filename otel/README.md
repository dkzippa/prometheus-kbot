<br/><br/>
>> Notice to mentors: 
>> on deadline 23.01.2024 v1 was ready only(it is with video below)
>> now I working on v2

<br/><br/>

# V2 - implementing Monitoring stack with Flux and distributed tracing  (in progress)
### Status: 

- ##### add metrics server to Kubernetes
	- `helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server/`
	- `helm repo update`
	- `helm upgrade --install metrics-server metrics-server/metrics-server -n kube-system --values=metrics-server.yaml`
	- if needed `helm uninstall metrics-server -n kube-system --debug`

	- check 
		- `k get po -n kube-system -l "app.kubernetes.io/name=metrics-server"`
		- `kubectl get --raw /api/v1/nodes/ip-172-31-7-243/proxy/metrics/resource`
		- `export METRICS_POD=$(k get po -n kube-system -l "app.kubernetes.io/name=metrics-server" -o name)`
		- `echo $METRICS_POD`
		- `k logs $METRICS_POD -n kube-system`
		- `k describe $METRICS_POD -n kube-system`

- ##### use `otel/helm` folder for helm values
<br/><br/>
- #####  using `log` namespace for testing purposes, `monitoring` could be used for real
	- `k create ns log` 



- ##### fluent bit: 
	- `helm repo add fluent https://fluent.github.io/helm-charts`
	- `helm repo update`
	- `helm upgrade -install fluent-bit fluent/fluent-bit -n log --values=fluent-bit.values.yaml`

	- check:
		`export FLUENT_POD=$(kubectl get pods --namespace log -l "app.kubernetes.io/name=fluent-bit,app.kubernetes.io/instance=fluent-bit" -o jsonpath="{.items[0].metadata.name}")`
		`kubectl --namespace log port-forward $FLUENT_POD 2020:2020`
		`curl http://127.0.0.1:2020`


- ##### OpenTelemetry Collector (contrib)

	- `helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts	`
	- `helm repo update`
	- `helm upgrade --install otel-collector open-telemetry/opentelemetry-collector -n log --values otel-collector.values.yaml`
	
	- check:
		- `export OTEL_POD=$(kubectl get po -n log -l "app.kubernetes.io/name=opentelemetry-collector" -o jsonpath="{.items[0].metadata.name}"); `
		- `k logs $OTEL_POD -n log`
		- `k top pod $OTEL_POD -n log`

		- `kubectl -n log port-forward $OTEL_POD 3030:3030`
		`curl http://127.0.0.1:3030/ -H'Content-Type: application/json'`
		`curl http://127.0.0.1:3030/api/v2/spans -H'Content-Type: application/json'`



- ##### Grafana + Loki + Tempo (I know there is a grafana-loki-stack with several tools included, but wanted to do it separately to better understand the process, which would helm to implement  TraceID later )

	- `helm repo add grafana https://grafana.github.io/helm-charts`
	- `helm repo update`


	- Loki:
		- `helm upgrade --install loki grafana/loki -n log --values loki.values.yaml `

		- check:
			- `k get po -n log -l app.kubernetes.io/instance=loki -o yaml | grep -A 5 labels`
			- `LOKI_POD="$(kubectl get pod -l "app.kubernetes.io/instance=loki,app.kubernetes.io/component=single-binary" -n log -o jsonpath='{.items[0].metadata.name}')"`
			- `echo $LOKI_POD`
			- `k logs -n log $LOKI_POD -f `
			- `kubectl -n log port-forward $LOKI_POD 3100`

	- Grafana
		
		- `helm install grafana grafana/grafana -n log --values grafana.values.yaml `

		- check:
			- `export GRAFANA_POD=$(kubectl get pods -n log -l "app.kubernetes.io/name=grafana,app.kubernetes.io/instance=grafana" -o jsonpath="{.items[0].metadata.name}")`
			- `echo $GRAFANA_POD`
			- `kubectl -n log port-forward $GRAFANA_POD 3100 #3000`

		- get pass(for admin user):
			- `kubectl get secret -n log grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo`
			
	- Tempo # for tracing
		- `helm upgrade --install tempo grafana/tempo-distributed -n log --values tempo.values.yaml `


- ##### Prometheus 
	- `helm repo add prometheus-community https://prometheus-community.github.io/helm-charts`
	- `helm repo update`
	- `helm upgrade --install kube-prometheus-stack -n log --values prometheus.values.aml `


- ##### For flux i plan to do my Chart with all these tools as dependencies and their values files
<br/><br/>

- ##### As for distributed tracing I plan:
	- in Kbot go app:
		- add tracing to go app with already used lib
			- "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
			- "go.opentelemetry.io/otel/sdk/trace"
		- create spans, create and send traceIds with trace.NewTracerProvider and traceExporter
		- when for example we receive payload from telegram API

	- implement passing traceID with labels/tags in otel collector, Fluent Bit, Loki, Prometheus and Grafana with Tempo
		- in process



<br/><br/><br/>




# V1 (23.01.2024) - base level with docker-compose implemented:

### status:
I am not in time to implement the middle/senior/principal, but I know what and how to do that along with TraceID
I'm not so proficient with kuberntes and monitoring, so head a little exploding:) But it is very interesting and challenging.

All would be done with Helm even for middle(without FluxCD). Though it can be done directly with manifests, 
like `kubectl apply -f .../releases/latest/download/opentelemetry-operator.yaml` 
I but want to do it with helm to adapt it later for Flux.

The hardest part for me is to make correct configs and labels/tags in configs for Kubernetes.
So I am in the process now.

![Stack video](_assets/otel.gif)
