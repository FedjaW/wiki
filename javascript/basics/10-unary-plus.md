```js
/**
 * What will be the output of the following statements?
 * Unary plus tries to convert the operand to a number
 */

let log = console.log

log(+"1") // 1
log(+true) // 1
log(+false) // 0
log(+"") // 0 (empty string is faly value)
log(+"hello") // NaN
log(+1) // 1
```