# Toy Alert Manager

An alertmanager that supports performing arbitrary enrichments on alerts and take appripriate action.

These enrichments and actions have to be preconfigured.
The API is simple and extensible enough so as to enable users to extend the framework.

We will refer to this as the `tam`

## Design

The tam is a simple webhook server.

We can configure the tam to enrich alerts by pulling data from external systems AND take actions.

The enrichments and the actions that are possible / relevant for each alert is highly context dependent and is upto the user to build and configure.

The collection of `enrichments` and `actions` for a given `alert` is called an `alertPipeline`. We will see how to configure such a pipeline below.

## Configuring an Alert-Pipeline

The tam is configured by using a config file (yaml format) which defines multiple alertpipelines.

Each alertpipeline is defined by

- an AlertName
- A list of Enrichments
- A list of Actions.

For example, a typlcal config would look like this

```
alert_pipelines:
  - alert_name: KubePodCrashLooping
    enrichments:
      - step_name: enrichment_step_1
        enrichment_name: GET_DATA
        enrichment_args: "promql"
    actions:
      - step_name: action_step_1
        action_name: NotifySLack
        action_args: "url"
```

## How does the TAM work ?

The tam accepts a JSON payload in the following format

```
{
  "version": "4",
  "groupKey": <string>,              // key identifying the group of alerts (e.g. to deduplicate)
  "truncatedAlerts": <int>,          // how many alerts have been truncated due to "max_alerts"
  "status": "<resolved|firing>",
  "receiver": <string>,
  "groupLabels": <object>,
  "commonLabels": <object>,
  "commonAnnotations": <object>,
  "externalURL": <string>,           // backlink to the Alertmanager.
  "alerts": [
    {
      "status": "<resolved|firing>",
      "labels": <object>,
      "annotations": <object>,
      "startsAt": "<rfc3339>",
      "endsAt": "<rfc3339>",
      "generatorURL": <string>,      // identifies the entity that caused the alert
      "fingerprint": <string>        // fingerprint to identify the alert
    },
    ...
  ]
}
```

note: This is detailed in the prometheus [webhook receiver docs](https://prometheus.io/docs/alerting/latest/configuration/#webhook_config)

The alerts object is a list that can contain multiple `alert`. Each of them are of the following format

```
{
  "annotations": {
    "description": "Pod customer is restarting 2.11 times / 10 minutes.",
    "runbook_url": "",
    "summary": "Pod is crash looping."
  },
  "labels": {
    "alertname": "KubePodCrashLooping",
    "cluster": "cluster-main",
    "container": "rs-transformer",
    "endpoint": "http",
    "job": "kube-state-metrics",
    "namespace": "customer",
    "pod": "customer",
    "priority": "P0",
    "prometheus": "monitoring/kube-prometheus-stack-prometheus",
    "region": "us-west-1",
    "replica": "0",
    "service": "kube-prometheus-stack-kube-state-metrics",
    "severity": "CRITICAL"
  },
  "startsAt": "2022-03-02T07:31:57.339Z",
  "status": "firing"
}
```

The tam uses the `labels.alertname` as a primary identifier to identify alerts and identify configured pipelines for said alerts. Thus, the above configured pipeline for `KubePodCrashLooping` would match this alert and then execute the enrichments and then the Actions.

While the Enrichments and Actions can be built by the user using a certain framework, it should be noted that the enrichment runtime has a full copy of the alert body it was configured for. Similarly the alert runtime as a full copy of the alert as well the enrichments and their corresponding output. We shall see how build our own enrichments and actions in a bit.

## Building Actions and Enrichments

[Actions](./action/README.md) and [Enrichments](./enrichment/README.md) live in their own directories. There are some sample alerts and enrichments pre-built for ease of use.

## SETUP

```
kind setup cluster
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prom-stack prometheus-community/kube-prometheus-stack -f values.yml

```

## Caveats

It is designed to run in a secure enironment, hence there is no support for authentication and authorization.

DO NOT EXPOSE THIS TO THE OPEN INTERNET
