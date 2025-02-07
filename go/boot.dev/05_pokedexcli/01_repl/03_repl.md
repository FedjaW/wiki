# REPL

REPL stands for "read-eval-print loop". It's a common type of program that allows for interactivity. You type in a command, and the program evaluates it and prints the result. You can then type in another command, and so on.

Your native terminal's command line is a REPL! Our Pokedex will also be a REPL. In this step, we're just going to build the bones of the REPL: a loop that parses and cleans an input and prints the first word back to the user. Here is what your program should be able to do after this step:

```cli
> go run .
Pokedex > help me
help
Pokedex > Exit
exit
Pokedex > WHAT even is a pokeman?
what
```

`Pokedex >` is our own custom command line prompt. The program waited and recorded my input (in the first instance, `help me`) and didn't continue until I pressed `enter`.

## Assignment

1. Remove your "Hello, World!" logic.
2. Create support for a simple REPL

- Wait for user input using bufio.NewScanner (this blocks the code and waits for input, once the user types something and presses enter, the code continues and the input is available in the returned scanner)
- Start an infinite `for` loop. This loop will execute once for every command the user types in (we don't want to exit the program after just one command)
    - Use fmt.Print to print the prompt `Pokedex >` without a newline character
    - Use the scanner's .Scan and .Text methods to get the user's input as a string
    - Clean the users input by trimming any leading or trailing whitespace, and converting it to lowercase. I used strings.ToLower and strings.Fields to split the input into a slice of words
  - Capture the first "word" of the input and use it to print: `Your command was: <first word>`

3. Test your program. Here's my example session:

```cli
wagslane@MacBook-Pro-2 pokedexcli % go run .
Pokedex > well hello there
Your command was: well
Pokedex > Hello there
Your command was: hello
Pokedex > POKEMON was underrated
Your command was: pokemon
```

You can terminate the program by pressing `ctrl+c`.

4. Run the CLI again and `tee` the output (copies the stdout) to a new file called `repl.log` (and `.gitignore` the log).

```cli
go run . | tee repl.log
```

- Use this as the first input: `CHARMANDER is better than bulbasaur`.
- Use this as the second input: `Pikachu is kinda mean to ash`.
- Terminate the program by pressing `ctrl+c`.
