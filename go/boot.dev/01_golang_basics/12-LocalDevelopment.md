# Local development

## PACKAGES

Every Go program is made up of packages.

You have probably noticed the package main at the top of all the programs you have been writing.

A package named "main" has an entrypoint at the main() function. A main package is compiled into an executable program.

A package by any other name is a "library package". Libraries have no entry point. Libraries simply export functionality that can be used by other packages. For example:

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

This program is an executable. It is a "main" package and imports from the fmt and math/rand library packages.

## MODULES

Go programs are organized into packages. A package is a directory of Go code that's all compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package (directory).

A repository contains one or more modules. A module is a collection of Go packages that are released together.

A GO REPOSITORY TYPICALLY CONTAINS ONLY ONE MODULE, LOCATED AT THE ROOT OF THE REPOSITORY.

A file named go.mod at the root of a project declares the module. It contains:

- The module path
- The version of the Go language your project requires
- Optionally, any external package dependencies your project has

The module path is just the import path prefix for all packages within the module. Here's an example of a go.mod file:

```go
module github.com/bootdotdev/exampleproject

go 1.22.1

require github.com/google/examplepackage v1.3.0
```

Each module's path not only serves as an import path prefix for the packages within but also indicates where the go command should look to download it. For example, to download the module golang.org/x/tools, the go command would consult the repository located at https://golang.org/x/tools.

> An "import path" is a string used to import a package. A package's import path is its module path joined with its subdirectory within the module. For example, the module github.com/google/go-cmp contains a package in the directory cmp/. That package's import path is github.com/google/go-cmp/cmp. Packages in the standard library do not have a module path prefix.

## SETTING UP YOUR MACHINE

Your machine will contain many version control repositories (managed by Git, for example).

Each repository contains one or more packages, but will typically be a single module.

Each package consists of one or more Go source files in a single directory.

The path to a package's directory determines its import path and where it can be downloaded from if you decide to host it on a remote version control system like GitHub or GitLab.

## A NOTE ON GOPATH

The $GOPATH environment variable will be set by default somewhere on your machine (typically in the home directory, `~/go`). Since we will be working in the new "Go modules" setup, you don't need to worry about that. If you read something online about setting up your GOPATH, that documentation is probably out of date.

These days you should avoid working in the `$GOPATH/src` directory. Again, that's the old way of doing things and can cause unexpected issues, so better to just avoid it.

## GET INTO YOUR WORKSPACE

Navigate to a location on your machine where you want to store some code. For example, I store all my code in `~/workspace`, then organize it into subfolders based on the remote location. For example,

`~/workspace/github.com/wagslane/go-password-validator` = https://github.com/wagslane/go-password-validator

That said, you can put your code wherever you want.

## Go run

Compiles and runs the code - more for debugging and testing purpose.

## Go build

Builds an executable.

## Go Install

### BUILD AN EXECUTABLE

Ensure you are in your hellogo repo, then run:

```go
go install
```

Go has compiled and installed the program globally.

If you get an error regarding "hellogo not found" it means you probably don't have your Go environment setup properly. Specifically, `go install` is adding your binary to your `GOBIN` directory, but that may not be in your `PATH`.

## Go get

`go get` is used to install remote packages (like `npm install`).
I see a difference to npm, `go get` will download from that given URI and does not need a registry like npm.

```go
go get github.com/wagslane/go-tinytime
```

## CLEAN PACKAGES

A main package isn't a library, there's no need to export functions from it.
