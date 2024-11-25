# Linked List Queue Recap

Let's recap a few key points about linked lists and queues:

1. The problem with our array (normal Python list) implementation of a queue was that it had O(n) push operations.
2. The linked list implementation of a queue has O(1) push operations.
3. Linked lists are usually a worse choice than standard arrays because:
    - They are less performant due to more memory overhead
    - They are more complex to work with, debug, and reason about
    - They have no random access (indexing to a specific element)
4. Linked lists are a better choice than arrays in very specific circumstances because they have O(1) insertions and deletions at both ends of the list.
