# Bubble Sort

See also: [Bubble Sort](../../primeagen/02-Sort/01-BubbleSort.md)

Bubble sort is a very basic sorting algorithm named for the way elements "bubble up" to the top of the list.

Bubble sort repeatedly steps through a slice and compares adjacent elements, swapping them if they are out of order. It continues to loop over the slice until the whole list is completely sorted. Here's the pseudocode:

1. Set `swapping` to `True`
2. Set `end` to the length of the input list
3. While `swapping` is `True`:
    1. Set `swapping` to `False`
    2. For `i` from the 2nd element to `end`:
        - If the `(i-1)`th element of the input list is greater than the `i`th element:
            1. Swap the `(i-1)`th element and the `i`th element
            2. Set `swapping` to `True`

    3. Decrement `end` by one

4. Return the sorted list

## Implementation

```py
def bubble_sort(nums):
    swapping = True
    end = len(nums)
    while swapping:
        swapping = False
        for i in range(1, end):
            if nums[i - 1] > nums[i]:
                temp = nums[i - 1]
                nums[i - 1] = nums[i]
                nums[i] = temp
                swapping = True
        end -= 1
    return nums
```

## Bubble Sort Big O

`O(n^2)`

Bubble sort's time complexity of `O(n^2)` comes from its nested loop structure. For each element in the list, it makes a pass through the list, potentially moving each element into its proper position. This results in roughly `n` operations for each of the `n` elements, leading to `n * n = n^2` iterations in the worst-case scenario.

### Best and Worst Case

Sometimes it's useful to know how the algorithm will perform based on what the input data is instead of just how much data there is. In the case of bubble sort (and many other algorithms), the best and worst case scenarios can actually change the time complexity.

- Best case: If the data is pre-sorted, bubble sort becomes really fast.
- Worst case: If the data is in reverse order, bubble sort becomes really slow (but still in the same complexity class as random data).

#### More Explanation

For the best case, if the list is already sorted, the algorithm only needs one pass through the list, resulting in a time complexity of `O(n)`. Think of it as the list taking a leisurely walk across the stepping stones with everything perfectly placed.

For the worst case, when the list is in reverse order, bubble sort must swap elements extensively in each pass, and it may go through nearly n passes, taking `O(n^2)` time. Imagine a throng of dancers having to rearrange until each step lines up perfectly!

That's why, amongst the choices, "At best `n`, at worst `n^2`" captures the essence of Bubble Sort.

### Why Bubble Sort?

Bubble sort is famous for how easy it is to write and understand.

However, it's one of the slowest sorting algorithms, and as a result is almost never used in practice. That said, we covered it because it's a useful thought exercise so that you can appreciate why the more complex and performant algorithms are better.
