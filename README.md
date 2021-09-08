# toggl_exporter

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This is a simple server that scrapes [toggl track](https://toggl.com/track/) stats and exports them via HTTP for Prometheus consumption.

## Installation and Usage

The `toggl_exporter` listens on HTTP port 9981 by default. See the `--help` output for more options.

*Note: You must pass toggl api key via `--toggl.api-key` command flag or `TOGGL_API_KEY` environemnt variables.*

### Kubernetes

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: toggl-exporter
  labels:
    app: toggl-exporter
spec:
  replicas: 1
  selector:
    matchLabels:
      app: toggl-exporter
  template:
    metadata:
      labels:
        app: toggl-exporter
    spec:
      containers:
      - name: toggl-exporter
        image: ghcr.io/44smkn/toggl_exporter:latest
        ports:
        - containerPort: 9981
        env:
        - name: TOGGL_API_KEY
          valueFrom:
            secretKeyRef:
              name: toggl-secret
              key: api-key
---
apiVersion: v1
kind: Secret
metadata:
  name: toggl-secret
type: Opaque
data:
  api-key: YWRtaW4=
```

### AWS SAM

This is Extremely rare use cases.  
See: https://github.com/44smkn/toggl_exporter_serverless

## Collectors

| Name            | Description |
| --------------- | ----------- |
| ProjectDuration | Title       |
