# Pointers

Pointers hold the memory address of a value.

The \* syntax defines a pointer:

```go
var p *int
```

A pointer's zero value is nil

The & operator generates a pointer to its operand.

```go
myString := "hello"
myStringPtr := &myString
```

The \* dereferences a pointer to gain access to the value

```go
fmt.Println(*myStringPtr) // read myString through the pointer
*myStringPtr = "world"    // set myString through the pointer
```

## NIL POINTERS

Pointers can be very dangerous.

If a pointer points to nothing (the zero value of the pointer type) then dereferencing it will cause a runtime error (a panic) that crashes the program. Generally speaking, whenever you're dealing with pointers you should check if it's nil before trying to dereference it.

## POINTER RECEIVERS

A receiver type on a method can be a pointer.

Methods with pointer receivers can modify the value to which the receiver points. Since methods often need to modify their receiver, pointer receivers are more common than value receivers. However, methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.

### POINTER RECEIVER

```go
type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "blue"
}
```

### NON-POINTER RECEIVER

```go
type car struct {
	color string
}

func (c car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "white"
}
```

### POINTER RECEIVER CODE

Methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.

```go
type circle struct {
	x int
	y int
    radius int
}

func (c *circle) grow(){
    c.radius *= 2
}

func main() {
    c := circle{
        x: 1,
        y: 2,
        radius: 4,
    }

    // notice c is not a pointer in the calling function
    // but the method still gains access to a pointer to c
    c.grow()
    fmt.Println(c.radius)
    // prints 8
}
```
