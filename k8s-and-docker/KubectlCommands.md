# Overall commands

```CLI
kubectl replace --force -f nginx.yml
```

```CLI
kubectl run bee --image=nginx --dry-run=client -o yaml > bee.yaml
```

```CLI
kubectl get pods -o wide
```

You can edit a deployment that you created imperative!
You can not edit a pod (only few properties). A pod has to be deleted and then created again.

```CLI
kubectl edit deploy <deployment-name>
```

All namespaces use flag: `-A`

```CLI
kubectl get pods -A
```

Show all pods and on what node they are running:

```CLI
kubectl get pods -o wide
```

Describe a pod from another namespace

```CLI
kubectl describe pod -n kube-system etcd-controlplane 
```