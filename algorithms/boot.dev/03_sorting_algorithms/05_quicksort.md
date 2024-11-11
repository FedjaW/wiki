# Quick Sort

See [Quick Sort](../../primeagen/02-Sort/02-Quicksort.md)

Quick sort is an efficient sorting algorithm that's widely used in production sorting implementations. Like merge sort, quick sort is a recursive divide and conquer algorithm.

## Divide

- Select a pivot element that will preferably end up close to the center of the sorted pack
- Move everything onto the "greater than" or "less than" side of the pivot
- The pivot is now in its final position
- Recursively repeat the operation on both sides of the pivot

## Conquer

- The array is sorted after all elements have been through the pivot operation

## Pseudocode

- Select a "pivot" element - We'll arbitrarily choose the last element in the list
- Move through all the elements in the list and swap them around until all the numbers less than the pivot are on the left, and the numbers greater than the pivot are on the right
- Move the pivot between the two sections where it belongs
- recursively repeat for both sections

## Implementation

```py
def quick_sort(nums, low, high):
    if low < high:
        p = partition(nums, low, high)
        quick_sort(nums, low, p - 1)
        quick_sort(nums, p + 1, high)


def partition(nums, low, high):
    pivot = nums[high]
    i = low
    for j in range(low, high):
        if nums[j] < pivot:
            nums[i], nums[j] = nums[j], nums[i]
            i += 1
    nums[i], nums[high] = nums[high], nums[i]
    return i
```

## Big O

On average, quicksort has a Big O of `O(n*log(n))`. In the worst case, and assuming we don't take any steps to protect ourselves, it can degrade to `O(n^2)`. `partition()` has a single for-loop that ranges from the lowest index to the highest index in the array. By itself, the `partition()` function is `O(n)`. The overall complexity of quicksort is dependent on how many times `partition()` is called.

Worst case: The input is already sorted. An already sorted array results in the pivot being the largest or smallest element in the partition each time, meaning `partition()` is called a total of `n` times.

Best case: The pivot is the middle element of each sublist which results in `log(n)` calls to `partition()`.

## Fixing Quick Sort

While the version of quicksort that we implemented is almost always able to perform at speeds of `O(n*log(n))`, its Big O is still technically `O(n^2)` due to the worst-case scenario. We can fix this by altering the algorithm slightly.

Two of the approaches are:

- Shuffle input randomly before sorting. This can trivially be done in `O(n)` time.
- Actively find the median of a sample of data from the partition, this can be done in `O(1)` time.

### Random Approach

The random approach is easier to code, which is nice if you're the one writing the code.

The function simply shuffles the list into random order before sorting it, which is an `O(n)` operation. The likelihood of shuffling a large list into sorted order is so low that it's not worth considering.

### Median Approach

Another popular solution is to use the "median of three" approach. Three elements (for example: the first, middle, and last elements) of each partition are chosen and the median is found between them. That item is then used as the pivot.

This approach has less overhead, and also doesn't require randomness to be injected into the function, meaning it can remain deterministic and pure.

## Why Use Quick Sort

Pros:

- Very fast: At least it is in the average case
- In-Place: Saves on memory, doesn't need to do a lot of copying and allocating

Cons:

- Typically unstable: changes the relative order of elements with equal keys
- Recursive: can incur a performance penalty in some implementations
- Pivot sensitivity: if the pivot is poorly chosen, it can lead to poor performance

All this said, quicksort is widely used in the real world because the trade-offs are often worth it. For example, it's a default in PostgreSQL, a popular open-source database.

I'd also like to shoutout timsort, which is a hybrid sorting algorithm that uses a combination of merge sort and insertion sort. It's the default sorting algorithm used by the sorted function in Python.