# Reduction to P

Let's take an exponential time algorithm and fix it up so it can run in polynomial time!

The Fibonacci sequence is a sequence of numbers where each number is the sum of the two numbers before it. Like this:

> 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...

We want a function that, given an index in the sequence, returns the Fibonacci number at that index. For example:

- `fib(0)` -> 0
- `fib(1)` -> 1
- `fib(2)` -> 1
- `fib(3)` -> 2
- `fib(4)` -> 3
- `fib(5)` -> 5
- `fib(6)` -> 8
- `fib(7)` -> 13

Here are the implementation details to do it in polynomial time:

1. The input `n` represents the index of the desired Fibonacci value
2. Initialize three values. One will hold the `current` value, one will hold the parent value (1 before), and one will hold the `grandparent` value (2 before the current). You'll need to hard-code the first two values of the sequence. `grandparent=0` and `parent=1`
3. Write a loop that iterates `n-1` times. (For example, the loop should not repeat when `n=2`.)
4. Inside the loop update the `current` based on the parents, then update the parents to their appropriate values.
5. Return the current value once the loop completes

## Implementation 

```py
def fib(n):
    if n <= 1:
        return n
    current = 0
    parent = 1
    grandparent = 0
    for i in range(0, n - 1):
        current = parent + grandparent
        grandparent = parent
        parent = current
    return current

```
