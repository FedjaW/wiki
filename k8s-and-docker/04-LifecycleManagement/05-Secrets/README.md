# Secrets

Secrets are used to store sensitive information like passwords or keys.

## Create Secret

Create a secret (imperative):

```CLI
kubectl create secret generic app-secret --from-literal=DB_HOST=mysql
```

or via file (imperative):

```CLI
kubectl create secret generic app-secret --from-file=app_secret.properties
```

Declarativ:

```YAML
apiVersion: v1
kind: Secret
metadata:
  name: app-secret
data:
  DB_HOST: mysql # -> bXlzcWw=
  DB_USER: root # -> ...
  DB_PASSWORD: paswrd # -> ...
```

Secrets must be provided in an encoded format. To encode on linux:

```CLI
echo -n 'mysql' | base 64
```

To list the secrets:

```CLI
kubectl get secrets
```

To see the encoded values:

```CLI
kubectl get secret app-secret -o yaml
```

To decode:

```CLI
echo -n 'bXlzcWw=' | base 64 --decode
```

## Inject Secret into the pod

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
        - secretRef:
            name: app-secret # name of the secret created earlier
```

## Notes

Secrets are not encrypted! They are only encoded!
