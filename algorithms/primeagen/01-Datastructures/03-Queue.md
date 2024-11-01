# Queue

A Queue is a specific implementation of a linked list.

First In First Out.

```js
type Node<T> = {
    value: T
    next?: Node<T>
}

export default class Queue<T> {
    public length: number
    private head?: Node<T>
    private tail?: Node<T>

    constructor() {
        this.head = this.tail = undefined
        this.length = 0
    }

    enqueue(item: T): void {
        const node = { value: item } as Node<T>
        this.length++
        if (!this.tail) {
            this.tail = this.head = node
            return
        }

        this.tail.next = node // let the current tail point to the new inserted node
        this.tail = node // let the new node be the tail now
    }

    deque(): T | undefined {
        if (!this.head) {
            return undefined
        }

        this.length--
        const head = this.head
        this.head = this.head.next

        return head.value
    }

    peek(): T | undefined {
        return this.head?.value
    }
}
```
