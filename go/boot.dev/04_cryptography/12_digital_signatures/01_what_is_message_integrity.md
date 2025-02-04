# What Is a Message Integrity?

Message integrity ensures that a message has not been changed.

Let's say you need to transmit a private key from your computer to your phone. Because your network connection isn't great, there's a chance that information is lost in transit. You want to be sure that the private key makes it to your phone in pristine condition. If there is data loss, you'll want to know about it so you can try again.

To accomplish this, we can take a hash of the private key, and send it alongside the key. When the private key is transmitted, you'll be able to check on your phone by hashing the private key and checking if the hashes match.

## Checksums

This "hash and check" method is called a checksum.

1. checksum = `hash(data)`
2. send data
3. send checksum
4. if `(hash(data) == checksum)` then the `message` has not been altered

## Assignment

Complete the `checksumMatches` function.

It should hash the `message` after converting it directly to a slice of bytes using sha256. Check if the lowercase hexadecimal encoding of the hash matches the `checksum` argument. If it does, return `true`. Otherwise, return `false`.

### Solution

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func checksumMatches(message string, checksum string) bool {
	h := sha256.New()
	h.Write([]byte(message))
	hex := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Printf(hex)
	return hex == checksum
}
```

