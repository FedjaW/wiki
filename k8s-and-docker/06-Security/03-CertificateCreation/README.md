# Certificate creation

There are many tools to create certificates. We will use `OPENSSL`.

## CERTIFICATE AUTHORITY (CA) - certificate

To create a certificate:

```CLI
openssl genrsa -out ca.key 2048
```

Generate a certificate signing request:

```CLI
openssl req -new -key ca.key -subj "/CN=KUBERNETES-CA" -out ca.csr
```

Sign the certificate:

```CLI
openssl x509 -req -in ca.csr -signkey ca.key -out ca.crt
```

## ADMIN USER - certificate

To create a certificate:

```CLI
openssl genrsa -out admin.key 2048
```

Generate a certificate signing request:

```CLI
openssl req -new -key admin.key -subj "/CN=kube-admin/O=system:masters" -out admin.csr
```

Sign the certificate:

```CLI
openssl x509 -req -in admin.csr -CA ca.crt -CAkey ca.key -out admin.crt
```

## Same procedure for other certifactes...