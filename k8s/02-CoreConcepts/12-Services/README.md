# Services

Services enable communaction within and outside of the application.

## Service Types

### Node Port

- Opens a port on the node to allow access from outside the node.
- _In other words:_ Makes an internal port in the kubernetes cluster accessible on a port on the node.

### Cluster IP

The Service creates a virtual IP inside the cluster to enable communiation between different services such as a set of front-end servers to a backend servers.

### Load Balancer

Provisions a load balancer for our application.

## Commands

```CLI
    kubectl create -f service-definition.yaml
```

```CLI
    kubectl get services
```
