package domain

type Todo struct {
	Id TodoId `json:"id"`
	Title TodoTitle `json:"title"`
	Person TodoPerson `json:"person"`
	Done TodoDone `json:"done"`
}

type TodoId struct {
	Value int `json:"value"`
}

type TodoTitle struct {
	Value string `json:"value"`
}

type TodoPerson struct{
	Value string `json:"value"`
}

type TodoDone struct {
	Value bool `json:"value"`
}

type CreateTodo struct {
	Title TodoTitle
	Person TodoPerson
}

type UpdateTodo struct {
	Title TodoTitle `json:"title"`
	Person TodoPerson `json:"person"`
	Done TodoDone `json:"done"`
}

func NewTodo(title string, done bool) Todo {
	return Todo{
		Title: TodoTitle{Value: title},
		Done: TodoDone{Value: done},
	}
}