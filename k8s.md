# k8s setup
We are using helm3 to deploy all required components.

## k8s Dashboard (Optional)

https://hub.helm.sh/charts/k8s-dashboard/kubernetes-dashboard

```
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard
helm repo update
helm install helm3-k8s-dashboard kubernetes-dashboard/kubernetes-dashboard --version 2.3.0
```

## Metrics Server

https://hub.helm.sh/charts/bitnami/metrics-server
```
helm repo add bitnami https://charts.bitnami.com/bitnami
helm upgrade helm3-metric-server bitnami/metrics-server  --set apiService.create=true
```

