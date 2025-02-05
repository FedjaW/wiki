# Toy HMAC

At Passly, we use HMACs to authenticate messages between our internal servers when they need to make requests to each other over the public internet.

To demonstrate to our Luddite manager why we should use an open-source crypto library instead of writing our own HMAC implementation, we decided to write our own and then prove its inferiority. Hopefully we don't get fired instead of getting our way.

## Assignment

Complete the hmac function. It should:

1. Split the key into two halves. The second half should be the larger half if key's length is odd
2. Return the result of sha256(keyFirstHalf + sha256(keySecondHalf + message)) as a string in lowercase hex

Cast strings directly to slices of bytes and don't use any delimiters when concatenating the data.

### Solution

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func hmac(message, key string) string {
	firstHalf := key[:len(key)/2]
	secondHalf := key[len(key)/2:]

	innerDat := sha256.Sum256([]byte(secondHalf + message))
	allDat := []byte(firstHalf)
	allDat = append(allDat, innerDat[:]...)

	sum := sha256.Sum256(allDat)
	result := fmt.Sprintf("%x", sum)
	return result
}
```

#### My Solution

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func hmac(message, key string) string {
	l := len(key)/2
	first := []byte(key[0:l])
	second := []byte(key[l:])

	h := sha256.New()
	second = append(second, []byte(message)...)
	h.Write(second)
	s := h.Sum(nil)
	
	h2 := sha256.New()
	first = append(first, s...)
	h2.Write(first)

	return fmt.Sprintf("%x", h2.Sum(nil))
}
```

#### Q & A

`Q:` why do I need to use three dots here? `allDat = append(allDat, innerDat[:]...)`

`A:` In Go, when you have innerDat[:], you're creating a slice of bytes. The append() function can take multiple arguments, and the ... operator is used to "unpack" a slice into individual arguments.

Let's break it down:

```go
// Without ..., you're trying to append a slice as a single argument
append(allDat, innerDat[:])    // This won't work!

// With ..., you're appending each element of the slice individually
append(allDat, innerDat[:]...) // This works!
```

Think of it like this: if `innerDat[:]` contains `[1,2,3]`, then:

- Without `...`: You're trying to append `[1,2,3]` as a single element
- With `...`: You're appending `1,2,3` as individual elements

Want to try a small experiment? What do you think would happen if we removed the `...` from the append operation in your code?
