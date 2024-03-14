package store

type Todo struct {
	Id     int
	Title  string
	Body   string
	Author string
}

type TodoStore interface {
	CreateUser(title string, body string, author string) error
	GetTodos(author string) ([]*Todo, error)
	GetTodo(id int) (*Todo, error)
}
