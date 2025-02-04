# Encryption and Decryption Explained

## tl;dr

Due to these properties:

- Prime numbers / prime factors
- Modular multiplicative inverses
- Big freaking numbers

RSA is a secure algorithm that allows us to encrypt and decrypt messages.

## Encryption Formula

```text
c = m^e (mod n)
```

- `c` is the ciphertext
- `m` is the message
- `e` is the public key exponent
- `n` is the public key modulus

Why Is "c" a Secure Ciphertext?

It's secure because it's not possible to get the original message `m` back from `c`. There is no mathematical inverse of this formula without knowing the private key `d`.

## Decryption Formula

```text
m = c^d (mod n)
```

- `m` is the message
- `c` is the ciphertext
- `d` is the private key exponent
- `n` is the public key modulus

## The Decryption Math Explained in Excruciating Detail

First, let's remember how `d` was generated: it's the modular multiplicative inverse of `e` in "mod phi":

```txt
ed ≡ 1 (mod ϕ(n))
```

Which, according to what we learned about these equations being equivalent:

```txt
n ≡ r (mod q)
n = qk + r
```

Can be converted to this equation for the key:

```txt
ed = ϕ(n)k + 1
```

Next, remember that the equation for decryption is:

```txt
c^d (mod n)
```

Which, through substitution of the encryption formula, is the same as:

```txt
(m^e)^d (mod n)
```

Which is the same as:

```txt
m^(ed) (mod n)
```

So we can plug the key equation into the ed in the decryption equation:

```txt
mϕ(n)k+1 (mod n)
```

Which can be rewritten as:

```txt
(mϕ(n))k * m (mod n)
```

Euler's Theorem tells us that:

```txt
m^ϕ(n) ≡ 1 (mod n)
```

So we can rewrite the equation as:

```txt
1^k * m (mod n)
```

Which is just:

```txt
m (mod n)
```

Which is the original unencrypted message `m!`

## Further Reading

Here's another good reference if you're curious to read more.
<https://kndrck.co/posts/zk_rsa_explained_with_examples/#mjx-eqn-eq%3Adecrypted_message_formula>