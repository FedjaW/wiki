# Tree

```js
type Node<T> = {
    value: T
    children: Node<T>[]
}

// for a binary tree
type BinaryNode<T> = {
    value: T
    left: BinaryNode<T>
    right: BinaryNode<T>
}
```

## Terminology

- root - the most parent node. The First. Adam.
- height - The longest path from the root to the most child node
- binary tree - a tree in which has at most 2 children, at least 0 children
- general tree - a tree with 0 or more children
  binary search tree - a tree in which has a specific ordering to the nodes and at most 2 children
- leaves - a node without children
- balanced - A tree is perfectly balanced when any node's left and right children have the same height.
- branching factor - the amount of children a tree has.

## Traverse a tree - Depth first traversal

DFS: We walk the left side of the tree until the last node is visisted of the left hand side.
Then we walk on level up and so on..

Depth First Search preserves shape of the tree, Breadth first search does not!
Hence, you can use DFS to check if two trees are equal!

```js
type BinaryNode<T> = {
    value: T
    left: BinaryNode<T>
    right: BinaryNode<T>
}

function walk(curr: BinaryNode<number> | null, path: number[]): void {
    if (!curr) {
        return
    }

    // do the 3 stepf of a recursion

    //  1. pre
    path.push(curr.value)

    //  2. recurse
    walk(curr.left, path)
    walk(curr.right, path)

    //  3. post
}

export default function pre_order_search(head: BinaryNode<number>): number[] {
    const path: number[] =  []
    walk(head, [])

    return path
}
```
