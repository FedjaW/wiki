# Constants Don't Matter

Big-O notation only describes the theoretical growth rate of algorithms. It doesn't deal with the actual time an algorithm takes to run on a given machine. As such, when doing Big O analysis, we don't let ourselves get bogged down in details.

For example, take a look at the following functions:

```py
def print_names_once(names):
    for name in names:
        print(name)
```

```py
def print_names_twice(names):
    for name in names:
        print(name)
    for name in names:
        print(name)
```

As you would expect, `print_names_once` will take half the time to run as `print_names_twice`. And in the real world of software engineering, cutting speed in half is awesome. The funny thing about Big O analysis is that we don't care. We're academics™.

Both of the functions above have the same rate of growth, `O(n)`. You might be tempted to say, "`print_names_twice` should be `O(2 * n)`" but you would be missing the whole point of Big O.

In Big O analysis we drop all constants because while they affect the runtime, they don't affect the change in the runtime.

- `O(2 * n)` -> `O(n)`
- `O(10 * n^2)` -> `O(n^2)`

