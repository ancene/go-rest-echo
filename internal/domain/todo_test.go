package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTodo(t *testing.T) {
	createDto := TodoCreateDTO{}
	todo := NewTodo(createDto)

	assert.NotNil(t, todo)
	assert.Equal(t, 11, len(todo.ID))
	assert.Equal(t, "", todo.Title)
	assert.Equal(t, "", todo.Description)
	assert.False(t, todo.Completed)
}

func TestNewTodos(t *testing.T) {
	todos := NewTodos()
	assert.Empty(t, todos)

	todos = NewTodos(&Todo{})
	assert.NotEmpty(t, todos)
	assert.Equal(t, 1, len(todos))
}

func TestTodo_TableName(t *testing.T) {
	// way 1
	todo := Todo{}
	assert.Equal(t, "todos", todo.TableName())

	// way 2
	todoPtr := &Todo{}
	assert.Equal(t, "todos", todoPtr.TableName())
}
