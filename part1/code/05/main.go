package main

import "fmt"

func main() {
	todos := []string{
		"Write Intro to Go (part 1)",
		"Listen to the Go Time podcast",
		"Read the Golang Weekly email",
	}

	todos = append(todos, "Read the Go 1.13 release notes", "WIP", "Hidden TODO")
	todos[4] = "Buy a Go Gopher plush"

	todo := todos[0]
	fmt.Println("First todo:", todo)

	fmt.Printf("%d TODOs (showing first 5):\n", len(todos))
	firstFive := todos[:5]
	for i, todo := range firstFive {
		fmt.Printf("%d: %s\n", i+1, todo)
	}
}
