# Tests

By now, you should be familiar enough with unit testing. Like anything, the more you do it, the better you will be at it. More importantly, you will begin to understand how to write code that is easier to test.

Currently our Pokedex is more like a read-print-loop instead of a read-EVAL-print-loop. We need to do something with the input. We'll use TDD for this part.

## Assignment

1. Create a new cleanInput(text string) []string function. For now it should just return an empty slice of strings.

The purpose of this function will be to split the users input into "words" based on whitespace. It should also lowercase the input and trim any leading or trailing whitespace. For example:

- `hello world` -> `["hello", "world"]`
- `Charmander Bulbasaur PIKACHU` -> `["charmander", "bulbasaur", "pikachu"]`

2. Create a new file for some unit tests. I called mine `repl_test.go`. (The only requirement is that it ends in `_test.go`.) Create a test suite for the `cleanInput` function. Here is the basic structure of the test file:

__All tests go inside TestXXX functions that take a *testing.T argument:__

```go
func TestCleanInput(t *testing.T) {
    // ...
}
```

__I like to start by creating a slice of test case structs, in this case:__

```go
cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	// add more cases here
}
```

__Then I loop over the cases and run the tests:__

```go
for _, c := range cases {
	actual := cleanInput(c.input)
	// Check the length of the actual slice
	// if they don't match, use t.Errorf to print an error message
	// and fail the test
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
	}
}
```

3. Once you have at least a few tests, run the tests using go test ./... from the root of the repo. We expect them to fail.
4. Implement the cleanInput function to make the tests pass.