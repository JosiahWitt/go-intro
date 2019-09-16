package main

import "fmt"

func main() {
	todos := map[string]bool{
		"Write Intro to Go (part 1)":    false,
		"Listen to the Go Time podcast": true,
		"Read the Golang Weekly email":  true,
	}

	todos["Read the Go 1.13 release notes"] = false
	todos["Buy a Go Gopher plush"] = false

	if done, ok := todos["Write Intro to Go (part 1)"]; ok {
		fmt.Println("Write Intro to Go (part 1) exists and has status of", done)
	}

	fmt.Printf("%d TODOs:\n", len(todos))
	for todo, done := range todos {
		checked := " "
		if done {
			checked = "x"
		}

		fmt.Printf("[%s] %s\n", checked, todo)
	}
}
