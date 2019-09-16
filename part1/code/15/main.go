package main

import (
	"github.com/JosiahWitt/go-intro/part1/code/15/todo"
)

func main() {
	group := todo.NewGroup("My TODOs")

	group.Append(todo.Todos{
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
	})

	group.Print()
}
