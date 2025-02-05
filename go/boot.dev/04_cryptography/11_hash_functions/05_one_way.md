# One Way

Hash functions are one-way functions. Assuming the hash function is a good one, it should be impossible to obtain the input given only the output. The only way to find a message that produces a given hash is to attempt a brute-force search of all possible inputs to see if they produce a matching hash.

A critical operation in most hash functions is the modulo operation (%).

## Modulo = Remainder After Division

```txt
16 % 7 = 2
```

Notice how the modulo operation is irreversible. If I only give you the output of `x % 7`, (which, when `x=16` produces `2`) you can't know that `x` was `16`. It could have been `9`, `23`, `30`, etc.
