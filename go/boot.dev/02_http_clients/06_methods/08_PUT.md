# HTTP PUT

The HTTP PUT method creates or more commonly, updates the target resource with the contents of the request's body. Unlike GET and POST, there is no http.Put function. You will have to create a raw http.Request that an http.Client can Do.

## POST vs PUT

While POST and PUT are both used for creating resources, PUT is more common for updates and is idempotent, meaning it's safe to make the request multiple times without changing the server state. For example:

```txt
POST /users/bob (create bob user)
POST /users/bob (create duplicate bob user)
POST /users/bob (create duplicate bob user)
```

```txt
PUT /users/bob (create bob user if it doesn't exist)
PUT /users/bob (update bob user with the same data)
PUT /users/bob (update bob user with the same data)
```