# Kubelet

The kubelet runs on the worker node

## Tasks

- Register the Node it runs on to the kube-apiserver
- Creates PODs by talking to the container runtime (Docker)
- Monitor Node & PODs

## Installing Kubelet

Attention: Kubeadm will not deploy kubelets.
You must always manually install the kubelets.
