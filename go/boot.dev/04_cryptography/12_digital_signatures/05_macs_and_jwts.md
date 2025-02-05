# MACs and JWTs

You might be thinking that it's inconvenient to only be able to validate messages with people who share a copy of the secret key. However, this is how many web applications handle authentication, including Boot.dev!

## JWTs for User Authentication

A JWT, or JSON Web Token, is just an HMAC where the message data is a JSON object. For example:

```go
message = {"userID": "11be9160-2243-4449-934b-e8245fe2feb0"}
hmac = sha256(key1 + sha256(key2 + message))

jwt = message + "." + hmac
```

JWTs that use HMACs are useful for authentication because it's the same server that's issuing and validating the tokens.

I use the golang-jwt package in Go quite often, it's a great library for creating and validating JWTs.