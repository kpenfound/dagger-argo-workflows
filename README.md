# dagger-argo-workflows

This is an example of how Dagger can operate with Argo Workflows

## Setup

- kind cluster [installation](https://kind.sigs.k8s.io/docs/user/quick-start/)
- argo-workflows default [installation](https://argoproj.github.io/argo-workflows/quick-start/)
- dagger engine daemonset in [daemonset.yaml](./daemonset.yaml)
  - `kubectl apply -f template.yaml`
- create PersistentVolumeClaim for dagger runtime dependencies in [gomodcache.yaml](./gomodcache.yaml)
  - `kubectl apply -f ./gomodcache.yaml`
- create kubernetes secret for dagger cloud token
  - `kubectl create secret generic dagger-cloud --from-literal=token='{YOUR CLOUD TOKEN}'`
- argo workflow in [workflow.yaml](./workflow.yaml)
  - `argo submit -n argo --watch ./workflow.yaml`
