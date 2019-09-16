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
			Done:  false,
		},
		{
			Title: "Read the Go 1.13 release notes",
			Done:  false,
		},
	}

	todos[2].Done = true

	todos = append(todos, Todo{
		Title: "Buy a Go Gopher plush",
		Done:  false,
	})

	fmt.Printf("%d TODOs:\n", len(todos))
	for _, todo := range todos {
		checked := " "
		if todo.Done {
			checked = "x"
		}

		fmt.Printf("[%s] %s\n", checked, todo.Title)
	}
}
