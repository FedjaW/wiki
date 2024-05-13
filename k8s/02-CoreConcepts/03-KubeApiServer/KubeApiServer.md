# Kube-apiserver

## Worklow of creating a pod

A behind the scenes example:

1. The plain http command without kubectl looks like this:

```CLI
    curl -X POST /api/v1/namespace/defailt/pods ....[other]
```

2. The kube-apiserver will authenticate the user, validate the request and then
3. Creating a pod object without assigning it to a node
4. Updates the information in the etcd-server
5. send back the information to the user that the pod has been created
6. The kube-scheduler monitors the api-server and realizes that there is new pod with no node assigned. The scheduler identifies the right node to put the pod on and communicates that back to the api-server
7. The api-server updates the information in the etcd-cluster
8. The api-server passes that information to the kubelet of a worker node
9. The kublet then creates the pod on the node and instructs the container runtime engine to deploy the container image
10. Once this is done the kubelet updates the satus back to the api-server
11. And the api-server then updates the data in the etcd-server

This process is folllowed whenever a change to the cluster is performend.

## Tasks of the Kube-apiserver

1. Authenticate User
2. Validate Request
3. Retrieve Data
4. Update ETCD
5. SCheduler
6. Kubelet

Note: The `kube apisever` is the only one who interacts with the `ETCD`-Datastore.

## View api-server - kubeadm

The kubeadm tool deploys the kube-apiserver as a pod in the kube-system namespace on the master node.

This command will list all the pods within the kube-system namespace.

```CLI
    kubectl get pods -n kube-system
```

I think within that pod you can get the yaml defintion file of the apiserver like this:

```CLI
    cat /etc/kubernetes/manifest/kube-apiserver.yaml
```

## View api-server - without kubeadm (manual setup)

In an non kubeadm environment you can inspect the api-server by accessing the api-server service.
As a difference to the kubeadm way here the api-server is not deployed in an own pod but running as a standalone service.

```CLI
    cat /etc/systemd/system/kube-apiserver.service
```

## Also a way of inspecting the kube-apiserver

Execute this on the master node:

```CLI
    ps -aux | grep kube-apiserver
```
