Hello OpenObserve
---

This is a sample of [OpenObserve](https://openobserve.ai/), a cloud native observability platform (Logs, Metrics, Traces).


## Description

This is a sample of logging, metrics and tracing of an application using [OpenObserve](https://openobserve.ai/).


## Structure

### Used language, tools, and other components

| language/tools                                 | description                                                                                                                   |
|------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------|
| [OpenObserve](https://openobserve.ai/)         | cloud native observability platform (Logs, Metrics, Traces)                                                                   |
| [Fluent Bit](https://fluentbit.io/)            | logging and metrics processor and forwarder                                                                                   |
| [OpenTelemetry](https://opentelemetry.io/)     | observability framework for instrumenting, generating, collecting, and exporting telemetry data such as traces, metrics, logs |
| [Go](https://github.com/golang/go)             | programming language                                                                                                          |
| [Kubernetes](https://kubernetes.io/)           | container orchestrator                                                                                                        |
| [Minikube](https://minikube.sigs.k8s.io/docs/) | tool for quickly sets up a local Kubernetes cluster                                                                           |
| [Skaffold](https://skaffold.dev/)              | tool for building, pushing and deploying your application                                                                     |

### Directories

```
.
├── .k8s          # => Kubernetes manifests
│     ├── base
│     └── overlays
├── api           # => API implementation
├── skaffold.yaml
└── (some omitted)
```


## Usage

1. Run the application in minikube

    ```shell
    make run
    ```

2. Call API

   ```shell
   curl http://hello-openobserve.localhost.com/app/api/hello
   ```

3. Confirm OpenObserve

   Please refer to the following for confirmation.

   - [Logs](https://github.com/hyorimitsu/hello-openobserve/blob/main/README.md#logs)
   - Metrics (TODO)
   - [Traces](https://github.com/hyorimitsu/hello-openobserve/blob/main/README.md#traces)

4. Stop the application in minikube

    ```shell
    make down
    ```


## Logs

### Architecture

This sample uses [Fluent Bit](https://fluentbit.io/) to ingest logs.  
Check the [official docs](https://openobserve.ai/docs/ingestion/logs/) for other support tools.

![logs_architecture](https://github.com/hyorimitsu/hello-openobserve/blob/main/docs/img/logs_architecture.png)

### Usage

To check the logs, please visit the following page in browser.  
(Please [see here](https://github.com/hyorimitsu/hello-openobserve/blob/main/.k8s/overlays/local/openobserve/configmap.yaml#L6-L7) for `Email` and `Password`)

http://hello-openobserve.localhost.com/web/logs?org_identifier=default

For example, the logs for this sample application can be obtained with the following query.

```shell
kubernetes_container_name='app-api'
```

![logs_ui](https://github.com/hyorimitsu/hello-openobserve/blob/main/docs/img/logs_ui.png)


## Traces

### Architecture

This sample uses [OpenTelemetry](https://opentelemetry.io/) to ingest traces.

![traces_architecture](https://github.com/hyorimitsu/hello-openobserve/blob/main/docs/img/traces_architecture.png)

### Usage

To check the traces, please visit the following page in browser.  
(Please [see here](https://github.com/hyorimitsu/hello-openobserve/blob/main/.k8s/overlays/local/openobserve/configmap.yaml#L6-L7) for `Email` and `Password`)

http://hello-openobserve.localhost.com/web/traces?org_identifier=default

For example, the traces for this sample application can be obtained with the following query.

```shell
service_name='hello-openobserve'
```

![traces_ui](https://github.com/hyorimitsu/hello-openobserve/blob/main/docs/img/traces_ui.png)
