# Binary Search

- Precondition: data must be sorted

Lets say we jump by 10%
- Then you can jump by a fixed number and look if the value is greater than the input value
- Then you jump again
- If your value is less than the searched input value you go back one jump and linear search for the value
- The BigO is O(10 + 0.1*N) whic is O(N)
    - Here 10 because we assume we jump by 10% (it is just an example here, no proof)

Bettwe we jump by 1/2

- Then after some math you find out the BigO is O(logN)

```js
// my own solution
export default function bs_list(haystack: number[], needle: number): boolean {
    // Remember: the haystack must be (is) sorted!
    // 0 ........x......... N
    // lo + ((hi - lo) / 2)
    var l = 0;
    var h = haystack.length;

    do {
        var mid = Math.floor(l + (h - l) / 2);
        var value = haystack[mid];

        if (value === needle) {
            return true;
        }
        if (value > needle) {
            h = mid;
        } else {
            l = mid + 1;
        }
    } while (l < h);

    return false;
}

```