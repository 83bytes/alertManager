# Toy Alert Manager

An alertmanager that supports performing arbitrary enrichments on alerts and take appripriate action.

These enrichments and actions have to be preconfigured.
The API is simple and extensible enough so as to enable users to extend the framework.

We will refer to this as the `tam`

## Quick Start

Over here, we will quickly set up the tam on a docker-compose and fire a test-event using curl to see how it works. <br>
Once we have a basic example running, we can then dive into the details.

_Note:_ This assumes that you have docker and docker-compose installed. Furthermore, you need a slack webhook url that is configured to send data to a channel.

1. Get a Slack Webhook and copy the secret (this is the part after the `https://hooks.slack.com/services/` in the webhook-url)
2. Put this secret in a (`alertmanager/.env` directory) as follows

```
WEBHOOK_SECRET=secret we copied in step 1
```

3. Run `make docker-build`
4. Run `make sed`
5. Run `docker compose up -d`
6. Send the basicWebhookPayload.json to the tam using curl

```
curl -v -H "Content-Type: application/json" -X POST localhost:8081/webhook -d @basicWebhookPayload.json
```

**NOTE** If everthing is configured correctly, then you should see a message in the channel that you have configured. If not, please look at the logs. The tam in docker-compose has debug logs enabled which are quite verbose.

Sample output

```
alert: NOOP_ALERT
action: SendToSlack
result of ENRICHMENT_STEP_1 enrichment(s):  ARG1,ARG2
```

# Design

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

For example, a typical config would look like this

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

We can use the `alertmanager` to generate a sample config. We can redirect this output to a file and then modify it to our needs.

```
$ ./alertmanager config generate-template
```

```
alert_pipelines:
    - alert_name: NOOP_ALERT
      enrichments:
        - step_name: ENRICHMENT_STEP_1
          enrichment_name: NOOP_ENRICHMENT
          enrichment_args: ARG1,ARG2
      actions:
        - step_name: ACTION_STEP_1
          action_name: NOOP_ACTION
          action_args: ARG1,ARG2
```

We can use the in-built config-validator to check if the config-file is up-to-spec or not

```
$ ./alertmanager config validate --config-file /path/to/file
```

The list of available [enrichments](enrichment/README.md) and [actions](action/README.md) are available in the respective docs.

## How does the TAM work ?

The tam accepts a JSON payload in the following format

```
{
  "version": "4",
  "groupKey": <string>, // key identifying the group of alerts (e.g. to deduplicate)
  "truncatedAlerts": <int>, // how many alerts have been truncated due to "max_alerts"
  "status": "<resolved|firing>",
  "receiver": <string>,
  "groupLabels": <object>,
  "commonLabels": <object>,
  "commonAnnotations": <object>,
  "externalURL": <string>, // backlink to the Alertmanager.
  "alerts": [
  {
    "status": "<resolved|firing>",
    "labels": <object>,
    "annotations": <object>,
    "startsAt": "<rfc3339>",
    "endsAt": "<rfc3339>",
    "generatorURL": <string>, // identifies the entity that caused the alert
    "fingerprint": <string> // fingerprint to identify the alert
  }
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

## SETUP on k8s (kind)

```
kind setup cluster
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
make sed
helm install prom-stack prometheus-community/kube-prometheus-stack -f deployment/kube-prometheus-stack.yml
kubectl apply -f deployment/toy_alert_manager.yml

```

## Caveats

It is designed to run in a secure enironment, hence there is no support for authentication and authorization.

DO NOT EXPOSE THIS TO THE OPEN INTERNET
