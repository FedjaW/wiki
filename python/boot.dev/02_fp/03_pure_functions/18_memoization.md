# Memoization

At its core, memoization is just caching (storing a copy of) the result of a computation so that we don't have to compute it again in the future.

For example, take this simple function:

```py
def add(x, y):
    return x + y
```

A call to add(5, 7) will always evaluate to 12. So, if you think about it, once we know that add(5, 7) can be replaced with 12, we can just store the value 12 somewhere in memory so that we don't have to do the addition operation again in the future. Then, if we need to add(5, 7) again, we can just look up the value 12 instead of doing a (potentially expensive) CPU operation.

The slower and more complex the function, the more memoization can help speed things up.

Note: It's pronounced "memOization", not "memORization". This confused me for quite a while in college, I thought my professor just didn't speak goodly...
