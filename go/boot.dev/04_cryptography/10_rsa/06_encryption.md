# Encryption

## The Public Key

Aside from the message itself, all we need to perform encryption are the numbers e and n. Together, these are the public key.

The security of RSA relies on the fact that given n, it's really hard to guess what `p` and `q` (the private keys) were. Finding prime factors of very large numbers is computationally expensive. Trying to brute-force `p` and `q` would take more than trillions of years on modern hardware.

Remember, `n = p * q`.

## The Encryption Formula

The math for encrypting a message with RSA follows this formula:

`ciphertext = m^e (mod n)`

- m = message
- e = public key exponent
- n = public key modulus

We'll talk more about this math later, but for now, let's just implement it.

## Assignment

Complete the encrypt function.

### encrypt(m, e, n *big.Int) *big.Int

Return the result of `m^e (mod n)`. Use the .Exp method.

### Solution

```go
package main

import (
	"math/big"
)

func encrypt(m, e, n *big.Int) *big.Int {
	ciphertext := new(big.Int)
	ciphertext.Exp(m, e, n)
	return ciphertext
}
```
