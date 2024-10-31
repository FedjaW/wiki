# First Class and Higher Order Functions

A programming language "supports first-class functions" when functions are treated like any other variable. That means functions can be passed as arguments to other functions, can be returned by other functions, and can be assigned to variables.

First-class function: A function that is treated like any other value
Higher-order function: A function that accepts another function as an argument or returns a function
Python supports first-class and higher-order functions.

## First-class example

```py
def square(x):
    return x * x

# Assign function to a variable
f = square

print(f(5))
# 25
```

## Higher-order example

```py
def square(x):
    return x * x

def my_map(func, arg_list):
    result = []
    for i in arg_list:
        result.append(func(i))
    return result

squares = my_map(square, [1, 2, 3, 4, 5])
print(squares)
# [1, 4, 9, 16, 25]
```
