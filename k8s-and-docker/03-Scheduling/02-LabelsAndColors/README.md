# Labels and Selectors

The mechanism enables you to mark resources by a label and then filter by those labels.

```YAML
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    app: App1 # add as many labels as you like
spec:
  containers:
  - name: nginx
    image: nginx
    ports:
      - containerPort: 8080
```

## Commands

```CLI
kubectl get pods --selector app=App1 
```

## Annotations

While labels and selectors are used to group and select objects, annotations are used to record other details for informatary purpose.

```YAML
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    app: App1 # add as many labels as you like
  annotations:
    buildversion: 1.34 # example annotation
spec:
  containers:
  - name: nginx
    image: nginx
    ports:
      - containerPort: 8080
```
