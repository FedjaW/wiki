# Immutability

In FP, we strive to make data immutable. Once a value is created, it cannot be changed. Mutable data, on the other hand, can be changed after it's created.

## Who cares?

Immutable data is easier to think about and work with. When 10 different functions have access to the same variable, and you're debugging a problem with that variable, you have to consider the possibility that any of those functions could have changed the value.

When a variable is immutable, you can be sure that it hasn't changed since it was created. It's a helluva lot easier to work with.

*Generally speaking, immutability means fewer bugs and more maintainable code.*

## Tuples vs Lists

Tuples and lists are both ordered collections of values, but tuples are immutable and lists are mutable.

You can append to a list, but you can not append to a tuple. You can create a new copy of a tuple using values from an existing tuple, but you can't change the existing tuple.

Lists are mutable

```py
ages = [16, 21, 30]
# 'ages' is being changed in place
ages.append(80)
# [16, 21, 30, 80]
```

Tuples are immutable

```py
ages = (16, 21, 30)
more_ages = (80,) # note the comma! It's required for a single-element tuple
# 'all_ages' is a brand new tuple
all_ages = ages + more_ages
# (16, 21, 30, 80)
```