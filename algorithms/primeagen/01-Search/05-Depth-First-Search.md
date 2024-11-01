# Depth-First-Search

- The tree must be somehow in the correct "shape".
  - meaning the left hand side must be lower or equal than the "pivot aka parent node"
  - and the right hand side must be greater

Related to this is [[02-BinarySearch]]

- Binary search is harder to implement but it has always a O(log(n))
- where DFS has O(log(n)) - O(log(n))
  - it depends on the shape of the tree
  - log(n) if you have a single line of nodes
  - log(n) if you have a balanced tree, see: [[06-Tree]]

```js
type BinaryNode<T> = {
    value: T
    left: BinaryNode<T>
    right: BinaryNode<T>
}

function search(curr: BinaryNode<number> | null, needle: number): boolean {
    if(!curr) {
        // curr is null
        return false
    }

    if (curr.value === needle) {
        return true
    }

    // traverse
    if (curr.value < needle) {
        return search(curr.right, needle)
    }
    return search(curr.left, needle)
}

export default function dfs(
    head: BinaryNode<number>,
    needle: number,
): boolean {
    return search(head, needle)
}
```
