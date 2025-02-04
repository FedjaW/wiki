# Message Authentication Code (MAC)

While checksums are useful on their own, they provide no authentication mechanism. For example, if the NSA intercepts a message in transit, they can maliciously alter the message and the checksum. When the message is received by the original recipient, they will think it hasn't been altered.

All a plain checksum does is protect against accidental alterations.

To get the protection we want from malicious attackers, we need to use a MAC, or Message Authentication Code

In short, a MAC is a checksum that is authenticated by a secret key. This means that if the message is altered by someone who doesn't have the secret key, the checksum will not match.

![img](./02_mac.png)

A MAC takes advantage of the irreversibility property of hash functions. A MAC is the hash of a message concatenated with a key.

```txt
mac = hash(message + key)
```

As with a checksum, a mac can be sent along with the message in a communication. A man in the middle won't be able to alter the message secretly, because they can't produce a new valid `mac` without access to the secret key.

1. mac = `hash(message + key)`
2. send mac
3. send message
4. if `mac == hash(message + key)` then the message is valid

The disadvantage of a MAC is that the receiver will also need a copy of the secret key. This is why MACs are generally only used when the sender and the receiver are the same entity.

## Assignment

Fix the `macMatches` function. It currently uses a plain checksum. Concatenate the key to the end of the message before hashing it.

### Solution

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func macMatches(message, key, checksum string) bool {
	h := sha256.New()
	h.Write([]byte(message + key))
	return checksum == fmt.Sprintf("%x", h.Sum(nil))
}
```