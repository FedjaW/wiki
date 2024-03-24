# ETCD

## External ETCD

To see information about etcd do -> ssh into the external etcd server and run:

```CLI
ps -ef | grep -i etcd
```

## How many nodes are part of the ETCD cluster that etcd-server is a part of?

address: `https://192.25.30.11:2379` is taken from `--listen-client-urls` from the command above.

--trusted-ca-file=/etc/etcd/pki/ca.pem
--cert-file=/etc/etcd/pki/etcd.pem
--key-file=/etc/etcd/pki/etcd-key.pem

```CLI
ETCDCTL_API=3 etcdctl --endpoints=192.25.30.11:2379 --cacert=/etc/etcd/pki/ca.pem --cert=/etc/etcd/pki/etcd.pem --key=/etc/etcd/pki/etcd-key.pem member list
```

## Take a backup of etcd on cluster1 and save it on the student-node at the path /opt/cluster1.db

On this cluster1 etcd is deploys as a pod (stacked) and so you can get the neccessary informations with:

```CLI
kubectl describe pod -n kube-system etcd-cluster1-controlplane
```

ssh into the node where the etcd is running -> cluster1-controlplane

```CLI
ssh cluster1-controlplane
```

Take a backup:

--advertise-client-urls=https://192.25.30.19:2379
--trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt
--cert-file=/etc/kubernetes/pki/etcd/server.crt
--key-file=/etc/kubernetes/pki/etcd/server.key

```CLI
ETCDCTL_API=3 etcdctl --endpoints=192.25.30.19:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key snapshot save /opt/cluster1.db
```

This will create a backup on the cluster1-node but we want it to be on a different node. So exit out and go to the desired node then copy the file over:

```CLI
scp cluster1-controlplane:/opt/cluster1.db /opt/cluster1.db
```

## Restore the backup

ssh into the etcd-server and then restore:

```CLI
ETCDCTL_API=3 etcdctl snapshot restore /root/cluster2.db --data-dir /var/lib/etcd-data-new
```

No certs are needed for the restore.
