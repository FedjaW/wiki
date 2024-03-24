# Deployments

With a deployment you can execute a rolling update of your services.

## Commands with kubeclt

Create a deployment (will create a replica set (will create pods)):

```CLI
    kubectl create -f deployment-definition.yaml
```

Show all deployments:

```CLI
    kubectl get deployments
```

Show all replicasets:

```CLI
    kubectl get replicaset
```

Show all pods:

```CLI
    kubectl get pods
```

To see all the created objects at once:

```CLI
    kubectl get all
```