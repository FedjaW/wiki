# Help and Exit

A REPL is only useful if it does something! Our REPL will work using the concept of "commands". A command is a single word that maps to an action.

We're going to support two commands in this step:

- `help`: prints a help message describing how to use the REPL
- `exit`: exits the program

## Assignment

1. Remove your logic that prints the first word (the command) back to the user
2. Add a callback for the exit command. Commands in our REPL are just callback functions with no arguments, but that return an error. For example:

```go
func commandExit() error
```

This function should print Closing the Pokedex... Goodbye! then immediately exit the program. I used os.Exit(0).

3. Create a "registry" of commands. This will give us a nice abstraction for managing the many commands we'll be adding. I created a struct type that describes a command:

```go
type cliCommand struct {
	name        string
	description string
	callback    func() error
}
```

Then I created a map of supported commands:



```go
map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
}
```

4. Register the exit command. Update your REPL loop to use the "command" the user typed in to look up the callback function in the registry. If the command is found, call the callback (and print any errors that are returned). If there isn't a handler, just print Unknown command.
5. Test your program (obviously).
6. Add a help command, its callback, and register it. It should print:

```txt
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
```

Hint: I dynamically generated the "usage" section by iterating over my registry of commands. That way the help command will always be up-to-date with the available commands.

Test your code again manually.

### Q & A

`Q:` what makes a function a callback

`A:`  A callback is simply a function that's passed as an argument to another function or stored in a data structure to be "called back" (executed) later.

What makes a function a callback is not its structure, but how it's used. In your Pokedex REPL, your callbacks are functions that:

1. Match a specific signature (in this case `func() error`)
2. Are stored in your command registry/map
3. Will be called later when specific events occur (when the user types a command)

Here's a simple example (unrelated to your assignment):

```go
type Button struct {
    onClick func() error
}

func sayHello() error {
    fmt.Println("Hello!")
    return nil
}

// sayHello becomes a callback when we assign it here
button := Button{
    onClick: sayHello,
}

// The function gets "called back" when we do this
button.onClick()
```

`Q:` why can I use func() error as a type, are funcs types in go?

`A:` In Go, functions are indeed first-class citizens, which means function types are just as valid as other types like `int` or `string`!

A function type in Go is defined by its signature - the combination of its parameter types and return types. `func() error` is a type that describes any function that:

1. Takes no parameters (empty parentheses)
2. Returns an error

Here are some other examples of function types (different from your assignment):

```go
// Function type that takes no params and returns nothing
type SimpleCallback func()

// Function type that takes an int and returns a string
type NumberFormatter func(int) string

// You can use them like any other type
var myCallback SimpleCallback
myCallback = func() {
    fmt.Println("I'm a callback!")
}
```

This is what allows you to store functions in structs or maps, just like you're doing in your REPL command registry. 