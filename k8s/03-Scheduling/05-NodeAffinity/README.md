# Node Affinity

With `Node Selectors` you can not express logical combinations like:

- Large OR Medium
- NOT Small

The `Node Affinity` feature provides us with advanced capabilities of pod placement on specific nodes. With great power comes great complexity.

The same configuration from `Node Selectors` looks like this with `Node Affinity`:

- `Node Selectors`

```YAML
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
  - name: data-processor
    image: data-processor
  nodeSelector:
    size: Large # the key and value "size: Large" are labels assigned to the corresponding node.
```

- `Node Affinity`

```YAML
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
  - name: data-processor
    image: data-processor
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: size
            operator: In # NotIn ...
            values:
            - Large # -Small ...
```

There are two `types` of `Node Affinity` available:

- requiredDuringSchedulingIgnoredDuringExecution
- preferredDuringSchedulingIgnoredDuringExecution

Read more here: <https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/>