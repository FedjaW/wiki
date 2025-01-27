# Formatting

As we've discussed, most cryptographic algorithms operate directly on bits of data. That is, they operate on strings of 0's and 1's.

For human-readability purposes, it is useful to know how to encode the data in a friendlier format like Hexadecimal. When working with binary data, it's most common to use:

- Base 2 (Binary)
- Base 10 (Decimal)
- Base 16 (Hexadecimal)

## Encoding in Go

To encode `byte` or `int` data in Go, we can use the `fmt` package's formatting verbs. Remember, all data in a computer is just a bunch of numbers at the end of the day.

```go
myData := 32

fmt.Printf("%b\n", myData)
// prints binary: 100000

fmt.Printf("%d\n", myData)
// prints decimal: 32

fmt.Printf("%x\n", myData)
// prints hexadecimal: 20
```

## Padding With Zeroes

When encoding data in binary, it's common to pad the data with zeroes to make it a certain length. For example, if we want to encode the number 32 in binary, we can pad it with zeroes to make it 8 bits long:

```go
myData := 32

fmt.Printf("%08b\n", myData)
// prints binary: 00100000

fmt.Printf("%02x\n", myData)
// prints hexadecimal: 20
```

## Assignment

It's common for our customers to want to audit their encrypted and encoded vault data, just to make sure Passly is doing its job. We're developing some tooling to make that easier for them.

Complete the getHexString and getBinaryString functions. They accept a []byte and return a string of the hexadecimal and binary representations of the data respectively.

Pad the strings with zeroes so that they are always the maximum length of a byte. Keep in mind, the hex representation of a byte will be much shorter than the binary version due to its larger alphabet. Additionally, place a colon as a delimiter between each byte. For example:

```
byte1:byte2:byte3:byte4:etc
```

### Note

The test suite casts a string to a []byte and passes it to your function. Your function should return the hex/binary representation of the raw data.

Also, be aware of the strings.Join() function. It's a great way to join a slice of strings together with a delimiter.

### Solution

```go
package main

import (
	"fmt"
	"strings"
)

func getHexString(b []byte) string {
	result := []string{}
	for _, v := range b {
		result = append(result, fmt.Sprintf("%02x", v))
	}
	return strings.Join(result, ":")
}

func getBinaryString(b []byte) string {
	result := []string{}
	for _, v := range b {
		result = append(result, fmt.Sprintf("%08b", v)) // Binary representation, padded to 8 bits
	}
	return strings.Join(result, ":") // Join with colon delimiter
}
```