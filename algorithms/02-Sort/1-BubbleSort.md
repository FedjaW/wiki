# Bubble Sort

Sorted if `Xi <= Xi+1`. X is the value of the array, i the index.

Very simple sorting algorithm. Compare Xi and Xi+1 and swap if Xi is bigger than Xi+1.

The first iteration will produce an array where the largest nummer is at the end.

Then you go over it again and again.

It bubbles the biggest numbers to the right of the array.

```js
export default function bubble_sort(arr: number[]): void {
    for (var i = 0; i < arr.length - 1; ++i) {
        for (var j = 0; j < arr.length - 1 - i; ++j) {
            if (arr[j] > arr[j + 1]) {
                // swap
                var temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
            }
        }
    }
}
```


Big O discussion:

First run you need N steps.
Second run you need N-1 steps. 
Third run you need N-2 steps. 
..
Last run you need N-N+1 steps.

Now remember `Little Gauß` rule: 1 + 2 + 3 ... + N = (N+1)*N/2
So this aligns with our algorithm. 

O((N+1)*N/2) = O((N+1)*N) = O(N^2+N) = O(N^2)

We will drop non significant values, so we drop the +N because N^2 will grow so much faster.