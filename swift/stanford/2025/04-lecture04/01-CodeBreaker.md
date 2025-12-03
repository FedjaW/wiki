# Code Breaker

```swift
//
//  CodeBreaker.swift
//  CodeBreaker
//
//  Created by Friedrich Wilms on 30.11.25.
//

import SwiftUI

typealias Peg = Color

// struct is immutable by default,
// therefore we have to tell swift if it has to copy on write the struct
// mutating func is doing that.
struct CodeBreaker {
    var masterCode: Code = Code(kind: .master)
    var guess: Code = Code(kind: .guess)
    var attempts: [Code] = [] // [Code]()
    let pegChoices: [Peg]

    init(pegChoices: [Peg] = [.red, .green, .blue, .yellow]) {
        self.pegChoices = pegChoices
        masterCode.randomize(from: pegChoices)
        print(masterCode)
    }

    mutating func changeGuessPeg(at index: Int) {
        let existingPeg = guess.pegs[index]
        if let indexOfExistingPegInPegChoices = pegChoices.firstIndex(of: existingPeg) {
            let newPeg = pegChoices[(indexOfExistingPegInPegChoices + 1) % pegChoices.count]
            guess.pegs[index] = newPeg
        } else {
            guess.pegs[index] = pegChoices.first ?? Code.missing
        }
    }

    mutating func attemptGuess() {
        var attempt = guess // copies the guess
        attempt.kind = .attempt(guess.match(against: masterCode)) // modify the copy to be of kind .attempt
        attempts.append(attempt)
    }
}

struct Code {
    var kind: Kind
    var pegs: [Peg] = Array(repeating: Code.missing, count: 4)

    // static -> static is a global variable with namespacing to Code,
    // therefore we need to refere to it by explicitly typing Code.mssing
    static let missing: Peg = .clear

    // Behave like an equatable
    // meaning it will have the == (static func ==)
    // otherwise we can't compare because of the assosiated data
    // Equatable will synthesise if for me: static func ==(lhs: Kind, rhs: Kind) -> Bool
    enum Kind: Equatable {
        case master
        case guess
        case attempt([Match]) // "assosiated data" only for the case that needs information about the matches
        case unknown
    }

    mutating func randomize(from pegChoices: [Peg]) {
        for index in pegChoices.indices {
            pegs[index] = pegChoices.randomElement() ?? Code.missing
        }
    }

    var matches: [Match] {
        switch kind {
        case .attempt(let matches): return matches // grap that "assoisated data" (matches) and return it
        default: return [] // any other kind one return empty array
        }
    }

    func match(against otherCode: Code) -> [Match] {
        var results: [Match] = Array(repeating: .nomatch, count: pegs.count)
        var pegsToMatch = otherCode.pegs
        for index in pegs.indices.reversed() {
            if pegsToMatch.count > index, pegsToMatch[index] == pegs[index] {
                results[index] = .exact
                pegsToMatch.remove(at: index)
            }
        }
        for index in pegs.indices {
            if results[index] != .exact {
                if let matchIndex = pegsToMatch.firstIndex(of: pegs[index]) {
                    results[index] = .inexact
                    pegsToMatch.remove(at: matchIndex)
                }
            }
        }
        return results
    }
}
```
