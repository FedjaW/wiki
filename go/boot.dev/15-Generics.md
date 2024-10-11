# Generics

As we've mentioned, Go does not support classes. For a long time, that meant that Go code couldn't easily be reused in many cases. For example, imagine some code that splits a slice into 2 equal parts. The code that splits the slice doesn't care about the _types_ of values stored in the slice. Before generics, we needed to write the same code for each type, which is a very un-DRY thing to do.

```go
func splitIntSlice(s []int) ([]int, []int) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
```

```go
func splitStringSlice(s []string) ([]string, []string) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
```

## TYPE PARAMETERS

Put simply, generics allow us to use variables to refer to specific types. This is an amazing feature because it allows us to write abstract functions that drastically reduce code duplication.

```go
func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
```

Example code to get the last element from a slice regardless of its type:

```go
func getLast[T any](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}

	lastElement := len(s) - 1
	return s[lastElement]
}
```

## CONSTRAINTS

Sometimes you need the logic in your generic function to know _something_ about the types it operates on. The example we used in the first exercise didn't need to know _anything_ about the types in the slice, so we used the built-in any constraint.

Constraints are just interfaces that allow us to write generics that only operate within the constraints of a given interface type. In the example above, the `any` constraint is the same as the empty interface because it means the type in question can be _anything_.

## INTERFACE TYPE LISTS

When generics were released, a new way of writing interfaces was also released at the same time!

We can now simply list a bunch of types to get a new interface/constraint.

```go
// Ordered is a type constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
        ~float32 | ~float64 |
        ~string
}
```

## PARAMETRIC CONSTRAINTS

Your interface definitions, which can later be used as constraints, can accept type parameters as well.

```go
// The store interface represents a store that sells products.
// It takes a type parameter P that represents the type of products the store sells.
type store[P product] interface {
	Sell(P)
}
```
