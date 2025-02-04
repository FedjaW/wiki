# Decryption

Decryption in RSA is almost unbelievably simple from a mathematics perspective but remains extremely secure.

## Assignment

Complete the decrypt function.

### decrypt(c, d, n *big.Int) *big.Int

Use the big.Int.Exp function to raise the encrypted message (`c`) to the power of the private key (`d`) within (mod n).

### Solution

```go
package main

import (
	"math/big"
)

func decrypt(c, d, n *big.Int) *big.Int {
	plainttext := new(big.Int)
	plainttext.Exp(c, d, n)
	return plainttext 
}
```