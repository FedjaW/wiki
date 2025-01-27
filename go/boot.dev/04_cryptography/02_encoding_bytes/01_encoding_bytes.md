# Encoding Bytes

Encoding and encryption are not the same. When we talk about encoding, we are talking about converting raw data (binary) into a format that can be stored or transmitted as text.

There are many popular ways to encode data as text. Some include:

- Hexadecimal
- Base64
- ASCII

Typically, it is a bad idea to create your own encoding scheme. However, for the sake of learning, we are going to create one for fun.

## Alphabets

Encoding schemes have an alphabet. An alphabet is just the set of available characters in the scheme.

### Arabic Numerals Alphabet

`0123456789`

### Hexadecimal Alphabet

`0123456789abcdef`

### Base64 Alphabet

`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`

The number of characters in the alphabet is the "base". Base64 has 64 characters.

## Binary Literals in Go

Binary integer literals are written in Go as follows:

```go
// zeroWrittenInBinary is just an int with the value 0
zeroWrittenInBinary := 0b0000
fmt.Println(zeroWrittenInBinary)
// prints: 0

// oneWrittenInBinary is just an int with the value 1
oneWrittenInBinary := 0b0001
fmt.Println(oneWrittenInBinary)
// prints: 1
```

## Assignment

Complete the `base8Char` function. It accepts a single `byte` of data and returns the associated character from our alphabet: `ABCDEFGH`.

For example, given these binary numbers, it will return the following characters:

```
0000 -> A
0001 -> B
0010 -> C
0011 -> D
etc
```

Note that because we only have `8` characters in our alphabet, we will ignore any numbers equal to or greater than `8`. Remember, an entire byte can contain the numbers `0` through `255`.

### Solution

```go
package main

func base8Char(bits byte) string {
	const base8Alphabet = "ABCDEFGH"
	i := int(bits)
	if i >= len(base8Alphabet) || i < 0 {
		return ""
	}
	return string(base8Alphabet[i])
}
```

## Q & A

`Q:` why do I need to use string() in string(base8Alphabet[i])

`A:`

In Go, when you index into a string like `base8Alphabet[i]`, you actually get back a byte (which is an alias for `uint8`), not a string. This is because strings in Go are just sequences of bytes.

Let me show you with a different example:

```go
alphabet := "XYZ"
firstByte := alphabet[0]
// firstByte is actually the byte/uint8 value 88 (ASCII code for 'X')
```

When we want to convert a single byte back into a string containing that character, we need to use the `string()` conversion. Without it, we'd just be working with the numeric byte value!

Would you like to try an experiment? What do you think would happen if you printed these values:

```go
fmt.Printf("%v\n", "ABC"[0])       // what prints here?
fmt.Printf("%v\n", string("ABC"[0])) // what about here?
```

Result: 

```go
fmt.Printf("%v\n", "ABC"[0])       // prints: 65
fmt.Printf("%v\n", string("ABC"[0])) // prints: A
```

The first line prints `65` because that's the ASCII/UTF-8 numeric value of the character 'A'.
The second line prints `A` because we converted that byte value back into a string.

This clearly shows why we need the `string()` conversion in your base8 encoding function - without it, you'd be returning the numeric value instead of the actual character!
