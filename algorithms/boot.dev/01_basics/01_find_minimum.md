# Find Minimum

An "algorithm" is just a set of instructions that can be carried out to solve a problem.

People use algorithms all the time without even realizing it. Practically every function you write in code is an algorithm (well, kinda), even if it's a simple one.

## Assignment

Implement the "find minimum" algorithm in Python by completing the `find_minimum()` function:

1. Set minimum to positive infinity: `float("inf")`.
2. For each number in the list `nums`, compare it to `minimum`. If `num` is smaller, set `minimum` to `num`.
3. `minimum` is now set to the smallest number in the list.

```py
def find_minimum(nums):
    if nums == []:
        return None
    min = float("inf")
    for n in nums:
        if n < min:
            min = n
    return min

```
