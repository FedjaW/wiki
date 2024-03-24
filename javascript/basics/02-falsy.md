```js
// Falsy values
// Falsy values sind in js das equivialent zu dem boolschen literal wie false 

let x = false // boolscher wert

if (x) {
    console.log("I will never be run ...");
}

// Falsy values
//   - false
//   - 0
//   - ''
//   - undefined 
//   - null

let y = null

// y ist ein falsy value, wird also in ein false umgewandelt
if (y) {
    console.log("I will never be run ...");
}

// logische operation

let a = 23
let b = 42
let result = a || b // 23
// js versucht a in einen boolschen wert umzuwandeln
// da 23 kein falsy wert ist, ist es im umkehrschluss ein truthy value
// da das a schon als true gewertet ist und das für oder (||) ausreichend ist wird gar nicht weiter evaluiert und result auf den zu letz evaluierten Wert gesetzt

let n = 0
let k = 42
let resultOr = n || k // 42
let resultAnd = n && k // 42, selbst wenn k falsy wäre, es wird immer der wert zurück gegeben der als letztes EVALUIERT wurde, also hier in jedem fall der wert der in k steht!
```
