# Order 1

`O(1)` means that no matter the size of the input, there is no growth in the runtime of the algorithm. This is also referred to as a "constant time" algorithm.

In Python, a dictionary offers the ability to look items up by key, which is an operation that is independent of the size of the dictionary:

```py
# this is a constant time lookup
org = organizations[org_id]
```

Dictionary lookups are `O(1)`. Which is one of the reasons dictionaries and dictionary-equivalents in other languages are used all over the place.
