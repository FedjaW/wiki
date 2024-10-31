# Functions as Values

In Python, functions are just values, like strings, integers, or objects. For example, we can assign an existing function to a variable:

```py
def add(x, y):
    return x + y

# assign the function to a new variable
# called "addition". It behaves the same
# as the original "add" function
addition = add
print(addition(2, 5))
# 7
```