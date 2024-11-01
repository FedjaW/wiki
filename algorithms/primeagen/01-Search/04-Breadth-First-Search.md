# Breadth First Search

- Related to this is: [[06-Tree]]. There we use recursion to traverse a tree.

- Visit one level of the tree at a time
- A `queue` is the thing we need here!
- The runtime of a BFS is O(N)
- We do not need recursion here

```js
type BinaryNode<T> = {
     value: T
     left: BinaryNode<T>
     right: BinaryNode<T>
}

export default function bfs(
    head: BinaryNode<number>,
    needle: number,
): boolean {
    const q: (BinaryNode<number> | null)[] = [head] // the queue

    while (q.length) {
        const curr = q.shift() as BinaryNode<number>

        // search
        if (curr?.value === needle) {
            return true
        }
        if (curr.left) {
            q.push(curr.left)
        }
        if (curr.right) {
            q.push(curr.right)
        }

    }

    return false
}
```
