# TLS in Kubernetes

All communication between the nodes in Kubernetes must be secured and must be encrypted. Also the communication between clients (e.g. Administrator) and the kube-apiserver must be secured. And also the communication between all the kubernetes components must be secured.

Therefor Server Certificates for Servers and Client Certificates for Clients must be used.

## Server Certificates for Servers

### KUBE-API SERVER

The kube-api server is a server and is used to interact with the kubernetes cluster via rest-api, so it requieres certificates to establish a secured communication to its clients. So we generate a certificate key pair.

- apiserver.crt
- apiserver.key

### ETCD SERVER

The etcd server stores all the information about the cluster. So it need a pair of keys for itself.

- etcdserver.crt
- etcdserver.key

### KUBELET SERVER

The kubelet services, which run on all worker nodes, expose an https endpoint that the kube-apiserver talks to, therefor it must also be secured by a pair of keys.

- kubelet.crt
- kubelet.key

## Client Certificates for Clients

### Admin

The clients who access the kube-apiserver is us (the admins) through kubectl or rest-api. The admin user requires a certificate and key to authenticate to the kube-apiserver.

- admin.crt
- admin.key

### KUBE-SCHEDULER

The scheduler talks to the kube-apiserver to look for pods that requires scheudling and then get the kube-apiserver to schedule the right pod onto the right worker nodes. The scheduler is a client that accesses the kube-apiserver as far the kube-apiserver is concerned the schedulder is jsut another client like the admin user, so the schuduler needs to validates its identity using a client TLS-certifcate. So its need its own pair.

- scheudler.crt
- scheudler.key

### KUBE-CONTROLLER-MANAGER

The kube-controller-manager is antoher client that talks to the kube-apiserver.

- contoller-manager.crt
- contoller-manager.key

### KUBE-PROXY

The kube-proxy also needs a client certificate to interact with the kube-apiserver.

- kube-prodxy.crt
- kube-prodxy.key

### KUBE-API SERVER -> ETCD

The kube-apiserver talks to the etcd server, so as far as the etcd server ist concerned the kube-apiserver is just a client. So it needs to authenticate. The kube-apiserver can use the same pair of keys as it uses for serving its own api service.

- apiserver.crt
- apiserver.key

### KUBE-API SERVER -> KUBELET

- apiserver.crt
- apiserver.key

## Certificate Authority (CA)

- ca.crt
- ca.key