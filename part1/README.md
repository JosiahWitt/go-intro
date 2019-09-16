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

`func finishTodo(todo *Todo) {...}`
  - Creates a new function `finishTodo`, which has a parameter `todo` which is a pointer to a `Todo`
  - If you don't know what pointers are, I suggest you [read about them](https://en.wikipedia.org/wiki/Pointer_(computer_programming))
    - Basically, they store the location in memory, instead of the item itself

`todo.Done = false`
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
