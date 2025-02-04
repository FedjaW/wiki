# Multiplicative Inverse

The multiplicative inverse of a number, say `x`, is written as `x^-1` or `1/x`.

It's a number that when multiplied by `x` equals `1`.

## Modular Multiplicative Inverse

In the modular world, a multiplicative inverse of `a` is `x` as defined by this formula:

```txt
a * x ≡ 1 (mod m)
```

For example, an inverse of `3` in "mod 11" is `4` because:

```txt
3 * 4 ≡ 1 (mod 11)
```

As another example, an inverse of `5` in "mod 11" is `9` because:

```txt
5 * 9 ≡ 1 (mod 11)
```

Some numbers don't have a multiplicative inverse in a given mod. For example, `2` has no multiplicative inverse in "mod 4" because, in the following equation, there is no `x` that will result in a congruence:

```txt
2 * x ≡ 1 (mod 4)
```

### Who Cares?

This will be important later when we talk about decryption, but it's important to understand that this formula:

```txt
y^ax (mod m)
```

can be reduced to:

```txt
y^1 (mod m)
```

which reduces to:

```txt
y (mod m)
```

Of course, this all assumes that `a` is the modular multiplicative inverse of `x` in "mod m".