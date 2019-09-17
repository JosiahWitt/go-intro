# Intro to Go (part 1)

## Overview

### What is Go?

- Go is a **compiled**, **fast**, and **safe** open source language developed by Google
  - **Compiled**
    - Everything required to run the binary is included in the binary (versus shared libraries in C)
    - Simple to deploy; run `go build` to build the binary, then upload and start the standalone binary on the server
    - Cross compilation is supported via `GOARCH` and `GOOS` environment variables
    - Great for projects that will run on another computer (ie. most things), including API Servers, CLIs, etc.

  - **Fast**
    - Since it is compiled and typed, it runs directly on the CPU
    - Compile times are fast
    - Development is fast, due to types, type inference, simplicity of the language, auto completion, auto import, documentation, etc.

  - **Safe**
    - Types catch typos at compile time, instead of during execution
    - Garbage collected; you don’t manage memory allocation by hand
    - Yes, there are pointers, but no, you can’t do anything more dangerous than having a `nil` pointer, which throws an error (this also happens in most languages, such as Ruby, JavaScript, etc.)
    - Go 1 compatibility promise: no breaking changes while on Go 1, and hopefully a Go 2 will never come, since they hope to always maintain backwards compatibility
    - Patch releases around monthly, minor releases twice a year
    - Language changes are well thought out
    - Large standard library (including HTTP clients and routers, JSON, etc.), which is included in the Go 1 compatibility promise
    - Package management is built into the language, including cryptographic verification


### Brief History

- Go was created by Robert Griesemer, Rob Pike (Plan 9, UTF-8), and Ken Thompson (Unix, C, Plan 9, UTF-8) at Google, and publicly announced in 2009
- Since 2009, it has gained traction in a number of companies and open source projects, such as Docker and Kubernetes

### Installing

- Want to play with Go without installing it yet?
  - Use the [Go Playground](https://play.golang.org)
  - You'll be able to follow along until we get to `12: Packages`

- MacOS
  - `brew install go`
  - `vi ~/.bashrc` or `vi ~/.zshrc`
    - Add:
      ```
      export GOPATH=$HOME/code/go # This is where my Go code lives
      export GOROOT=/usr/local/opt/go/libexec
      export PATH=$PATH:$GOPATH/bin
      export PATH=$PATH:$GOROOT/bin
      ```
  - `source ~/.bashrc` or `source ~/.zshrc`
  - `mkdir -p $GOPATH "$GOPATH/src" "$GOPATH/pkg" "$GOPATH/bin"`

- Your Go code lives in the `$GOPATH/src` folder under your repository
  - eg. `$GOPATH/src/github.com/<my_username>/<my_repo>`
  - These directories don't need to be `git init`ialized for one-off projects
  - With the new Go Modules (hopefully finalized in Go 1.14, though currently in Go 1.11 - 1.13, just not the default), you'll no longer be tied to your `$GOPATH`, but for now, I find it to be best way for working on Go projects

### Tools

- `go`
  - Most tooling you need is in the `go` CLI
  - `go run file_name.go`
    - Runs the `file_name.go` file
  - `go build`
    - Builds the current project into a binary with the directory name
  - `go get github.com/some/repo`
    - Downloads the provided repo to the `$GOPATH`, so it can be imported

### Learn More

- [Go Website](https://golang.org)
- [Go Tour](https://tour.golang.org)
- [Effective Go](https://golang.org/doc/effective_go.html) - doc about writing code the "Go Way"

---

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

- Notice that there are no semicolons at the end of expressions

---

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

---

### [03: Loops](code/03/main.go) - [Playground](https://play.golang.org/p/oNdEYw7_DAW)

- `for i := 0; i < 5; i++ {...}`
  - `for` is the only type of loop
  - Notice that there are no parenthesis
  - `i := 0` assigns the variable `i` to be an `int` with value `0`
  - `i < 5` says "while `i` is less than `5`"
  - `i++` increments `i` by one

- `for i < 5 {...}`
  - `for` doubles as a while loop, since the "init" and "post" sections are optional
  - Any `bool` value can be used in place of `i < 5`

- If you need an infinite loop (eg. processing TCP connections), you can simply write `for {...}`

- Variables declared inside a loop don't escape the block

---

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

- Variables declared inside an if statement don't escape the block

- If you want to declare a value just for use within the if statement as well as for use in the if block, you can use the following:
```
if val := 2 + 3; val <= 5 {
  fmt.Println(val)
}
```

---

### [05: Slices and Arrays](code/05/main.go) - [Playground](https://play.golang.org/p/UXgN_whHC9i)

- Our project is getting quite messy, let's clean up!

- `todos := []string{...}`
  - Creates and fills the `string` slice
  - If nothing is provided (eg. `[]string{}`), an empty slice is created

- `todos = append(todos, ...)`
  - Appends more items to a copy of the `todos` slice, and assigns that new slice to the `todos` variable

- `todos[4] = "Buy a Go Gopher plush"`
  - Changes the value at index `4` to "Buy a Go Gopher plush"

- `todo := todos[0]`
  - Gets the value at index `0`
  - The program panics at runtime if the index is out of bounds

- `fmt.Printf("%d TODOs (showing first 5):\n", len(todos))`
  - `len(todos)`
    - Returns the length of the slice `todos`
  - Read about `fmt.Printf` below

- `firstFive := todos[:5]`
  - "Slices" the slice to the first `5` elements of the slice
  - Can specify starting and ending point
    - eg. `mySlice[start:end]`, where `start` and `end` are both positive integers
    - Start is inclusive, end is exclusive
    - Skipping start defaults to 0
      - eg. `mySlice[:10] == mySlice [0:10]`
    - Skipping end defaults to `len(mySlice)`; the rest of the slice
      - eg. `mySlice[3:] == mySlice [3:len(mySlice)]`

- `for i, todo := range todos {...}`
  - Ranges over the `todos` slice
  - Returns the index and value for each item, stored in `i` and `todo`, respectively
  - If you don't want the index, use the `_` (blank identifier) to specify that that value won't be used
    - eg. `for _, todo := range todos {...}`
  - If you don't want the value, skip it
    - eg. `for i := range todos {...}`

- `fmt.Printf("%d: %s\n", i+1, todo)`
  - [`fmt.Printf`](https://golang.org/pkg/fmt/#Printf) uses format printing, for a list of verbs, see the [docs](https://golang.org/pkg/fmt/#hdr-Printing)
  - `%d` is for a decimal (base 10) integer
  - `%s` is for a string
  - `%v` (not used here) uses the default for the type

- Go also has arrays, but they have fixed length, so usually slices are used, since they can always be appended to
  - To create an array, specify the length in the type
    - eg. `todos := [3]string{}`
    - Creates a string array of length 3

- Can make slices or arrays of any type

---

### [06: Maps](code/06/main.go) - [Playground](https://play.golang.org/p/PFdjJ0yjDHY)

- We're making progress; now we can display a list of TODO items to a user, but we have no way of knowing if an item has been completed or not

- `todos := map[string]bool{...}`
  - Creates and fills a map with `string` keys and `bool` values
  - If nothing is provided (eg. `map[string]bool{}`), an empty map is created

- `todos["Read the Go 1.13 release notes"] = false`
  - Sets the key `"Read the Go 1.13 release notes"` equal to `false`

- `done, ok := todos["Write Intro to Go (part 1)"]`
  - Returns the value for the key "Write Intro to Go (part 1)" and if it was ever set
  - `ok` is a `bool`; `true` => the value was set, `false` => the value was not set
  - Can skip the "comma ok"
    - Returns the zero value if not set
    - eg. `done := todos["Write Intro to Go (part 1)"]`
    - Zero values in Go:
      - `0` for numeric types (`int`, `float64`, etc.)
      - `false` for `bool`s,
      - `""` (the empty string) for `string`s
      - `nil` for pointers

- `fmt.Printf("%d TODOs:\n", len(todos))`
  - `len(todos)`
    - Returns the length of the map

- `for todo, done := range todos {...}`
  - Ranges over the `todos` map
  - Returns the key and value for each item, stored in `todo` and `done`, respectively
  - Follow the same rules as slices to ignore the key or value

- Can make maps of any type

- Maps are randomized when using `range`
  - This is to prevent anyone relying on the ordering of maps, since the ordering relied on the underlying implementation

---

### [07: Structs](code/07/main.go) - [Playground](https://play.golang.org/p/W9k6fdl_Q-Z)

- Nice! We're starting to make a TODO list. However, it doesn't seem right to use a map to store TODO values and completion status. What if we wanted to add a due date. Would we add another level of maps? No, in Go, we would create a struct.

- A struct is a lightweight way to structure data

- `type Todo struct {...}`
  - Declares a new type called `Todo`, which is a `struct`
  - Hint: You can also create other types, like a custom integer type
    - eg. `type MyInts int`

- `Title string`
  - The struct field `Title` is of type `string`

- `Done bool`
  - The struct field `Done` is of type `bool`

- `todos := []Todo{...}`
  - Creates and fills a slice that holds `Todo`s

- `Title: "Write Intro to Go (part 1)",`
  - Sets the `Title` field on a the `Todo`
  - Same for `Done`

- `todos[2].Done = true`
  - Set the 3rd Todo to be `Done`
  - Struct fields are accessed using dot notation

- `todos = append(todos, Todo{...})`
  - Creates and appends a `Todo` to the `todos` slice

- If a struct field is not set, it defaults to the zero value

---

### [08: Functions](code/08/main.go) - [Playground](https://play.golang.org/p/Zi1WSIjuU7U)

- Using a struct to structure a TODO item is much better

- `func printTodos(todos []Todo) {...}`
  - Creates a new function with the name `printTodos`, which accepts a parameter named `todos` which is a slice of `Todo`, and returns nothing

- `printTodos(todos)`
  - Calls the `printTodos` function, and passes in the `Todo` slice

- `printTodosFn := printTodos`
  - Functions are values, so you can store a function reference in a variable

- `newTodo := func(title string, done bool) Todo {...}`
  - Creates an anonymous function with two parameters:
    - `title` is a `string`
    - `done` is a `bool`
  - Returns a `Todo`

- `return Todo{...}`
  - Creates a new `Todo` and returns it

---

### [09: Functions & Errors](code/09/main.go) - [Playground](https://play.golang.org/p/yfCNH4ehKY-)

- `func printTodos(todos []Todo) (int, error) {...}`
  - Updates `printTodos` function to return two values:
    - `int` which contains the number of Todos printed
    - `error` which contains an error, if any, printing the Todos
    - The error is typically returned as the last value

- `return 0, errors.New("no todos to print")`
  - Returns a "no todos to print" error

- `numPrinted, err := printTodos(todos)`
  - Calls the `printTodos` function, and stores the returned values in `numPrinted` and `err`
  - The variable `err` is typically used to store variables

- `if err != nil {...}`
  - Checks if an error is present
  - Usually you'll just return the error if it is present, but you may also wrap it with more context, or do some clean up work

- `if _, err := printTodos([]Todo{}); err != nil {...}`
  - Calls the `printTodos` function with an empty slice, then checks if an error was returned
  - Since we don't need the number of Todos printed, we use the `_` (blank identifier) to ignore that value

---

### [10: Pointers](code/10/main.go) - [Playground](https://play.golang.org/p/-TXD4c8AIU8)

- Pointers are essential to Go, since by default values are copied and passed by value ([eg](https://play.golang.org/p/kjRNQ9ySTgh))
  - Except:
    - Slices, whose elements can be modified without explicit use of pointers ([eg](https://play.golang.org/p/-vY8xxAAJH-))
    - Maps, who are essentially pointers in the underlying implementation ([eg](https://play.golang.org/p/FEbvy5nE9KM))

- `func finishTodo(todo *Todo) {...}`
  - Creates a new function `finishTodo`, which has a parameter `todo` which is a pointer to a `Todo`
  - If you don't know what pointers are, I suggest you [read about them](https://en.wikipedia.org/wiki/Pointer_(computer_programming))
    - Basically, they store the location in memory, instead of the item itself

- `todo.Done = false`
  - Sets the `Done` to false on the `todo`
  - Notice that we can still use the dot notation even though it is a pointer; Go dereferences the pointer behind the scenes

- `todos := []*Todo{...}`
  - Updates the `Todo` slice to be a slice of pointers to `Todo`s

- `finishTodo(todos[2])`
  - Calls the `finishTodo` function with a pointer to the 3rd Todo

- `todo := *todos[0]`
  - Dereferences 1st the Todo, using `*`
  - This stores a copy in `todo`

- `finishTodo(&todo)`
  - Calls the `finishTodo` function with a pointer to the `todo` variable
  - To get the pointer (address) location, we use `&`
  - Notice that when printing `todo` it is Done, however, since we made a copy by dereferencing it first, the original Todo has not been updated

---

### [11: Methods](code/11/main.go) - [Playground](https://play.golang.org/p/fHNeFLVVorm)

- So far, so good, but it feels strange to define a `finishTodo` function and pass in a `Todo`
  - This is why Go allows methods on types

- `func (t *Todo) finish() {...}`
  - Creates the `finish` method on the `Todo` struct
  - Just like a function, you can pass in parameters and return multiple values
  - Methods can have pointer receivers (as shown here) or non-pointer receivers

- `todos[2].finish()`
  - Calls the `finish` method on the 3rd Todo

---

### [12: Packages](code/12/)

- So far, everything has been in the `main` package
  - While this is fine for a tiny program, to scale out, we'll want to utilize other packages
  - Let's create the `todo` package

- The `todo` directory is where the `todo` package lives
  - The directory name typically matches the package name

- Public items are specified with an uppercase letter (eg. types, functions, methods, struct fields, etc.)
  - Private items are specified with a lowercase letter
    - Private items are accessible at a package level

- [`todo/todo.go`](code/12/todo/todo.go)
  - `func (t *Todo) Finish() {...}`
    - The `Finish` method is now uppercase, since we want to access it from outside the package

- [`todo/group.go`](code/12/todo/group.go)
  - Added the `Group` type, which holds a group of `Todo`s
  - Overall, it should be familiar
  - `todos []*Todo`
    - The `Group` struct field `todos` is private; it is not accessible outside the package
  - `g.Todos = append(g.Todos, todos...)`
    - `todos...` spreads the slice of `*Todo` for use with the builtin `append` function

- [`main.go`](code/12/main.go)
  - `import "github.com/JosiahWitt/go-intro/part1/code/12/todo"`
    - Imports the `todo` package
    - Packages are referenced by full path, which typically includes the domain
  - `todo, err := group.Get(3)`
    - Calls the `Get` method on the `group`
  - `todo.Finish()`
    - Calls the `Finish` method on the 3rd `Todo`
  - `group.Print()`
    - Calls the `Print` method on the `group`

---

### [13: Custom Types](code/13/)

- [`todo/todos.go`](code/13/todo/todos.go)
  - `type Todos []*Todo`
    - Defines the type `Todos` to be a slice of `*Todo`
    - These types can have methods as well
    - So you could define a custom `int` type (`type MyInt int`) with an `Abs` method

- [`todo/group.go`](code/13/todo/group.go)
  - `Todos Todos`
    - Define the `Todos` field in struct `Group` to be of type `Todos`
  - `for _, todo := range g.Todos {...}`
    - Since `Todos` is an underlying slice, we can still `range` over it

- [`main.go`](code/13/main.go)
  - `group.Append(todo.Todos{...})`
    - We can now use `todo.Todos` instead of `[]*todo.Todo`

---

### [14: Interfaces](code/14/)

- In Go, interfaces are ad-hoc; they can be created anywhere, and as long as the type implements the methods requested on the interface, it is said to satisfy the interface

- `type TodoGroup interface {...}`
  - We define the `TodoGroup` interface by providing the function signatures that the type must implement to satisfy the interface

- `func createGroup() TodoGroup {}`
  - We create a new function `createGroup` that returns a type that satisfies the `TodoGroup` interface

- `return todo.NewGroup("My TODOs")`
  - Since `Group` satisfies the interface, it can be returned
  - Notice that `Group` has more methods than those on the interface, so in order to access them (or the struct fields), you need to use type assertion

- `originalGroup, ok := group.(*todo.Group)`
  - Use type assertion to assert that the type underlying the interface is a pointer to a `Group`
  - `ok` is a `bool` that indicates if the type assertion is true
  - `originalGroup` is of type `*Group`

- If you need to accept any type, you use the empty interface as your type: `interface{}`

---

### [14: Conclusion](code/15/)

- We've covered a lot!

- I've cleaned up the code in preparation for the next part.
