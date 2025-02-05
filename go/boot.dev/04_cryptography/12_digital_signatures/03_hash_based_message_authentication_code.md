# Hash-Based Message Authentication Code

An HMAC is a kind of MAC. All HMACs are MACs but not all MACs are HMACs. The main difference is that an HMAC uses two rounds of hashing instead of one (or none). Each round of hashing uses a child key that's derived from the secretKey.

Here's a naive implementation:

```go
secretKey = 'thisIsASecretKey1234'
childKey1 = 'thisIsASe'
childKey2 = 'cretKey1234'
hash(childKey1 + hash(childKey2 + 'the message we want to send'))
```

This is a simplified version of the function given in RFC-2104.

## Why Use HMAC? Why Do We Need to Hash Twice?

With some MACs, depending on the hash function, it is possible to change the message (without knowing the key) and obtain another valid MAC. This is called a length extension attack. There are no known extension attacks against the current HMAC specification, so you should prefer HMACs over MACs.

## Note About the Child Keys

In the naive example above, we created 2 child keys by splitting the original key. That's not the way a secure HMAC would be implemented, instead, we would derive child keys using a slightly more complex process. That said, the principle is the same: using a single key we can derive two separate child keys.

If you're curious about how that might work in production, you can check out the implementation here.
