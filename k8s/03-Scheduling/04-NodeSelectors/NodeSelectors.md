# Node Selectors

Example: You have three nodes with different hardware resources. On of them has more horsepower than the other and you have three container to run with different resource needs. You want to place the most hungry resource pod on the node with the most horstpower. To solve this we can set a limitation on the pods so that they run only the particular nodes. There are two options to do this:

- Node Selectors
- Node Affinity (see next chapter)

## Commands to setup Node Selectors

```YAML
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
  - name: data-processor
    image: data-processor
  nodeSelector:
    size: Large # the key and value "size: Large" are labels assigned to the corresponding node.
```

The scheduler uses labels to identidy the correct node to place the pod on.

To label the node:

```CLI
kubectl label node <node-name> <label-key>=<label-value>
```

```CLI
kubectl label node node-1 size=Large
```

When the label is set you can create the pod:

```CLI
kubectl create -f pod-definition-yml
```
