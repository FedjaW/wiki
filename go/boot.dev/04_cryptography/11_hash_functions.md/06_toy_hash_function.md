# Toy Hash Function

Let's build a toy hash function!

Our goal will be to build a function that:

- Is hard to reverse
- Has a fixed output size
- Is deterministic

Our function will not be suitable for production use, but will be useful for the Passly marketing team to explain how our security systems work from a high level.

## Shifting Bits in Go

To shift a byte left by numBits:

```go
numBits := 3
// original = 11110000
shifted := original << numBits
// shifted = 10000000
```

## XOR in Go

```go
result := a ^ b
```

## Assignment

Complete the `hash()` function. It takes an arbitrarily sized `[]byte` and returns a fixed size `[4]byte`.

It should do the following:

1. Rotate the bits in each byte left 3 bits, do this one byte at a time
2. Shift the bits in each byte left 2 bits, do this one byte at a time
3. Create a new empty "final" array of exactly 4 bytes
4. XOR each byte with its corresponding byte in the final array, and save the result back to the final array. Do this one byte at a time for each byte in the rotated/shifted input. The "corresponding" byte in the final array is at index `i%4`, where `i` is the index of the byte in the input.

### Solution

```go
package main

import (
	"math/bits"
)

func hash(input []byte) [4]byte {
	for i, b := range input {
		rotated := bits.RotateLeft8(uint8(b), 3)
		input[i] = byte(rotated)
	}

	for i, b := range input {
		input[i] = b << 2
	}

	final := [4]byte{}
	for i, b := range input {
		// final[i%4] = final[i%4] XOR b
		final[i%4] = final[i%4] ^ b
	}

	return final
}
```