# Static pods

Let's say there is no master at all. Let's say we have only a `kubelet` isntalled on a `worker node`; and on that node docker is installed.

The only thing the `kubelet` knows to do is create pods. But in our scenario we do not have an `API-Server` to provide pod details. Usually we need a pod.yaml file to provide the pod details.

## Commands

To inspect the satic pods we can not use the kubeclt because we have no kube-apiserver:

```CLI
docker ps
```

## Notes

The kubelet can create pods from the defintion files in the `etc/kubernetes/manifests` folder, but it also provides an http api endpoint. The `kube-apiserver` uses the endpoint to provide input to the kubelet.

The apiserver is aware of all pods (even the static ones):
This command will give you all pods:

```CLI
kubectl get pods
```

## Use Case

You can use the static pods to deploy the control plan components itself as pods on a node.

1. Install kubelet on the master nodes
2. Create pods definition files that uses docker images of the various control plane coponents. (apiserver.yml; controller-manager.yaml; etcd.yaml)
3. Place the definition files in the designated manifest folder (/etc/kubernetes/manifests)

The `kubeadmin tool` sets up a cluster in such a way. That is why we can see the controlplane components as pods when executing `kubectl get pods -n kube-system`.

## Common question

What is the difference between daemonsets and static pods?

### Static pod

- Created by the kubelet
- Deploy controlplane components as static pods
- Ignored by the kube-scheduler

### Daemonsets

- Created by kube-apiserver (daemonset controller)
- Deploy monitoring agents, Loggin, Agents on nodes
- Ignored by the kube-schedule

## Practical - Notes

To identify a static pod look out for the name ending of the pod. If the ending is the name of a node, it is a static pod. (Not sure why this is.)

Here the name end with controlplane which is a node name

- kube-apiserver-controlplane

