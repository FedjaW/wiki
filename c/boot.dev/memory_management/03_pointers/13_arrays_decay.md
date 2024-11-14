# Arrays Decay to Pointers

So we know that arrays are like pointers, but they're not exactly the same. Arrays allocate memory for all their elements, whereas pointers just hold the address of a memory location. In many contexts, arrays decay to pointers, meaning the array name becomes "just" a pointer to the first element of the array.

## When arrays decay

Arrays decay when used in expressions containing pointers:

```c
int arr[5];
int *ptr = arr;          // 'arr' decays to 'int*'
int value = *(arr + 2);  // 'arr' decays to 'int*'
```

And also when they're passed to functions... so they actually decay quite often in practice. That's why you can't pass an array to a function by value like you do with a struct; instead, the array name decays to a pointer.

## When arrays don't decay

- `sizeof` Operator: Returns the size of the entire array (e.g., `sizeof(arr)`), not just the size of a pointer.
- `&` Operator Taking the address of an array with `&arr` gives you a pointer to the whole array, not just the first element. The type of `&arr` is a pointer to the array type, e.g., `int (*)[5]` for an `int` array with `5` elements.
- Initialization: When an array is declared and initialized, it is fully allocated in memory and does not decay to a pointer.
