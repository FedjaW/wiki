# Pure Functions

If you take nothing else away from this course, please take this: Pure functions are fantastic. They have two properties:

They always return the same value given the same arguments.
Running them causes no side effects
In short: pure functions don't do anything with anything that exists outside of their scope.

## Example of a pure function

```py
def find_max(nums):
    max_val = float("-inf")
    for num in nums:
        if max_val < num:
            max_val = num
    return max_val
```

## Example of an impure function

```py
# instead of returning a value
# this function modifies a global variable
global_max = float("-inf")

def find_max(nums):
    global global_max
    for num in nums:
        if global_max < num:
            global_max = num
```

## Pure Function Review

Pure functions have a lot of benefits. Whenever possible, good developers try to use pure functions instead of impure functions. Remember, pure functions:

Return the same result if given the same arguments. They are deterministic.
Do not change the external state of the program. For example, they do not change any variables outside of their scope.
Do not perform any I/O operations (like reading from disk, accessing the internet, or writing from the console).
These properties result in pure functions being easier to test, debug, and think about.

One of the biggest differences between good and great developers is how often they incorporate pure functions into their code. Pure functions are easier to read, easier to reason about, easier to test, and easier to combine. Even if you're working in an imperative language like Python, you can (and should) write pure functions whenever reasonable.

There's nothing worse than trying to debug a program where the order functions are called needs to be juuuuust right because they all read and modify the same global variable.
