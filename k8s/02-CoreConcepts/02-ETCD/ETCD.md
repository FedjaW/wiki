# ETCS

## What is ETCD

ETCD is a distributed reliable key-value store that is simple, secure and fast.
A key value store stores information in the form of documents or pages.

Example:

| key   |      Value      |
|----------|:-------------:|
| Name |  John Doe |
| Age |    45   |
| Salary | 5000 |

| key   |      Value      |
|----------|:-------------:|
| Name |  Dave Smith|
| Age |    34   |
| Salary | 4000 |

## Install ETCD

1. Download Binaries

```CLI
    curl -L https://githubg.com/etcd-io/etcd/releases/download/etcd-v3.3.11-linux-amd64.tar.gz -o etcd-v3.3.11-linux-amd64.tar.gz
```

2. Extract

```CLI
    tar xzvf etcd-v.3.3.11-linux-amd64.tar.gz
```

3. Run ETCD Service

```CLI
    ./etcd
```

When you run ETCS it starts a service that listens on port 2379 by default. You can connect clients to it, to store and retrieve information. A default client, that comes with etcd is the etcd-control-client. This client is a command line client for etcd.

(For version 2)

```CLI
    ./etcdctl set key1 value1
    ./etcdctl get key1
```

## ETCD Versions

Find out version (for API-Version 2):

```CLI
    ./etcdctl --version
```

Set API version bevor a command:

```CLI
    ETCDCTL_API=3 ./etcdctl version
```

Set API version permanent for the session:

```CLI
    export ETCDCTL_API=3 ./etcdctl version
```

Save and retrieve information (For version 3):

```CLI
    ./etcdctl put key1 value1
    ./etcdctl get key1
```

## ETCDs Role in Kubernetes

The ETCD datastore stores information about the cluster for:

- Nodes
- PODs
- Configs
- Secrets
- Accounts
- Roles
- Bindings
- Others

Every information you see when you run the `kubectl get` command is from `ETCD` server.
Every change you make to your cluster such as adding nodes are updated in the `ETCD` server.

Different ways of deploying ETCD:

1. Deployed from scratch (manual)
2. With the Kubeadm tool

A deployment from scratch will run etcd as a service in your master node.
Deploying with kubeadm will deploy the etcd server in the kube-system namespace as an own container.

To show all pods within the kube-system namespace run:

```CLI
    kubectl get pods -n kube-system
```

To list all etcd keys run following command (inside the etcd-master POD):

```CLI
    kubectl exec etcd-master -n kube-system etcdctl get / --prefix -keys-only
```

Below is the final form (with all certs):

```CLI
kubectl exec etcd-master -n kube-system -- sh -c "ETCDCTL_API=3 etcdctl get / --prefix --keys-only --limit=10 --cacert /etc/kubernetes/pki/etcd/ca.crt --cert /etc/kubernetes/pki/etcd/server.crt  --key /etc/kubernetes/pki/etcd/server.key"
```

__What I get when executing those commands (kubectl exec etcd-docker-desktop -n kube-system -- etcdctl version):__

```CLI
kubectl exec [POD] [COMMAND] is DEPRECATED and will be removed in a future version. Use kubectl exec [POD] -- [COMMAND] instead.
```

__This works fine:__

```CLI
kubectl exec etcd-docker-desktop -n kube-system -- etcdctl version
```
