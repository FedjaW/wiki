```js
setTimeout(() => {
    console.log("Hallo Welt")
}, 1_000) 

// 1_000 - numeric seperator - besser leserlich

const sayHello = async function () {
    await sleep(1_000)
    console.log("Hallo Welt")
}

const main = async function () {
    await sayHello()
}

// Top-level await, darf nur einmal aufgeruden werden
// await main()

// oder per hand händeln
main().catch(e => console.error(e))
```