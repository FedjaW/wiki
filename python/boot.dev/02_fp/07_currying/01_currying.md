# Currying

Function currying is a specific kind of function transformation where we translate a single function that accepts multiple arguments into multiple functions that each accept a single argument.

This is a "normal" 3-argument function:

```py
box_volume(3, 4, 5)
```

This is a "curried" series of functions that does the same thing:

```py
box_volume(3)(4)(5)
```