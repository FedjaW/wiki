```js
// javascript hat keine Klassen!

// Wenn man this nutzt sollte es eigentlich Dog geschrieben sein 
// ist aber nur Konvention
const dog = function (name) {
    this.name = name
}

console.log(typeof dog) // => function

// ich rufe hier eine function mit new auf
// das new bewirkt dass unter der haube ein neues object generiert wird
// dieses neue object wird an die dog function als implizieter parameter übergeben
// auf dieses object kann ich mit "this" zugreifen
const myDog = new dog('Alice')

// FALLSTRICK - Aufrug ohne new
const myOtherDog = dog('Werner')
// hier wird die dog() function ohne das new aufgerufen
// in der dog function wrid aber this benutzt
// this will auf das, unter der haube erzeugte, object zugreifen
// das object wird unter der haube aber nur erzeugt wenn 
// die function mit dem new keyword aufgerufen wurde

// was ist this denn nun?

// in dem Falle ist this das globale Object in js (window object)
// man würde also die name property von dem globalen window object überschreiben -> TRORLOLOLOLOL

// im strict mode wäre this dann glücklichwerweise undefined -> würde dann eine exception per Laufzeit geben

//-----------------------------------------------------------------------------------------------
// ECMAScript 2015
//-----------------------------------------------------------------------------------------------
// Ist nur syntactic sugar
class Cat {
    constructor (name) {
        this.name = name
    }

    miau () {
        console.log("MIAU");
    }
}

console.log(typeof Cat) // => function (und nicht class) TORORLOLOL
```
