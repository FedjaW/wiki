# No-Op

A no-op is an operation that does... nothing.

If a function doesn't return anything, it's probably impure. If it doesn't return anything, the only reason for it to exist is to perform a side effect.

## Example no-op

This function performs a useless computation because it doesn't return anything or perform a side-effect. It's a no-op.

```py
def square(x):
    x * x
```

## Example side-effect

This function performs a side effect. It changes the value of the y variable that is outside of its scope. It's impure.

```py
y = 5
def add_to_y(x):
    global y
    y += x

add_to_y(3)
# y = 8
```

The *global* keyword just tells Python to allow access to the outer-scoped y variable.

