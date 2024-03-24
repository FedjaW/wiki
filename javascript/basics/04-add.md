```js
const add = function (left, right) {
    return
        left + right // unreachable code
}

const sum = add(23, 42) // will return undefined

console.log(sum);


// more examples ----------------------------

const add2 = function (left, right) {
    return
        (left + right) // unreachable code
}


const add3 = function (left, right) {
    return (
        left + right) // this is ok
}


// more examples ----------------------------

const sumx = add(23) // second parameter is undefined
const sumxxxx = add(23, 42, 7, 5) // does not break anything


// more examples ----------------------------

// alles ab dem dritten parameter wird gesammelt in ein array gepackt
// dieser muss nicht rest heißen (frei benennbar)
const addrest = function (left, right, ...rest) {
    console.log(rest)
    return left + right
}

// default wert festlegen
const adddefault = function (left, right = 0) {
    return left + right
}
```
