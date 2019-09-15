package main

import "fmt"

func main() {
	todos := []string{
		"Write Intro to Go (part 1)",
		"Listen to the Go Time podcast",
		"Read the Golang Weekly email",
	}

	todos = append(todos, "Read the Go 1.13 release notes", "Buy a Go Gopher plush")

	for i, todo := range todos {
		fmt.Printf("%d: %s\n", i+1, todo)
	}
}
