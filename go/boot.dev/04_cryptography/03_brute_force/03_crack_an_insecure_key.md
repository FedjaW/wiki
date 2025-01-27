# Crack an Insecure Key

At Passly we practice Cryptanalysis, meaning that we try to break our own encryption algorithms. We do this to make sure that our encryption is secure!

## Assignment

Complete the findKey function. It accepts the encrypted and decrypted versions of messages as arguments. It should return the key that was used to encrypt the password.

If a key is found, return it and a nil error. Otherwise, return an empty byte slice and an error that reads key not found.

### Steps

1. Iterate over all the numbers from `0` to `2^24`.
2. Convert each number to a byte slice (the next key to try) using the `intToBytes` function.
3. Decrypt the `encrypted` message using the current key and the `crypt` function.
4. If the decrypted value matches the expected decrypted value (which was passed in as an argument), you found the key! (You can cast the byte slice directly to a string to compare it to the expected value)

Keep in mind, this code might take a while to run, after all, you're brute forcing 24 bits of data!

### Solution

```go
package main

import (
	"fmt"
	"math"
)

func findKey(encrypted []byte, decrypted string) ([]byte, error) {
	for i := 0; i < int(math.Pow(2, 24)); i++ {
		potentialKey := intToBytes(i)
		decryptedMessage := crypt(encrypted, potentialKey)
		if string(decryptedMessage) == decrypted {
			return potentialKey, nil
		}
	}
	return nil, fmt.Errorf("key not found")
}
```


#### Helper Functions

```go
package main

import (
	"bytes"
	"encoding/binary"
)

// Helper function: crypt performs XOR-based encryption/decryption
func crypt(dat, key []byte) []byte {
	final := []byte{}
	for i, d := range dat {
		final = append(final, d^key[i])
	}
	return final
}

// Helper function: intToBytes converts an integer to a 3-byte slice (little-endian)
func intToBytes(num int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int64(num))
	if err != nil {
		return nil
	}
	bs := buf.Bytes()
	if len(bs) > 3 {
		return bs[:3]
	}
	return bs
}
```

#### Q & A

`Q:` why is intToBytes not just: `[]bytes(number)`?

`A:`

Let me explain why we need the special `intToBytes` helper function instead of a simple type conversion.

In Go, you can't directly convert an integer to `[]byte` using `[]byte(number)`. That syntax only works for strings. If you tried `[]byte(42)`, the compiler would give you an error.

The `intToBytes` function does something specific:

1. It converts the integer to exactly 3 bytes
2. It uses little-endian byte ordering
3. It handles the binary representation of the number

Think of it like this: if you have the number 300, its binary representation would be:

```
300 (decimal) = 100101100 (binary)
```

The `intToBytes` function will store this across 3 bytes, padding with zeros as needed.

`Q:` Why three bytes for 2^24 numbers?

`A:`

Let's break it down:

- 1 byte = 8 bits
- 3 bytes = 24 bits

With 24 bits, you can represent 2^24 different numbers (from 0 to 2^24 - 1). That's about 16.7 million possibilities.

In this cryptanalysis problem, we're trying all possible 3-byte keys. Why 3 bytes? Because if you look at the encrypted messages in the test cases, they're all 3 bytes long:

```go
{[]byte{0x1b, 0x2c, 0x3d}, "yes", ...}  // 3 bytes encrypted
{[]byte{0x2a, 0xff, 0xea}, "car", ...}  // 3 bytes encrypted
```

Since the XOR operation (used in the crypt function) needs the key to be the same length as the message, we need a 3-byte key to decrypt a 3-byte message.

`Q:` err := binary.Write(buf, binary.LittleEndian, int64(num)) Why cast to int64?

`A:`

The binary.Write function requires a fixed-size number type for consistent binary representation across different systems.

The regular int type in Go is platform-dependent - it could be 32 bits on some systems and 64 bits on others. By explicitly casting to int64, we ensure consistent behavior regardless of the platform.

Think of it like this: if we passed a regular int:

- On a 32-bit system: int is 4 bytes
- On a 64-bit system: int is 8 bytes

By using int64, we're saying "always use 8 bytes" before we trim it down to 3 bytes with:

```go
if len(bs) > 3 {
    return bs[:3]
}
```