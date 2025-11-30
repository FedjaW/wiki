# Intro

Read the comments and how the code progresses to learn about Views and syntax.

```swift
import SwiftUI

// ContentView "behaves" like a View
// A View is a protocal; Swift is a protocal oriented language
// What is a protocal? A protocal is a description of a "behaviour"
struct ContentView: View {
    // Int is a struct
    var i: Int = 0 // means it is an Int, an Int is an concrete type

    // To behave like a View you need to implement the "body" variable!
    // some View... means this has to be concrete type but it has to be some View
    // The actual concrete type will be detemined by the code in the brackets after the some View
    // The content after the brackets is a computed property
    // body is a property of a struct and is computed every single time the variable is requested.
    var body: some View {
        VStack {
            Image(systemName: "globe") // Image is a struct that behaves like a View
            Text("Hello, world!") // Text is a struct that behaves like a View
            Circle()
        }
    }
}

// What is this weird syntax of VStack {..}
// Let's rewrite it, (see above and then here)
struct ContentViewPart2: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 1, content: greetings)
    }

    func greetings() -> Text {
        return Text("greetings!")
    }
}

// Let's go on rewriting it
struct ContentViewPart3: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 1, content: greetings)
    }

    // I want it to return a "bag of lego" View not only Text.
    // @ViewBuilder will interpretate the stuff inside of it as a list of views
    // Notice the return type, a TupleView
    // We never type TupleView, we type "some View"
    // "some View", the compiler will figure out the return type.
    @ViewBuilder
    func greetings() -> TupleView<(Image, Text, Circle)> { // some View
        Image(systemName: "globe")
        Text("greetings!")
        Circle()
    }
}

// Let's go on rewriting it
// We never write functino with @ViewBuilder like this, we inline it
struct ContentViewPart4: View {
    var body: some View {
        // NOTICE: What about the @ViewBuilder?!
        // The @ViewBuilder is part of the content in the VStack definition
        // Go and ctrl+RIGHTCLICK on VStack and inspect it, there you will find the @ViewBuilder

        /*
        init(
            alignment: HorizontalAlignment = .center,
            spacing: CGFloat? = nil,
            @ViewBuilder content: () -> Content
        )
        */

        VStack(alignment: .leading, spacing: 1, content: {
            Image(systemName: "globe")
            Text("greetings!")
            Circle()
        })
    }
}

// How to we get to what we had in the beginning, what was VStack {..}
struct ContentViewPart5: View {
    // If you have a function or a creation of a struct and the last paramater of it is a function,
    // then you can remove the label (here content:) dn put the function on the outside.
    var body: some View {
        VStack(alignment: .leading, spacing: 1) {
            Image(systemName: "globe")
            Text("greetings!")
            Circle()
        }
    }
}

struct ContentViewPart6: View {
    var body: some View {
        // If VStack does use the default values AND there is a closure afterwards,
        // then we can remove the paranthesis
        // Here we are
        VStack {
            Image(systemName: "globe")
            Text("greetings!")
            Circle()
        }
    }
}
```
