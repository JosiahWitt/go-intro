package main

import "fmt"

type Todo struct {
	Title string
	Done  bool
}

func main() {
	todos := []Todo{
		{
			Title: "Write Intro to Go (part 1)",
			Done:  false,
		},
		{
			Title: "Listen to the Go Time podcast",
			Done:  true,
		},
		{
			Title: "Read the Golang Weekly email",
			Done:  true,
		},
		{
			Title: "Read the Go 1.13 release notes",
			Done:  false,
		},
	}

	printTodos(todos)

	printTodosFn := printTodos
	printTodosFn(todos)

	newTodo := func(title string, done bool) Todo {
		return Todo{
			Title: title,
			Done:  done,
		}
	}

	todos = append(todos, newTodo("Buy a Go Gopher plush", false))
	printTodos(todos)
}

func printTodos(todos []Todo) {
	fmt.Printf("%d TODOs:\n", len(todos))
	for _, todo := range todos {
		checked := " "
		if todo.Done {
			checked = "x"
		}

		fmt.Printf("[%s] %s\n", checked, todo.Title)
	}
}
