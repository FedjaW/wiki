```ts
// Status ist nun der Typalias für die Union 'open' | 'done' | 'discarded' 
type Status = 'open' | 'done' | 'discarded'  

// interface sind für objecte gedacht
interface Todo<TData> { // TData ist hier ein Generic
    description: String
    satus: Status // 'open' | 'done' | 'discarded' // union
    assignee?: String // ? -> optional
    data: TData
}

interface Metada {
    createdDateString: string
    modifiedDateString: string
}

const todo: Todo<Metada> = {
    description: 'Typescript lernen',
    satus: "open",
    data: { 
        createdDateString: 'gestern',
        modifiedDateString: 'heute'
    }
}

let text: string = "Hallo welt"
let text2 = "Hallo welt"

console.log(text);

const add = function (left: number, right: number): number {
    return left + right
}

// functionen in js haben eigentlich immer einen return wert.
// wenn man nix returned ist der return wert undefined
const printValue = function (value: unknown): void { // void, never, unknown, any
    console.log(value);
}

// - void
// mit void zeige ich an das diese function keinen rückgabewer hat, 
// aber in wirklichkeit wird trotzdem undefined zurück gegeben

// - never
// never ist eine besondere Variante von void
// never gibt an, dass die Funktion niemals zurückkehrt
// z.b. indem sie eine exception wirft

// - unknown
// unknown heißt ich bekomme irgendeinen Typ übergeben
// z.b. ich weiß im Vorgeld nicht was als parameter übergeben wird
// unknown sagt aus, dass der Typ mit typeof erst einmal geprüft werden sollte

// - any
// any ist eine alte Variante von unknown
// any lässt jede Operation zu (nach dem Motto, Entwickler weiß was er da tut...)

// ................

// C# und konsorten verwenden ein nominales typ-system
// ts verwendet ein strukturelles typ-system
```

