# Manual Scheduling

How does scheduling work? Let us start with a simple example:

Concider a pod-definition.yml file:

```YAML
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    name: nginx
spec:
  containers:
  - name: nginx
    image: nginx
    ports:
      - containerPort: 8080
  nodeName:
```

Every pods has a field `nodeName` which by default is not set.
The scheduler looks out for those and will schedule it onto a node and puts the node-name in the nodeName field.

If there is no scheduler the pods will stay in a pending state.
You can schedule the pods on your own by set the nodeName field while creating the pod. It can not be updated later.
Another way to assign a pod to a node is to create a binding object and send a post request to the pods binding api.
This mimicking what the actual schedulder does.

```YAML
apiVersion: v1
kind: Binding
metadata:
  name: nginx
target:
  apiVersion: v1
  kind: Node
  name: node02
```

```CLI
curl --header "Content-Type:application/json" --request POST --data '{"apiVersion":"v1", "kind":"Binding" ...}' gttp://$SERVER/api/v1/namespaces/default/pods/$PODNAME/binding/
```
