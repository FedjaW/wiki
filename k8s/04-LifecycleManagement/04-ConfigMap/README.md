# ConfigMaps

ConfigMaps are injected into the container as environment variables.

Two steps are needed:

__1. Creating ConfigMap__

`Imperative:`

- Direct

```CLI
kubectl create configmap app-config --from-literal=APP_COLOR=blue
```

- From file

```CLI
kubectl create configmap app-config --from-file=app_config.poperties
```

`Declarative:`

```CLI
kubectl create -f configmap-definition.yaml
```

The configmap-definition.yaml:

```YAML
apiVersion: v1
kind:
metadata:
  name: app-config # (*)
data:
  APP_COLOR: blue
  APP_MODE: prod
```

__2. Inject into the pod__

To inject the configmap into the container specify the env variable:

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
      envFrom: # --------------------------------- ENV variables
        configMapRef:
          name: app-config # (*)
```

## ENV, SINGLE ENV, VOLUME

### ENV

```YAML
...
      envFrom: # --------------------------------- ENV variables
        configMapRef:
          name: app-config # (*)
```

## SINGLE ENV

```YAML
...
        env: # --------------------------------- ENV variables
        - name: APP_COLOR
          valueFrom:
            configMapKeyRef:
              name: app-config
              key: APP_COLOR 
```

## VOLUME

```YAML
...
        volumes: # --------------------------------- Volume
        - name: app-config-volume
          configMap:
            name: app-config
```
