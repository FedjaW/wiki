# Linear search

- Going through every element of the array
- O(N)
- Even if the first index matches, we consider the worst case
- It might not be in the array (this is the worst case)

```js
export default function linear_search(
    haystack: number[],
    needle: number,
): boolean {
    for (let i = 0; i < haystack.length; i++) {
        if (haystack[i] === needle) {
            return true;
        }
    }
    return false;
}

```
