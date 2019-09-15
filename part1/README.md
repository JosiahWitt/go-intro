# Intro to Go (part 1)

## Overview

TODO: Fill in language overview

## Notes

### [01: Hello, World!](code/01/main.go) - [Playground](https://play.golang.org/p/HkcX8SIizVQ)

- `package main`
  - The main package for the application; each executable Go application must have a main package
  - We'll talk more about packages in a bit

- `func main() {...}`
  - The main function for the main package; each main package must have a main function.
  - This is where the application starts

- `import "fmt"`
  - Imports the [`fmt`](https://golang.org/pkg/fmt/) package
  - This package is in the standard library, and "...implements formatted I/O..."

- `fmt.Println("Write Intro to Go (part 1)")`
  - Calls the [`Println`](https://golang.org/pkg/fmt/#Println) function in the `fmt` package, providing the string `"Write Intro to Go (part 1)"`
  - Prints "Write Intro to Go (part 1)" plus a new line in the terminal


### [02: Variables](code/02/main.go) - [Playground](https://play.golang.org/p/yDhsbXG84c2)

- `var item1, item2 string`
  - Declares the `item1` and `item2` variables to be of type `string`
  - `item1` and `item2` can only hold strings
  - Their zero value is the empty string (`""`)
  - Unused variables in Go is a compile time error

- `item2 = "Write Intro to Go (part 1)"`
  - Assigns the string `"Write Intro to Go (part 1)"` to the variable `item2`

- `var item3 string = "Listen to the Go Time podcast"`
  - Declares the `item3` variable to be of type `string` and assigns the string to `"Listen to the Go Time podcast"`

- `item4 := "Read the Golang Weekly email"`
  - Uses the `:=` assignment operator for type inference
  - Declares the `item4` variable to be of type `string` and assigns the string to `"Read the Golang Weekly email"`
  - The preferred way to set variables

- Go types include: `bool`, `string`, `int`, `int64`, `byte`, `rune` (unicode), `float64`, and more


### [03: Loops](code/03/main.go) - [Playground](https://play.golang.org/p/oNdEYw7_DAW)

- `for i := 0; i < 5; i++ {...}`
  - `for` is the only type of loop
  - Notice that there are no parenthesis
  - `i := 0` assigns the variable `i` to be an `int` with value `0`
  - `i < 5` says "while `i` is less than `5`"
  - `i++` increments `i` by one

- `for i < 5 {...}`
  - `for` doubles as a while loop
  - Any `bool` value can be used in place of `i < 5`


### [04: If Statements](code/04/main.go) - [Playground](https://play.golang.org/p/OdLylnQjvgC)

- `if item1 == "" {...}`
  - Checks if `item1` is the empty string
  - Notice that there are no parenthesis
  - Any `bool` value can be used in place of `item1 == ""`

- `else if item1 == "todo" || item1 == "wip" {...}`
  - Optional block that is evaluated if _all_ previous conditions are false and `item1 == "todo"` or `item1 == "wip"` evaluates to `true`

- `else {...}`
  - Optional block that is evaluated if _all_ previous conditions are false

- `if i < 5 && i != 4 {...}`
  - Block is evaluated if `i < 5` and `i` does not equal `4`
  - Yes, I know, this should be written `i < 4`, but I wanted to show _not_ and _and_
