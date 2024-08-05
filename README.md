# Stonewall, the best secret vault ever

## How to run an app

### Prerequisites

Installed kubectl >= v1.30.1, skaffold >= v2.12.0, minikube >= 1.33.1 and docker >= 26.1.3 + docker-buildx

### Start a local cluster with minikube on your machine

```bash
minikube start
```

### Run following an one-liner in your project root folder

```bash
skaffold dev
```

### Open a kube-proxy tunnel between your host and services of the cluster (caddy & cocroachdb)

```bash
minikube tunnel --bind-address=127.0.0.1
```

Ctrl+c to stop the app.


