# Code Breaker Game V1

Learn to use ForEach loop with correct id and trailing closure syntax.

```swift
import SwiftUI

struct ContentView: View {
    var body: some View {
        VStack {
            pegs(colors: [.red, .green, .green, .yellow])
            pegs(colors: [.blue, .red, .green])
        }
        .padding()
    }

    func pegs(colors:  Array<Color>) -> some View {
        HStack {
            // ForEach over the index, colors.indices is used because it is unique.
            // The id: \.self references the struct itself, and its unique because of indices.
            // every struct has a .self
            ForEach(colors.indices, id: \.self) { index in
                Circle().foregroundStyle(colors[index])
            }

            // ForEach is the way to go, no this version of the code
            // to be failsafe no to access a index that does not exist.
            /*
            Circle().foregroundStyle(colors[0])
            Circle().foregroundStyle(colors[1])
            Circle().foregroundStyle(colors[2])
            Circle().foregroundStyle(colors[3])
            */
        }
    }
}
```

Here we can see the next step to really make a game with some squares and circles.
Also we see that the pegs func is to larg, we want to do a decomposition, see comment and the next version of it.

```swift
import SwiftUI

struct ContentView: View {
    var body: some View {
        VStack {
            pegs(colors: [.red, .green, .green, .yellow])
            pegs(colors: [.blue, .red, .green, .yellow])
            pegs(colors: [.blue, .red, .green, .yellow])
        }
        .padding()
    }

    // This is now a very long function
    // Let's do some "DECOMPOSITION"
    func pegs(colors:  Array<Color>) -> some View {
        HStack {
            ForEach(colors.indices, id: \.self) { index in
                RoundedRectangle(cornerRadius: 10)
                    .aspectRatio(1, contentMode: .fit)
                    .foregroundStyle(colors[index])
            }
            HStack {
                VStack {
                    Circle().fill(.green)
                    Circle().strokeBorder(.primary, lineWidth: 3).aspectRatio(1, contentMode: .fit)
                }
                VStack {
                    Circle()
                    Circle().opacity(0)
                }
            }
        }
    }
}
```

Here we extracted code into a new View. We even could move MatchMarkers into a new file.

```swift
struct ContentView: View {
    var body: some View {
        VStack {
            pegs(colors: [.red, .green, .green, .yellow])
            pegs(colors: [.blue, .red, .green, .yellow])
            pegs(colors: [.blue, .red, .green, .yellow])
        }
        .padding()
    }

    func pegs(colors:  Array<Color>) -> some View {
        HStack {
            ForEach(colors.indices, id: \.self) { index in
                RoundedRectangle(cornerRadius: 10)
                    .aspectRatio(1, contentMode: .fit)
                    .foregroundStyle(colors[index])
            }
            MatchMarkers()
        }
    }
}

struct MatchMarkers: View {
    var body: some View {
        HStack {
            VStack {
                Circle().fill(.green)
                Circle().strokeBorder(.primary, lineWidth: 3).aspectRatio(1, contentMode: .fit)
            }
            VStack {
                Circle()
                Circle().opacity(0)
            }
        }
    }
}
```

Extracted into a new file and implemented the logic to display the correct marker syling.

ContentView.swift

```swift
import SwiftUI

struct ContentView: View {
    var body: some View {
        VStack {
            pegs(colors: [.red, .green, .green, .yellow])
            pegs(colors: [.blue, .red, .green, .yellow])
            pegs(colors: [.blue, .red, .green, .yellow])
        }
        .padding()
    }

    func pegs(colors:  Array<Color>) -> some View {
        HStack {
            ForEach(colors.indices, id: \.self) { index in
                RoundedRectangle(cornerRadius: 10)
                    .aspectRatio(1, contentMode: .fit)
                    .foregroundStyle(colors[index])
            }
            MatchMarkers(matches: [Match.exact]) // currently we pass only a single Match, don't know if this makes any sense.
        }
    }
}
```

MatchMarkers.swift

```swift
import SwiftUI

// An enum can have functions but no "stored" variables.
enum Match {
    case nomatch
    case exact
    case inexact
}

struct MatchMarkers: View {
    var matches: [Match] // same as Array<Match>

    var body: some View {
        HStack {
            VStack {
                matchMarker(peg: 0)
                matchMarker(peg: 1)
            }
            VStack {
                matchMarker(peg: 2)
                matchMarker(peg: 3)
            }
        }
    }

    func matchMarker(peg: Int) -> some View {
        // Functional programming right here
        let exactCount: Int = matches.count(where: { match in match == .exact })
        let foundCount: Int = matches.count(where: { match in match != .nomatch })

        // Instead of return you can use @ViewBuilder on func matchMarker.
        // Why does it work to put @ViewBuilder?
        // @ViewBuilder is a list of views and can have lets and if in there.
        return Circle()
            .fill(exactCount > peg ? Color.primary : Color.clear)
            .strokeBorder(foundCount > peg ? Color.primary : Color.clear, lineWidth: 2).aspectRatio(1, contentMode: .fit)
    }

    // This is valid code even withour @ViewBuilder nor return statement.
    // This is because in swift a "oneliner" is returned, it is syntactic sugar.
    func matchMarkerExample(peg: Int) -> some View {
        Circle()
    }
}
```

Explanation of inline func

```swift
import SwiftUI

enum Match {
    case nomatch
    case exact
    case inexact
}

struct MatchMarkers: View {
    var matches: [Match]

    var body: some View {
        HStack {
            VStack {
                matchMarker(peg: 0)
                matchMarker(peg: 1)
            }
            VStack {
                matchMarker(peg: 2)
                matchMarker(peg: 3)
            }
        }
    }

    func matchMarker(peg: Int) -> some View {
        // where: takes a func that takes a Match and returns a Bool
        // That is what isExact is, a func that takes a Match and returns a Bool
        // let exactCount = matches.count(where: isExact)

        // The final verion, see below for explanation
        let exactCount = matches.count { $0 == .exact }
        let foundCount = matches.count { $0 != .nomatch }

        return Circle()
            .fill(exactCount > peg ? Color.primary : Color.clear)
            .strokeBorder(foundCount > peg ? Color.primary : Color.clear, lineWidth: 2).aspectRatio(1, contentMode: .fit)
    }

    func isExact(match: Match) -> Bool {
        return match == Match.exact
    }
    func isExact2(match: Match) -> Bool {
        // A single line Statement, can remove the return
        match == Match.exact
    }
    func isExact3(match: Match) -> Bool {
        // Can remove the Match because .exact can be inferred because we compare with match
        match == .exact
    }
    // Now we make this func inline and put it into where param of count func
    // the "match in" represents the type of the isExact func
    // let exactCount = matches.count(where: { (match: Match) -> Bool in match == .exact })

    // now we can reduce more because of type inference
    // let exactCount = matches.count(where: { match in match == .exact })

    // A step further
    // $0 is the first argument of the input parameter
    // let exactCount = matches.count(where: { $0 == .exact })

    // Trailing closure syntax
    // let exactCount = matches.count { $0 == .exact }
}
```
