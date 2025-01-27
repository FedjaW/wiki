# Decoding

We've encoded some raw data into hex and binary strings, now let's write some code to decode the hexadecimal version.

## Assignment

Complete the `getHexBytes` function. It takes a string of the following format as input:

```
48:65:6c:6c:6f
```

It should return a `[]byte` that represents the `string` encoded in each hex value. We're essentially reversing the process we used to encode the data in the last assignment.

The following functions will help:

- strings.Split()
- hex.DecodeString()

### Solution

```go
package main

import (
	"encoding/hex"
	"strings"
)

func getHexBytes(s string) ([]byte, error) {
	final := []byte{}
	parts := strings.Split(s, ":")
	for _, v := range parts {
		dec, err := hex.DecodeString(v)
		if err != nil {
			return nil, err
		}
        // (three dots) dec... is the variadic argument operator
		final = append(final, dec...)
	}
	return final, nil
}
```
