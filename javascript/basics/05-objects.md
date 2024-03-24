```js
// Ein Object in js ist etwas viel einfacheres als Objects in anderen Sprachen
// Ein Object in js ist eine Liste von key/value Paaren (entspricht in anderen Sprachen eher einem Dictionary)

const company = {
    name: 'papas-rezepte',
    location: 'Bielefeld',
    foundedIn: 2023,
    fullName: function () {
        return `${this.name}.de`
    } 
}

// es geht auch: (method shorthand syntax)
const company2 = {
    name: 'papas-rezepte',
    location: 'Bielefeld',
    foundedIn: 2023,
    fullName () {
        return `${this.name}.de`
    } 
}

console.log(company.name)               // papas-rezepte
console.log(company.company['name'])    // papas-rezepte

console.log(company.company['country'])    // undefined

company.country = 'Germany' // es wird ein weiteres Property angefügt
console.log(company.company['country'])    // Germany

const person = {} // leeres dict aka "object"
const person2 = null // null zeigt an, dass hier ein object erwatet wird aber keins zugewiesen ist
// deswegen auch typeof null -> object
```
