package main

import (
	"fmt"

	"github.com/JosiahWitt/go-intro/part1/code/14/todo"
)

type TodoGroup interface {
	Print()
	Append(todo.Todos)
}

func main() {
	group := createGroup()

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
			Done:  false,
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

	originalGroup, ok := group.(*todo.Group)
	if ok {
		todo, err := originalGroup.Get(3)
		if err != nil {
			fmt.Println("TODO DNE", err)
			return
		}
		todo.Finish()
	}

	group.Print()
}

func createGroup() TodoGroup {
	return todo.NewGroup("My TODOs")
}
