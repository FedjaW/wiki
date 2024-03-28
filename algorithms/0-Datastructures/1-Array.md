# Array

By the way...

```js
// javascript:
const a = [] // this is not an array!
```

- An Array is an contigious (unbreaking) memory space
- Arrays do not grow, neither shrink

What is the Big O of getting a specific value from the array?

- You access it by its starting memory address and taking the offset based on the size of the type
- a + width * offset
    - (a) beeing the address
    - (width) beeing the size of the type
    - (offset) beeing the input
- Big O is constant time because nothing depends on the input
- We just access the value directly