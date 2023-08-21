# dagger-argo-workflows

This is an example of how Dagger can operate with Argo Workflows

## Setup

- kind cluster [installation](https://kind.sigs.k8s.io/docs/user/quick-start/)
- argo-workflows default [installation](https://argoproj.github.io/argo-workflows/quick-start/)
- dagger engine daemonset in [template.yaml](./template.yaml)
  - `kubectl apply -f template.yaml`
- `kubectl create secret generic dagger-cloud --from-literal=token='a276ce43-e1ca-4427-a6ee-200d77b85b56'`
- argo workflow in [workflow.yaml](./workflow.yaml)
  - `argo submit -n argo --watch ./workflow.yaml`
