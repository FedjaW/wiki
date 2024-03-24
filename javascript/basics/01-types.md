```js
//-----------------------------------------------------------------------------------------
// *** Dynamisches Typsystem: ***
//-----------------------------------------------------------------------------------------
// Variablen nehmen implizit den Typ an des zugewiesenen Wertes an.
// Dieser kann sich zur Laufzeit ändern.

// Datentypen:

// 1. number (= double)
x = 42
pi = 3.1415
result = 23 / 0 // = infinity

// 2. string
text = "Hallo Welt!"
text = 'Hallo Welt!'
text = `Hallo ${pi} Welt!` // template string

// 3. boolean
isSunny = true
isRainy = false

// 4. undefined
x = undefined // x enthält den Wert undefined und hat den Typ undefined!

// 5. null
x = null // x enthält den Wert null und hat den Typ null!

console.log(typeof 3.14) // -> number
console.log(typeof undefined) // -> undefined

// Notiz: das der typeof Operator für null -> object zurück gibt ist ein Programmierfehler
// der seit über 26 Jahren nicht behoben wurde.
console.log(typeof null) // -> object (nicht null wie zu erwarten wäre)


//-----------------------------------------------------------------------------------------
// *** Werte & Referenztypen ***
//-----------------------------------------------------------------------------------------
// Wertetypen: -> number, boolean, undefined, null
// Wertetypen werden auf dem Stack verwaltet.
// Wertetypen haben immer die gleiche größe im Speicher.
// number ist so ein Wertetyp. 
// Egal ob ich 42 oder 3.4545461561654 oder 12651465465465 speichere, es nimmt immer den gleichen Speicher ein.

// Referenztypen: -> string
// Referentypen werden auf dem Heap verwaltet.
// string ist ein Referenztyp
// string hat eine sonderregel, er is immutable, also werden immer neue string erzeugt und keine verändert.
// string ist aus technischer sicht ein Referenztyp, verhält sich aber in vielen Dingen wie ein Wertetyp.

// js hat eine Garbage Collection (kümmert sich um die Bereinigung des Heap)

// Spannender Exkurs: Gleichheit string
"Hallo Welt" == "Hallo Welt" // true oder false?
// Antwort: 
//   Es kommt darauf an javascript aus diesen zwei string-literalen macht
//   Links von dem "==" (also das erste mal) wird "Hallo Welt" sicher auf dem Heap abgelegt
//   Rechts von dem "==" (also das zweite mal) wird dieses string-literal auch auf dem Heap abgelegt oder nicht?
//     String ist eine Sonderregel, es passiert ein sogenanntes interning, das heißt, dass erkannt wird, dass es diesen string schon gibt und er wird wiederverwendet
//   Also sind sie gleich. -> true

// Gleichheit string und number
0 == '0' // true
// 0 wird unter der haube stringyfied -> 0.toString()
x = 0
x.toString() == '0' // true

// Gleichheit number 0 und ''
0 == '' // true
// js kommt mit toString() nicht weiter also wird eine weiter function aufgerufen:
// valueOf()
y = ''
value = valueOf(y) // retuns 0
// 0 == 0

// Das alles ist etwas unintuitiv, daher gibt es den Operator ===
// === -> typsicherer Vergleich
0 === '' // false
// Das kann nicht gleich sein aufgrund der unterschiedliche typen
// === erspart sich valueOf() und toString()
// daher ist === etwas performanter als ==

//-----------------------------------------------------------------------------------------
// *** Variablen ***
//-----------------------------------------------------------------------------------------
x = 42 // globale variable
let x = 42 // lokale variable
const y = 42 // Konstante
```

