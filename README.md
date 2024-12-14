# Stonewall, the best secret vault ever

## How to run an app

### Prerequisites

Installed kubectl >= v1.30.1, skaffold >= v2.12.0, minikube >= 1.33.1 and docker >= 26.1.3 + docker-buildx enabled, nginx-ingress-controller

### Start a local cluster with minikube on your machine

```bash
minikube start
```

### Installation of nginx ingress controller via glasskube just once (https://glasskube.dev/docs/)

```bash
glasskube install ingress-nginx
```

### Run following an one-liner in your project root folder

```bash
skaffold dev
```

### Open a kube-proxy tunnel between your host and services of the cluster (nginx & cocroachdb)

```bash
minikube tunnel --bind-address=127.0.0.1
```

Ctrl+c to stop the app.

### Endpoints and URLs

```text
localhost/db            - CockroachDB Cluster WebUI    
localhost/mail          - Mailhog WebUI    
localhost/api/v1/*      - REST/Api Enpoints    
localhost/*             - NextJS App    
```

