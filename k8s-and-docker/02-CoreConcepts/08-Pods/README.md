# POD

A `POD` is the smallest instance in kubernetes. Usually pods have a one to one relation to a container. But there is the concept of multi-container pods.

Sometimes you need a helper container. This helper container lives a longside with your "main" container. Both container can communicate via localhost because they share the same network space.

## pod-definition.yml

The example show a bare minimum of yaml configuration to deploy a pod.

To deploy it run:

```CLI
kubectl create -f pod-definition.yml
```

To inspect the pods run:

```CLI
kubectl get pods
```

To inspect a pod details run:

```CLI
kubectl describe pod myapp-pods
```

## Commands with kubectl

To get all pods (in the default namespace):

```CLI
    kubectl get pods
```

Create a new pod with the nginx image:

```CLI
    kubectl run nginx --image=nginx
```

To inspect details of a pod (nginx is the name of the pod to inspect):

```CLI
    kubectl describe pod nginx
```

To dry run a new pod with the nginx image:

```CLI
    kubectl run redis --image=redis --dry-run=client -o yaml > redis.yaml
```

To apply the created yaml from the dry run:

```CLI
    kubectl apply -f redis.yaml
```
