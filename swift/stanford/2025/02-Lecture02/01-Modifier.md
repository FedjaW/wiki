# View Modifier

Read the comments in the code to learn about view modifier.

```swift
import SwiftUI

struct ContentView: View {
    var body: some View {
        VStack {
            Image(systemName: "globe")

            // Let's us a font modifier.
            // .font() is a function defined on the View protocol!
            // .font() returns a new View!
            // So we can chain onto that View, e.g.: with foregroundStyle()
            Text("Welcoem to CS193p!")
                .font(.largeTitle)
                .foregroundStyle(.green)

            Text("Greetings!")
            Circle()
        }
    }
}

struct ContentView2: View {
    var body: some View {
        VStack {
            Image(systemName: "globe")
            Text("Welcoem to CS193p!")
                .font(.largeTitle)
                .foregroundStyle(.green)
            Text("Greetings!")

            // Look at that...
            // Here .font() won't have any effect
            // but it exists even on Circle(), because
            // Circle() behaves like a View and the View protocol
            // has a font modifier
            Circle().font(.largeTitle)
        }
    }
}

struct ContentView3: View {
    var body: some View {
        VStack {
            Image(systemName: "globe")
            Text("Welcoem to CS193p!")
                .font(.largeTitle)
                .foregroundStyle(.green)
            Text("Greetings!")

            // Look at that...
            // Here .font() won't have any effect
            // but it exists even on Circle(), because
            // Circle() behaves like a View and the View protocol
            // has a font modifier
            Circle().font(.largeTitle)
        }
    }
}

// Let's try what will happen if we put a modifier onto VStack
// VStack behaves like a View, so the same modifiers can be used
// as for the other Views
// What did happen?
// The VStack itself can't do much with a .font(.largeTitle)
// so it passes it on to its content (to every View inside)
// until it can do something with it, itself, see ContentView5
struct ContentView4: View {
    var body: some View {
        VStack {
            Image(systemName: "globe")
            Text("Welcoem to CS193p!")
                .foregroundStyle(.green)
            Text("Greetings!")
            Circle().font(.largeTitle)
        }
        .font(.largeTitle)
    }
}

// Here .padding() will modify Text.
// The second .padding() on VStack will put a padding around
// the whole content of VStack and not pass it to every View inside
struct ContentView5: View {
    var body: some View {
        VStack {
            Image(systemName: "globe")
            Text("Welcoem to CS193p!")
                .padding()
                .foregroundStyle(.green)
            Text("Greetings!")
            Circle().font(.largeTitle)
        }
        .font(.largeTitle)
        .padding()
    }
}
```
