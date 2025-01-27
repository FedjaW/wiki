# Better Keys

Keys are typically encoded in software as raw binary data. In Go, this means we usually use a slice of bytes as the data type. For example:

```go
[]byte{0xAA, 0xBB}
```

`0xAA` and `0xBB` are hexadecimal values, and hex is one of the more common ways to write raw binary in code. They represent the raw binary data `10101010` and `10111011` respectively.

## Code

Create a cipher. Here are the docs for aes.NewCipher(): <https://pkg.go.dev/crypto/aes#NewCipher>

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func keyToCipher(key string) (cipher.Block, error) {
	keyEncoded := []byte(key)
	return aes.NewCipher(keyEncoded)
}
```
