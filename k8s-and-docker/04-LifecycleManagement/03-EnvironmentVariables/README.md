# Environment Variables

CLI command with docker:

```CLI
docker run -e APP_COLOR=pink simple-webapp-color
```

YAML representation:

```YAML
apiVersion: v1
kind: Pod
metadata:
  name: simple-webapp-color
spec:
  containers:
    - name: simple-webapp-color
      image: simple-webapp-color
      ports:
        - containerPort: 8080
      env: # --------------------------------- ENV variables
        - name: APP_COLOR
          value: pink
```

## Various types of environment variables

### Plain Key Value

```YAML
...
        env: # --------------------------------- ENV variables
        - name: APP_COLOR
          value: pink
```

### ConfigMap

```YAML
...
        env: # --------------------------------- ENV variables
        - name: APP_COLOR
          valueFrom:
            configMapKeyRef: 
```

### Secrets

```YAML
...
        env: # --------------------------------- ENV variables
        - name: APP_COLOR
          valueFrom:
            secret:KeyRef:
```
