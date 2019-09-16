package main

import (
	"errors"
	"fmt"
)

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
		{
			Title: "Buy a Go Gopher plush",
			Done:  false,
		},
	}

	numPrinted, err := printTodos(todos)
	if err != nil {
		fmt.Println("There was an error printing the todos:", err)
	}
	fmt.Println("Printed", numPrinted)

	if _, err := printTodos([]Todo{}); err != nil {
		fmt.Println("There was an error printing the todos:", err)
	}
}

func printTodos(todos []Todo) (int, error) {
	if len(todos) == 0 {
		return 0, errors.New("no todos to print")
	}

	fmt.Printf("%d TODOs:\n", len(todos))
	for _, todo := range todos {
		checked := " "
		if todo.Done {
			checked = "x"
		}

		fmt.Printf("[%s] %s\n", checked, todo.Title)
	}

	return len(todos), nil
}
