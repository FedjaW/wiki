# Binary Trees

Trees aren't particularly useful data structures unless they're ordered in some way. One of the most common types of ordered tree is a Binary Search Tree or `BST`. In addition to the properties we've already talked about, a `BST` has some additional constraints:

1. Instead of an unbounded list of children, each node has at most 2 children
2. The left child's value must be less than its parent's value
3. The right child's value must be greater than its parent's value
4. No two nodes in the `BST` can have the same value

By ordering the tree this way, we'll be able to add, remove, find, and update nodes much more quickly.
