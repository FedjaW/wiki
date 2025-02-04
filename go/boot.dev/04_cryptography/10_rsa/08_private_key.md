# Private Key

While `(n, e)` is the public key, `d` is the private key. That said, `d` can be derived from `p` and `q`, so you can also consider the pair of `p` and `q` the private key, but it's just different ways to look at it.

At the end of the day, `p`, `q`, tot and `d` should all be kept secret, there's no reason to share anything other than `(n, e)`.

## Generating D

`d` is simply the modular multiplicative inverse of `e` `(mod tot)`.

## Assignment

Complete the `getD` function.

### getD(e, tot *big.Int) *big.Int

Use the ModInverse method to calculate `d` and return it.

### Solution

```go
package main

import (
	"math/big"
)

// Get the private exponent
func getD(e, tot *big.Int) *big.Int {
	d := new(big.Int)
	d.ModInverse(e, tot)
	return d
}
```