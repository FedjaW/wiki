# Model and UI & Swift Type System

## Model and UI

Seperating "Logic and Data" from "UI"

It cannot be overemphasized how important it is to be clear about "source of thruth".
Duplicating "truth" in your Views can lead to lots of logic erros.

Luckily, SwiftUI makes this very simple.
Firstly, it uses struct. structs are generally immutable and are copied.
Thus you have to go out of your way to create a mutable copy of a soruce of thruth.

SwiftUI also makes you mark all sources of thruth with either @State or @Observable.
We'll see how we use @State soon.

Any "truth sharing" is also marked (@Binding, @Observable)

## Type System

### struct

What is in a struct?

```swift
// Stored vars
var masterCode: Code

// Computed vars
var body: some View {
    return Text("Hello World")
}

// Constant lets (values never change)
let widht = 3

// Functions
func multiply(operand: Int, by: Int) -> Int {
    // Weird to write because wanna say
    // operand times other operand
    return operand * by
}
// Weird naming because of operand
// Wanna say multiply 5 by 6
multiply(operand: 5, by: 6)

// Solution
func multiply(_ operand: Int, by otherOperand: Int) -> Int {
    return operand * otherOperand
}
multiply(5, by: 6)

// Initializers
struct RoundedReactangle {
    init(cornerRadius: CGFloat) {
        // init
    }
    init(cornerSize: CGSize) {
        // init
    }
}
```

### struct vs class

There is a huge difference between a struct and a class!

A struct (and enum) is value type and a class is reference type.

A value type stores its value (itself) directly in a var/let.

A reference type stores a pointer (a reference) to its value in a var/let.

Therefore...

The mutability of a struct is explicit (var vs. let).

Not so with a class (they are always mutable, no matter wether you put it in a var or let).

Views are structs and so are completly read-only (they are essentially in a let somewhere).

Thus, you have to use a special notation (@State) to make a stored var in a View be modifiable.

You also have to use a special notation (@Binding) to share a struct between Views.

A class does have one big advantage: sharing. So we will only use a class in SwiftUI when we want extrem sharing.

classes also have a strong sense of identity (the pointer to them).

structs need to explicitly declare their identity (usually throug a var named id).

### init

init is a special statuc func which is used to create an instance of a struct.

e.g. RoundedReactangle(cornerRadius: 10) calls this func...

```swift
struct RoundedReactangle: View, Shape {
    init(cornerRadius: CGFloat) {
        // init all vars in a RoundedReactangle that results in the given cornerRadius
    }
}
```

Notes...

1. no need to put statuc before init (all inits are, by definition, static)
2. no need to return the created gthing inside, just initielize all the struct's vars
3. multiple inits with different arguments are ok
4. if you don't give a struct an init, you'll get a free one that includes all vars

### enum

Min: 41

https://www.youtube.com/watch?v=B42CuI0RO7Y&list=PLoROMvodv4rPHblRXKsJCQs8TLGpiCTrG&index=4
