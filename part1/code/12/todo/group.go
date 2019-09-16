package todo

import (
	"errors"
	"fmt"
)

type Group struct {
	Title string
	todos []*Todo
}

func NewGroup(title string) *Group {
	return &Group{
		Title: title,
	}
}

func (g *Group) Append(todos []*Todo) {
	g.todos = append(g.todos, todos...)
}

func (g *Group) Get(num int) (*Todo, error) {
	i := num - 1
	if i < 0 || i >= len(g.todos) {
		return nil, errors.New("todo does not exist")
	}

	return g.todos[i], nil
}

func (g *Group) Print() {
	fmt.Printf("%s: %d TODOs:\n", g.Title, len(g.todos))
	for _, todo := range g.todos {
		checked := " "
		if todo.Done {
			checked = "x"
		}

		fmt.Printf("[%s] %s\n", checked, todo.Title)
	}
}
