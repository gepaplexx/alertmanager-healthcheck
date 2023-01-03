# alertmanager-healthcheck

This repository includes an application written in Golang that serves the purpose of 
receiving heartbeats from Prometheus alertmanagers. It exposes two endpoints over HTTP on port 2112:

- `/inc`: Receives POST-Requests from alertmanagers in this 
  [format](https://prometheus.io/docs/alerting/latest/configuration/#webhook_config).

- `/metrics`: Endpoint for Prometheus metrics in the OpenMetrics format.

The `/metrics` endpoint provides a metric called `alertmanager_status`, which is a counter vector
that increments per alert of an alertmanager. The Alerts pushed to `/inc` should
include a label called `gepardec_cluster` in order to differentiate the counters of alertmanagers of
different clusters. 

## Usage

Pull from GitHub Container Registry and Run: 

```
docker run -d -p 2112:2112 ghcr.io/gepaplexx/alertmanager-healthcheck:main
```

## Integration

This application is best used with the preconfigured 
[Watchdog](https://docs.openshift.com/container-platform/4.11/monitoring/managing-alerts.html#applying-custom-alertmanager-configuration_managing-alerts)
alert of RedHat OpenShift Container Platform.

The exposed `alertmanager_status` metric can be used to detect outages of remote alert systems with a 
local alertmanager. 
