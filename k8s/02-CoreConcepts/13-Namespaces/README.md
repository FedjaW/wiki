# Namespace

In order to talk from one namespace to a resource that is in another namespace you can refere to it by: `<your-service>.dev.svc.cluster.local`. This is possible because, when the service is created, a `DNS` entry is added automatically in this format.

Let us break this format down:

`cluster.local` -> default domain name of the kubernetes cluster
`svc` -> subdomain for the service
`dev` -> namespace
`your-service` -> name of your service

## Commands

This command lists all the pods, but only in the default namespace.

```CLI
kubectl get pods
```

This command list pods in another namespace.

```CLI
kubectl get pods --namespace=kube-system
```

To create pods in a specific namespace other than default:

```CLI
kubectl create -f pod-definition.yml --namespace=dev
```

Or put the namespace in the pod-definition.yml like it is in the example file.

To create a new namepsace:

```CLI
kubectl create -f namespace-dev.yml
```

Or

```CLI
kubectl create namespace dev
```

To switch the namespace permantly:

```CLI
kubectl config set-context $(kubectl config current-context) ---namespace=dev
```

This list all pods:

```CLI
kubectl get pods --all-namespaces
```
