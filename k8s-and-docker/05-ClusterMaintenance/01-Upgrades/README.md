# OS Upgrades

Sometimes you have to take down nodes as part of your cluster e.g. for maintenance purpose like upgrading software.

We will see the options to handle such cases.

## The scenario: 1 Master Node + 3 Worker Nodes

The worker nodes have pods running on them and one of them will go down.

Now, the pods on that node are not accessible any more. If you have replica of the pod on another node the user will notice nothing.
If the pod has no replicas of it running on a different node, the pod is temporarly not available.

For pods that are not within a replicaset: How does kubernetes handle such a scenario?

If the node comes back up fast enough the kubelet service starts and will create the pods again.
If the node was down for more than 5 min then the pods are terminated from that node. Kubernetes conciders them as dead. If the pod was part of a replicaset or the deployment then they will be recreated on other nodes.
The time it waits for a pod to come back online is known as the `pod-eviction-timeout` and is set on the `controller-manager`.

When the node comes back online after the `pod-eviction-timeout` it comes up blank without a pod scheduled on it. The pods that are part of a replicaset are scheduled on another node. The pod that are not part of a replicaset are just gone.

There is a better way to handle such a scenario:

You can drain the node of all the workloads from a node to other nodes. Technically the pods are gracefully terminated and are recreated on other nodes.

The node is then also marked as cordaned (marked as unschedulable).

```CLI
kubectl drain node-1
```

When the node comes back online it is still unschedulable, you need to uncordon it:

```CLI
kubectl uncordon node-1
```

The node that comes back online will have no pods running on it even after you mark it as uncordoned, because only pods that are created will be scheduled. The one that where put to other nodes will keep running on this other nodes.