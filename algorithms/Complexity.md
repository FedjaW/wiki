# Complexity O(n)

Big O is a notation to categorize your algorithms time and memory requirements based on input. It is meant to generalize the growth of your algorithm not to be meant an exact measurement.

## Important Concepts

1. Growth is with respect to the input
2. Constants are dropped
    - O(2N) -> O(N) and this makes sense. That is because Big O is meant to describe the upper bound of the algorithm. The constant eventually becomes irrelevant
3. Concider worst case

## Constant time: 1

No matter how many elements we are working with, the algorithm/operation/whatever will always take the same time.

## Logarithmic time: log(n)

You have this if doubling the number of elements you are iterating over does not double the amount of work. Always assume that searching operations are log(n)

## Linear time: n

Iterating through all elements in a colleciton of data. If you see a for loop spanning from 0 to array.lenght, you probably have n or linear runtime.

## Quasilinear time: n*log(n)

You have this if doubling the number of elements you are iterating over does not double the amount of work. Always assume that any sorting operation is n*log(n)

## Quadratic time: n^2

Every element in a collection has to be compare to every other element. THe handshake problem.

## Exponential time: 2^n

If you add a single element to a collection, the processing power required doubles.

## Identifying Runtime Complexity

- Iterating with a simple for loop through a single collection -> "Probably" O(n) -> Linear

- Iterating through half of a collection -> Still O(n). There are no constants in runtime.

- Iterating through two different collections with separate for loops -> O(n + m)

- Two nested for loops iterating over the same collection -> O(n^2)

- Two nested for loops iterating over different collections -> O(n*m)

- Sorting -> O(n*log(n))

- Searching a sorted array -> O(log(n))
