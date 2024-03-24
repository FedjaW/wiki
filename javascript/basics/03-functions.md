```js
// ------------------------------------------------------------------
// Function statement
// ------------------------------------------------------------------
const sum = add(23, 42)

console.log(sum);

function add (left, right) {
    return left + right
}

// - man kann die Funktion aufrufen, bevor sie deklariert wurde
// js scanned die Datei vorher


// ------------------------------------------------------------------
// Function expression
// ------------------------------------------------------------------
const add2 = function (left, right) {
    return left + right
}

const sum2 = add2(23, 42)

console.log(sum2);

// - man kann die Funktion NICHT aufrufen, bevor sie deklariert wurde
// weil eine Function expression eine ZUWEISUNG ist und wie jede andere Zuweisung auch
// wird sie erst per Laufzeit ausgeführt
```