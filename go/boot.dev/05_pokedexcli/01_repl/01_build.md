# Build

By the end of this step, you'll have a simple working program in Go!

## Assignment

1. Create a main.go file. It should be part of package main in the root of your project and have a main() function that just prints the text "Hello, World!".
2. Create a Go module in the root of your project. Here's the command:

```cli
go mod init MODULE_NAME
```

I recommend naming the module by its remote Git location (you should store all your projects in Git!). For example, my GitHub name is wagslane so my module name might be github.com/wagslane/pokedexcli.

3. Build your program:

```cli
go build
```

Note: The executable will be named after the directory containing the main.go file, hence mine was pokedexcli. You should .gitignore the executable.

4. Run your program:

```cli
./pokedexcli
```

It should print "Hello, World!" to the console!
