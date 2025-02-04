# Modular Arithmetic

Modular arithmetic and modular exponentiation are widely used in cryptography, and RSA is no exception.

## (Mod N)

In modular arithmetic, numbers "wrap around" when reaching a certain value, called the "modulus". The calculation is simple.

`a = b (mod n)`

Using the modulo operator `%`, we first do `a % n = b`. Now we know number `a` is congruent with the remainder, `b`, in `(mod n)`. Likewise, all numbers with this remainder are congruent in `(mod n)`.

### (Mod 2)

In the world of `(mod 2)`, `0` is congruent with `0 (mod 2)`. So we would write:

```txt
0 ≡ 0 (mod 2)
```

And so on:

- 0 ≡ 0 (mod 2)
- 1 ≡ 1 (mod 2)
- 2 ≡ 0 (mod 2) wraps back to zero!
- 3 ≡ 1 (mod 2)
- 4 ≡ 0 (mod 2) wraps back to zero!
- ...

Note: `≡` is the congruence symbol.
see: <https://www.aimath.org/news/congruentnumbers/modulo.html>

### (Mod 3)

As you saw above, in the world of (mod 2), there are only 2 possible values: {0, 1}. However, in the world of (mod 3), there are 3 possible values: {0, 1, 2}.

- 0 ≡ 0 (mod 3)
- 1 ≡ 1 (mod 3)
- 2 ≡ 2 (mod 3)
- 3 ≡ 0 (mod 3) wraps back to zero!
- 4 ≡ 1 (mod 3)
- 5 ≡ 2 (mod 3)
- 6 ≡ 0 (mod 3) wraps back to zero!
- ...

## More

It will be important to understand a simple transformation that's possible with modular arithmetic. These 2 formulas are effectively equivalent:

```txt
n ≡ r (mod q)
n = qk + r
```

Where n is the number we're working with, r is the remainder, q is the quotient, and k is the number of times q goes into n. This allows us to convert an equation in the world of modular arithmetic into a more familiar equation in the normal world.

For example,

```txt
9 ≡ 1 (mod 2)
```

Can be converted to:

```txt
9 = 2k + 1
```
