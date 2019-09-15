# Intro to Go (part 1)

## Overview

TODO: Fill in language overview

## Notes

### [01: Hello, World!](code/01/main.go)

- `package main`
  - The main package for the application; each executable Go application must have a main package
  - We'll talk more about packages in a bit

- `func main() {}`
  - The main function for the main package; each main package must have a main function.
  - This is where the application starts

- `import "fmt"`
  - Imports the [`fmt`](https://golang.org/pkg/fmt/) package
  - This package is in the standard library, and "...implements formatted I/O..."

- `fmt.Println("Hello, World!")`
  - Calls the [`Println`](https://golang.org/pkg/fmt/#Println) function in the `fmt` package, providing the string `"Hello, World!"`
  - Prints "Hello, World!" plus a new line in the terminal
