# Asymmetric JWTs

JWTs that use HMACs are symmetric signatures. We talked about how they use the same key to sign and validate the tokens. This works well in some scenarios, but let's take the example of third-party authentication, like "Sign in with Google".

## ECDSA JWTs

Asymmetric JWTs use a private key to sign the JSON payload, and a public key to validate it, it's similar to asymmetric encryption, but instead of encrypting data, we're signing it.

ECDSA, or Elliptic Curve Digital Signature Algorithm, is a type of asymmetric signature algorithm. It works using the same elliptic curve math that we talked about in the chapter on asymmetric encryption.

Lifecycle of an ECDSA (Asymmetric) JWT
Let's take a look at how the Boot.dev web app uses asymmetric JWTs for our "Sign in with Google" button.

1. User clicks "Sign in with Google" and enters their credentials
2. If the credentials are valid, Google creates a JWT by signing it with their private key
3. The signed JWT is given to the client
4. The client sends the JWT in an HTTP header along with every request that requires authentication
5. The Boot.dev server uses Google's public key to validate the JWT in each subsequent request

## Assignment

"Sign in with Google", "Sign in with GitHub", "Sign in with Twitter", etc. are all examples of third-party authentication. They reduce the friction of creating an account on a new website, so many websites use them.

Passly wants to be an issuer of JWTs. Our massive egos demand that we push clients into using "Sign in with Passly".

Complete the `createECDSAMessage` function.

### createECDSAMessage()

1. Hash the message using SHA-256
2. Create a signature of the hashed message using the private key and the SignASN1 function.
3. Return a new token in the following format:

```go
MESSAGE.signature
```

Where `MESSAGE` is the original message, and signature is the signature of the hashed message in lowercase hex.

Keep in mind, this isn't a full JWT, it's an arbitrary message and a signature.

### Solution

```go
package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func createECDSAMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	hash := sha256.Sum256([]byte(message))
	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`%v.%x`, message, sig), nil
}
```