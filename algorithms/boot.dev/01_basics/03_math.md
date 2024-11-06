# Math

## Exponents

I promise we'll get to how this relates to coding, but first we need to review some math stuff.

An exponent indicates how many times a number is to be multiplied by itself.

For example:

`53 = 5 * 5 * 5 = 125`

Sometimes exponents are also written using the caret symbol `(^)`:

`5^3 = 53`

### Exponent syntax

The ** operator calculates an exponent in Python.

```py
square = 2 ** 2
# square = 4

cube = 2 ** 3
# cube = 8
```

### Vocabulary

`5^3 = 53`

53 is read as "5 to the power of 3"
5 is the "base"
3 is the "exponent"

## Logarithms

A logarithm is the inverse of an exponent.

`24 = 16`

`log216 = 4`

"log216" can be read as "log base 2 of 16", and means "the number of times 2 must be multiplied by itself to equal 16".

"log216" might also be written as `log2(16)`

### Python Syntax

There isn't a language-level operator to calculate a logarithm, but we can import the math library and use the math.log() function.

```py
import math

print(f"Logarithm base 2 of 16 is: {math.log(16, 2)}")
# Logarithm base 2 of 16 is: 4.0
```

## What is nice for algorithms

Exponents grow very quickly, and logarithms grow very slowly. A logarithm is the inverse of an exponent.

Generally speaking, it's nice when we can write code that uses `log(n)` time to run, where `n` is the amount of data to process. For example, let's say we have a list of `1,000,000` users, and we want to write an algorithm that finds the user with the most followers.

If it takes `1` millisecond to check one user (let's just pretend), a linear algorithm would take `1,000,000` milliseconds, or about `16` minutes and `40` seconds.

A quadratic algorithm (exponent) would take `1,000,000,000,000` milliseconds, or about 31.7 years.

## Factorials

The factorial of a positive integer `n`, written `n!`, is the product of all positive integers less than and equal to `n`.

`5! = 5 * 4 * 3 * 2 * 1 = 120`

The results of a factorial grow *even faster* than exponentiation!
