# Code Breaker View

```swift

//
//  CodeBreakerView.swift
//  CodeBreaker
//
//  Created by Friedrich Wilms on 27.11.25.
//

import SwiftUI

struct CodeBreakerView: View {
    // Start always with a let, then if you have to modify it, change it to a var
    @State var game = CodeBreaker(pegChoices: [.brown, .yellow, .orange, .black])

    var body: some View {
        VStack {
            view(for: game.masterCode)
            ScrollView {
                view(for: game.guess)
                ForEach(game.attempts.indices.reversed(), id: \.self) { index in
                    view(for: game.attempts[index])
                }
            }
        }
        .padding()
    }

    var guessButton: some View {
        Button("Guess") {
            withAnimation {
                game.attemptGuess()
            }
        }
        .font(.system(size: 80)) // system font, system size if very big
        .minimumScaleFactor(0.1) // 10%, i want this font to be 80 but minimumScalFactor will say I take a small as 10% of that, which is 8
        // Notice that the view modifiers are used on the Button.
        // The Button itself can't do anything with it, so it passes them to the inner content of the Button
        // And there will be probably a Text that can use the modifier.
    }

    func view(for code: Code) -> some View {
        HStack {
            ForEach(code.pegs.indices, id: \.self) { index in
                RoundedRectangle(cornerRadius: 10)
                    .overlay {
                        if code.pegs[index] == Code.missing {
                            RoundedRectangle(cornerRadius: 10)
                                .strokeBorder(Color.gray)
                        }
                    }
                    .contentShape(Rectangle()) // here to get the multytouch events, otherwise we would need a little bit of opacity, because currently we use clear.
                    .aspectRatio(1, contentMode: .fit)
                    .foregroundStyle(code.pegs[index])
                    .onTapGesture {
                        if code.kind == .guess {
                            game.changeGuessPeg(at: index)
                        }
                    }
            }
            MatchMarkers(matches: code.matches)
                .overlay {
                    // overlay will fit into the space that you are overlaying it,
                    // other than a ztack.
                    if code.kind == .guess {
                        guessButton
                    }
                }
        }
    }
}

#Preview {
    CodeBreakerView()
}

```
