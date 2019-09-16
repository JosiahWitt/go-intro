package todo

type Todo struct {
	Title string
	Done  bool
}

func (t *Todo) Finish() {
	t.Done = true
}
