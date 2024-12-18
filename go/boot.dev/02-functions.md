# Functions

## NAMED RETURN VALUES

Return values may be given names, and if they are, then they are treated the same as if they were new variables defined at the top of the function.

Named return values are best thought of as a way to document the purpose of the returned values.

According to the tour of go:

> A return statement without arguments returns the named return values. This is known as a "naked" return. Naked return statements should be used only in short functions. They can harm readability in longer functions.

```go
func getCoords() (x, y int){
  // x and y are initialized with zero values

  return // automatically returns x and y
}
```

Is the same as:

```go
func getCoords() (int, int){
  var x int
  var y int
  return x, y
}
```
