# Variables

Go's basic variable types are:

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

## Zero values

Variables declared without an explicit initial value are given their zero value.

The zero value is:

- 0 for numeric types,
- false for the boolean type, and
- "" (the empty string) for strings.

## SHORT VARIABLE DECLARATION

Inside a function (like the main function) the `:=` short assignment statement can be used in place of a var declaration. The `:=` operator infers the type of the new variable based on the value. It's colloquially called the walrus operator because it looks like a walrus... sort of.

```go
// These two lines of code are equivalent:
var empty string
empty := ""

// ---
numCars := 10 // inferred as an integer
temperature := 0.0 // temperature is inferred as a float because it has a decimal
var isFunny = true // inferred as a boolean
```

## SAME LINE DECLARATIONS

You can declare multiple variables on the same line:

```go
mileage, company := 80276, "Tesla"
```

## TYPE SIZES

Ints, uints, floats, and complex numbers all have type sizes.

```go
int  int8  int16  int32  int64 // whole numbers

uint uint8 uint16 uint32 uint64 uintptr // positive whole numbers

float32 float64 // decimal numbers

complex64 complex128 // imaginary numbers (rare)
```

The size (8, 16, 32, 64, 128, etc) represents how many bits in memory will be used to store the variable. The "default" int and uint types refer to their respective 32 or 64-bit sizes depending on the environment of the user.

The standard sizes that should be used unless the developer has a specific need are:

- int
- uint
- float64
- complex128

## PREFER "DEFAULT" TYPES

A problem arises when we have a uint16, and the function we are trying to pass it into takes an int. We're forced to write code riddled with type casts like int(myUint16).

This style of development can be slow and annoying to read. When Go developers stray from the “default” type for any given type family, the code can get messy quickly.

Unless you have a good reason to, stick to the following types:

- bool
- string
- int
- uint
- byte
- rune
- float64
- complex128

## CONSTANTS

Constants are declared with the const keyword. They can't use the := short declaration syntax.

```go
const pi = 3.14159
```

Constants can be character, string, boolean, or numeric values. They can not be more complex types like slices, maps and structs, which are types we will explain later.

As the name implies, the value of a constant can't be changed after it has been declared.

## COMPUTED CONSTANTS

Constants must be known at compile time. They are usually declared with a static value:

```go
const myInt = 15
```

However, constants can be computed as long as the computation can happen at compile time.

For example, this is valid:

```go
const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName
```

That said, you cannot declare a constant that can only be computed at run-time like you can in JavaScript. This breaks:

```js
// the current time can only be known when the program is running
const currentTime = time.Now()
```

## FORMATTING STRINGS IN GO

The %v variant prints the Go syntax representation of a value, it's a nice default.

```go
s := fmt.Sprintf("I am %v years old", 10)
// I am 10 years old

s := fmt.Sprintf("I am %v years old", "way too many")
// I am way too many years old
```

If you want to print in a more specific way, you can use the following formatting verbs:

STRING

```go
s := fmt.Sprintf("I am %s years old", "way too many")
// I am way too many years old
```

INTEGER

```go
s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old
```

FLOAT

```go
s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old
```

If you're interested in all the formatting options, you can look at the fmt package.

## THE INITIAL STATEMENT OF AN IF BLOCK

An if conditional can have an "initial" statement. The variable(s) created in the initial statement are only defined within the scope of the if body.

```go
if INITIAL_STATEMENT; CONDITION {
}
```

This is just some syntactic sugar that Go offers to shorten up code in some cases. For example, instead of writing:

```go
length := getLength(email)
if length < 1 {
    fmt.Println("Email is invalid")
}
```

We can do:

```go
if length := getLength(email); length < 1 {
    fmt.Println("Email is invalid")
}
```

Not only is this code a bit shorter, but it also removes length from the parent scope, which is convenient because we don't need it there - we only need access to it while checking a condition.