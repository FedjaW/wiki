# Args and kwargs

In Python, *args and **kwargs allow a function to accept and deal with a variable number of arguments.

- *args collects positional arguments into a tuple
- **kwargs collects keyword (named) arguments into a dictionary

```py
def print_arguments(*args, **kwargs):
    print(f"Positional arguments: {args}")
    print(f"Keyword arguments: {kwargs}")

print_arguments("hello", "world", a=1, b=2)
# Positional arguments: ('hello', 'world')
# Keyword arguments: {'a': 1, 'b': 2}
```

## Positional arguments

Positional arguments are the ones you're already familiar with, where the order of the arguments matters. Like this:

```py
def sub(a, b):
    return a - b

# a=3, b=2
res = sub(3, 2)
# res = 1
```

## Keyword arguments

Keyword arguments are passed in by name. Order does not matter. Like this:

```py
def sub(a, b):
    return a - b

res = sub(b=3, a=2)
# res = -1
res = sub(a=3, b=2)
# res = 1
```

## A note on ordering

Any positional arguments must come before keyword arguments. This will not work:

```py
sub(b=3, 2)
```