# Taints and Tolerations

Nodes can have taints and pods can have tolerations.
Pods that are not tolerant to the taint of a node can not be placed on that node.

The scheduler will try to place the node on another node that has no taint set or a taint that the pod is tolerant to.

If a pod is tolerant to a taint, does not mean this pod will always be placed there. If there are other nodes with no taints or in other words the pod is also tolerant to other nodes the scheduler can place the pod wherever he wants to.
=> Taints and tolerations does not tell the pod to go to a particular node!
=> It tells the node to accept only pods with certain tolerations!

## Taints - Node

Blueprint of the command:

```CLI
kubectl taint nodes node-name key=value:taint-effect
```

What will happen to the pods that do not tolerate this taint?
There are three `taint-effect` options:

- `NoSchedule` -> The pods will not be scheduled on the node (explanation above)
- `PrefereNoSchedule` -> The system will try to avoid placing the pods on the node but it is not garantueed
- `NoExecute` -> New pods will not be scheduled on the node and existing pods on the node if any will be evicted if they do not tolerate the taint.

Example command:

```CLI
kubectl taint nodes node1 app=blue:NoSchedule
```

## Tolerations - Pods

```YAML
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
spec:
  container:
  - name: nginx-container
    image: nginx
  tolerations: # all values have to be encoded in double quotes
  - key: "app"
    operator: "Equal"
    value: "blue"
    effect: "NoSchedule"
```

## Interesting notes

For now we only had worker nodes in mind. But what is about the controlplain (master node)?

The master node is technically just another node that has all the capabilites of hosting a pod plus it runs all the management software.
But the scheduler does not place any pods an the master node. Why is that?

When the kubernetes cluster is first setup a taint is set on the master node automatically that prevents any pods beeing scheduled on the node.

It is best practice to not schedule workload on the master node.

To see taint set:

```CLI
kubectl describe node kubemaster | grep Taint
```

To remove taint:

```CLI
kubectl taint node kubemaster <taint-name>-
```

Example

```CLI
kubectl taint node controlplane node-role.kubernetes.io/control-plane:NoSchedule-
```
