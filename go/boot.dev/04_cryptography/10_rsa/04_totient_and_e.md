# Totient and E

Now that we have p, q and n, we need to calculate:

- The totient of n, which we'll call tot.
- A random number that is relatively prime to tot, which we'll call e

The public key for RSA encryption is the pair of numbers: (n, e)

## The Totient Function (Phi)

Euler's totient function counts the positive integers up to a given integer, n in our case, that are relatively prime to it.

In other words, the totient is the number of integers between 2 and n whose greatest common divisor is 1.

`tot = ϕ(n) = (p - 1) * (q - 1)`

Remember `p` and `q` are prime, which means their totient's are just `p-1` and `q-1`. Because we know `n = p * q`, we know the totient of n is `(p - 1) * (q - 1)`.

## e

`e` is a random number between `1` and `tot` that is relatively prime to `tot`. This means that the greatest common divisor of `e` and `tot` is `1`.

## Assignment

Complete the getTot and getE functions.

### getTot(p, q *big.Int) *big.Int

Use the math/big package to calculate `(p-1)(q-1)` and return it as a pointer to a big.Int. This is the "totient" of `n`, which we can also call "phi of n", or `ϕ(n)`.

### getE(tot *big.Int) *big.Int

Use the math/big package to generate a random number e that adheres to the following constraints:

- `e` is greater than `1`
- `e` is less than `tot`
- `e` and `tot` have a greatest common divisor of `1`

The `gcd` function is provided for you. It calculates the greatest common divisor of two big ints.

Generate random `e` values in the range of `[2, tot)` until you find one that satisfies the constraints. Use `crand.Int` with the globally provided `randReader` to generate random big ints. You'll need to do some manual arithmetic to get the range you want because `crand.Int` only generates random numbers in the range of `[0, max)`

### Solution

```go
package main

import (
	crand "crypto/rand"
	"math/big"
)

func getTot(p, q *big.Int) *big.Int {
	tot := new(big.Int)
	tot.Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	return tot
}

func getE(tot *big.Int) *big.Int {
	totMinusTwo := new(big.Int).Sub(tot, big.NewInt(2))

	e, _ := crand.Int(randReader, totMinusTwo)
	e.Add(e, big.NewInt(2))

	for gcd(e, tot).Cmp(big.NewInt(2)) != 0 {
		e, _ = crand.Int(randReader, totMinusTwo)
		e.Add(e, big.NewInt(2))
	}
	return e
}
```

// helpers.go


```go
package main

import (
	"errors"
	"fmt"
	"io"
	"math/big"
	mrand "math/rand"
)

func generatePrivateNums(keysize int) (*big.Int, *big.Int) {
	p, _ := getBigPrime(keysize)
	q, _ := getBigPrime(keysize)
	return p, q
}

func gcd(x, y *big.Int) *big.Int {
	xCopy := new(big.Int).Set(x)
	yCopy := new(big.Int).Set(y)
	for yCopy.Cmp(big.NewInt(0)) != 0 {
		xCopy, yCopy = yCopy, xCopy.Mod(xCopy, yCopy)
	}
	return xCopy
}

func firstNDigits(n big.Int, numDigits int) string {
	if len(n.String()) < numDigits {
		return fmt.Sprintf("%v", n.String())
	}
	return fmt.Sprintf("%v...", n.String()[:numDigits])
}

var randReader = mrand.New(mrand.NewSource(0))

func getBigPrime(bits int) (*big.Int, error) {
	if bits < 2 {
		return nil, errors.New("prime size must be at least 2-bit")
	}
	b := uint(bits % 8)
	if b == 0 {
		b = 8
	}
	bytes := make([]byte, (bits+7)/8)
	p := new(big.Int)
	for {
		if _, err := io.ReadFull(randReader, bytes); err != nil {
			return nil, err
		}
		bytes[0] &= uint8(int(1<<b) - 1)
		if b >= 2 {
			bytes[0] |= 3 << (b - 2)
		} else {
			bytes[0] |= 1
			if len(bytes) > 1 {
				bytes[1] |= 0x80
			}
		}
		bytes[len(bytes)-1] |= 1
		p.SetBytes(bytes)
		if p.ProbablyPrime(20) {
			return p, nil
		}
	}
}
```