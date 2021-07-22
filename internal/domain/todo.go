package domain

import "github.com/ancene/go-rest-echo/pkg/nanoid"

type TodoCreateDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoUpdateDTO struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Completed   bool   `json:"completed,omitempty"`
}

type Todo struct {
	ID          string
	Title       string
	Description string
	Completed   bool
}

type Todos []*Todo

func NewTodo(dto TodoCreateDTO) *Todo {
	return &Todo{
		ID:          nanoid.NanoID(),
		Title:       dto.Title,
		Description: dto.Description,
		Completed:   false,
	}
}

func NewTodos(todos ...*Todo) Todos {
	result := make(Todos, 0)
	if len(todos) > 0 {
		result = append(result, todos...)
	}
	return result
}

func (Todo) TableName() string {
	return "todos"
}
