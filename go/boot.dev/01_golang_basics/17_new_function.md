# new()

It allocates memory for a variable of the given type and returns a pointer to it. Magical, isn't it?

For example, you can use new with built-in types or custom types:

```go
i := new(int)       // Creates a pointer to an integer, initialized to 0
s := new(string)    // Creates a pointer to a string, initialized to ""
b := new(bool)      // Creates a pointer to a boolean, initialized to false
```

For custom types:

```go
type Wizard struct {
    name    string
    power   int
}

w := new(Wizard) // Creates a pointer to a Wizard, fields initialized to default zero values
w.name = "Merlin"
w.power = 100
```

`Q:` Now then, can you think of a reason why we might use `new` to allocate memory when pointers (`&` operator) could accomplish similar tasks? Why not simply write `p := &Wizard{}?` What's the advantage of `new`?

`A:` Both new and the & operator can create pointers, but the choice often comes down to style, clarity, and the specific task at hand.

Let's compare the two approaches:

1. Using new:

```go
i := new(int)
```

This creates a pointer to an integer with a default value of 0. It's simple and concise when you just need a zero-valued object.

2. Using the & operator:

```go
i := &int{}
```

This achieves the same result: it also creates a pointer to an integer with a default value of 0. However, this syntax includes explicit braces, which may look more complex.


Now, why choose one over the other?

- Use `new` when you want to keep the code simple and focus on explicitly creating a zero-value object.
- Use `&{}` __or a struct literal__ when you want to create a specific initial state, especially for composite types like structs. For example:

```go
w := &Wizard{name: "Merlin", power: 100}
```

This cannot be accomplished using `new`.
