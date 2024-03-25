# Daemon Sets

Daemon sets are like replica sets as it helps you to deploy instances of pods, but daemon sets runs one copy of your pod on each node of your cluster.
Whenever a new node is added to the cluster a replica of the pod is automtically added to that node.

The daemon set ensures that one copy is always present on all nodes on the cluster.

## Usecases

- Monitoring-Agent as a pod on all nodes
- Log-Agent as a pod on all nodes
- kube-proxy is a perfect example for a daemon set
- weave-net is for networking

## Yaml definition

```YAML
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: monitoring-daemon
spec:
  selector:
    matchLabels:
      app: monitoring-agent
  template:
    metadata:
      labels:
        app: monitoring-agent
    spec:
      containers:
      - name: monitoring-agent
        image: monitoring-agent
```

## Commands

```CLI
kubectl create -f daemon-set.yaml
```

```CLI
kubectl get daemonsets
```

## How is a daemonset scheduled?

Till v1.12 the nodeName in the pod is set to bypass the scheduler and place the pod on a specific node.

From v1.12 - uses NodeAffinity and default scheduler.
