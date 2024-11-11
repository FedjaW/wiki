# Insertion Sort

Insertion sort builds a sorted list one item at a time. It's much less efficient on large lists than merge sort because it's `O(n^2)`, but it's actually faster (not in Big O terms, but due to smaller constants) than merge sort on small lists.

## Pseudocode

1. For each index in the input list:
    1. Set a `j` variable to the current index
    2. While `j` is greater than `0` and the element at index `j-1` is greater than the element at index `j`:
        1. Swap the elements at indices `j` and `j-1`
        2. Decrement `j` by `1`
2. Return the list

### Implementation

```py
def insertion_sort(nums):
    for i in range(len(nums)):
        j = i
        while j > 0 and nums[j - 1] > nums[j]:
            nums[j], nums[j - 1] = nums[j - 1], nums[j]
            j -= 1
    return nums

```

### Big O

Insertion sort has a Big O of `O(n^2)`, because that is its worst case complexity.

The outer loop of insertion sort always executes `n` times, while the inner loop depends on the input.

- Best case: If the data is pre-sorted, insertion sort becomes really fast.
- Average case: The average case is `O(n^2)` because the inner loop will execute about half of the time.
- Worst case: If the data is in reverse order, it's still `O(n^2)` and the inner loop will execute every time.

## Why Use Insertion Sort?

- Fast: for very small data sets (even faster than merge sort and quick sort, which we'll cover later)
- Adaptive: Faster for partially sorted data sets
- Stable: Does not change the relative order of elements with equal keys
- In-Place: Only requires a constant amount of memory
- Online: Can sort a list as it receives it

## Why is insertion sort fast for small lists?

Many production sorting implementations use insertion sort for very small inputs under a certain threshold (very small, like `10`-ish), and switch to something like quicksort for larger inputs. They use insertion sort because:

- There is no recursion overhead
- It has a tiny memory footprint
- It's a stable sort as described above
