# Kube Proxy

In a kubernetes cluster every pod can reach every other pod. This is accomplished by a POD-Network.

The IP addresses of the pods are not garantueed to be the same. They can change over time as pods get replaced. So a `service` is introduced to make sure any given pod is reachable even the IP address changes.

A `service` has no interfaces an no listening port. A `service` is a virtual component that only lives in the kubernetes memory.

A `service` is accessible by any node within the cluster. That is achived with `kube-proxy`. `kube-proxy` is a process that runs on each pod. Its job is to look out for new `services` and for every `service` it creates the appropriate rules on each node to forward traffic to those `services` to the backend pods. One way it does this is using IP-table rules.

A `kube-proxy` is deployed as a deamonset by kubeadm.