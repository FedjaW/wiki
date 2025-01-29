# Key Schedules

The simplest versions of block ciphers operate on single blocks of data, however, those blocks of data typically go through many rounds of encryption. Each round of encryption needs its own key, but it's nearly impossible for users to keep track of more than a single key.

A key schedule is an algorithm that a block cipher employs to split an original key into multiple "round keys" or "subkeys". These round keys are deterministically derived from the original key, meaning that the same original key will always produce the same round keys.

For example the following "master key":

```txt
key = 1111011101010011
```

might produce the following 8 "round keys" in an 8-round block cipher:

```txt
roundKey0 = 1110001001111010
roundKey1 = 0100101101011100
roundKey2 = 0011110110101110
roundKey3 = 0101101011100010
roundKey4 = 1011100010010110
roundKey5 = 0010010110101110
roundKey6 = 0101011100010010
roundKey7 = 1101011111110010
```

## Assignment

In the real world, there are many different production algorithms for key schedules. At Passly, we've been asked to implement a simple key schedule that can be used in our test suite. It doesn't need to be "production-level-secure", it just needs to get the high-level idea across.

Complete the `deriveRoundKey()` function. It accepts a master key, which is `4` bytes long and represented as an array of bytes. It will also accept a "round number", which is just an `int` representing which round key is being derived. The key schedule produces a round key where each byte in the round key is the original byte from the master key `XOR`ed with the binary representation of the round number.

For example:

```txt
masterKey     = 01101100 01110000...
roundNumber   = 00000001
roundKey      = 01101101 01110001...
```

or when `roundNumber = 5`

```txt
masterKey     = 01101100 01110000...
roundNumber   = 00000101
roundKey      = 01101001 01110101...
```

### Solution

```go
package main

func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
	for i, v := range masterKey {
		masterKey[i] = v ^ byte(roundNumber)
	}
	return masterKey
}
```

#### Q&A

`Q:` Why can I do `byte(roundNumber)` to cast an `int` to a `byte`

`A:`

n Go, a `byte` is just an alias for `uint8`, which means it's an 8-bit unsigned integer that can hold values from 0 to 255. When you use `byte(roundNumber)`, you're telling Go to convert the `int` value to fit within a single byte.

Let's break down what happens:

1. An `int` in Go is typically 32 or 64 bits depending on your system
2. When converting to `byte`, Go takes just the least significant 8 bits
3. Any bits beyond the first 8 are truncated

For example, if we had a number like:

```go
roundNumber := 5  // In binary: 00000000 00000000 00000000 00000101
b := byte(roundNumber) // Results in: 00000101
```

### Key Schedule Review

A key schedule is just an algorithm that splits a master key into multiple "round keys" or "sub keys". This ensures that each round in a block cipher can use its own unique key for that round of encryption, increasing the security of the cipher.

Key schedules must be deterministic - there can be no randomness in the derivation of the round keys.