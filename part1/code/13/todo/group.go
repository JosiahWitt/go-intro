package todo

import "fmt"

type Group struct {
	Title string
	Todos Todos
}

func NewGroup(title string) *Group {
	return &Group{
		Title: title,
	}
}

func (g *Group) Append(todos Todos) {
	g.Todos = append(g.Todos, todos...)
}

func (g *Group) Print() {
	fmt.Printf("%s: %d TODOs:\n", g.Title, len(g.Todos))
	for _, todo := range g.Todos {
		checked := " "
		if todo.Done {
			checked = "x"
		}

		fmt.Printf("[%s] %s\n", checked, todo.Title)
	}
}
