# Hash Functions

Hash functions have 3 basic goals:

- Hash functions scramble data deterministically
- No matter the input, the output of a hash function always has the same size
- The original data can not be retrieved from the scrambled data (one-way function)

## SHA-256

SHA-256 is one of the most widely used hash functions. It's fast, secure, and does a great job meeting the 3 goals stated above.

## Checking Integrity

At Passly, we use SHA-256 for many things, but one of the most important is to ensure the integrity of a password vault. Each time we save a vault, we hash the vault's contents and store the hash in our database. Later, if we need to verify the integrity of the vault, we can hash the vault's contents again and compare the two hashes. If they match, we know the vault has not been tampered with.

## Assignment

Add the following functions and methods to the program:

- `newHasher`
- `h.Write`
- `h.GetHex`

### newHasher

Returns a pointer to a new `hasher`. Uses `sha256.New()` to create a new `hash.Hash`.

### h.Write

A method on a pointer to a `hasher`. Uses `hash.Write()` to write data to the hasher. It should accept a string and cast the string to a `[]byte`. It should pass along the return values, that is, it returns the number of bytes written from p `(0 <= n <= len(p))` and any error encountered that caused the write to stop early.

### h.GetHex

A method on a pointer to a `hasher`. Uses `hash.Sum()` to get the hash value of the data written to the hasher. It should encode the hash value as a lowercase hex string and return it.

### Solution

```go
package main

import (
	"crypto/sha256"
	"fmt"
	"hash"
)

type hasher struct {
	hash hash.Hash
}

func newHasher() *hasher {
	return &hasher{
		hash: sha256.New(),
	}
}

func (h *hasher) Write(s string) (n int, err error) {
	return h.hash.Write([]byte(s))
}

func (h *hasher) GetHex() string {
	return fmt.Sprintf("%x", h.hash.Sum(nil))
}
```