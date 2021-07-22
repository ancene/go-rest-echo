package mysqldb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ancene/go-rest-echo/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewTodoRepository(t *testing.T) {
	mockSQL := new(sql.DB)
	use := NewTodoRepository(mockSQL)
	assert.NotNil(t, use)
}

func TestTodo_FindByID(t *testing.T) {
	ctx := context.Background()
	todo := domain.NewTodo(domain.TodoCreateDTO{
		Title:       "test",
		Description: "test",
	})

	t.Run("success", func(t *testing.T) {
		// initial mocking db
		db, mock, _ := sqlmock.New()
		defer func() { db.Close() }()

		// setup data to mock
		row := sqlmock.NewRows([]string{"id", "title", "description", "completed"}).
			AddRow(todo.ID, todo.Title, todo.Description, todo.Completed)

		// runnning mock
		query := fmt.Sprintf("SELECT id, title, description, completed FROM %s WHERE id = ?", todo.TableName())
		mock.ExpectQuery(query).WillReturnRows(row)

		// testing
		repo := NewTodoRepository(db)
		item, err := repo.FindByID(ctx, todo.ID)
		assert.NoError(t, err)
		assert.NotNil(t, item)
		assert.Equal(t, todo.ID, item.ID)
		assert.Equal(t, todo.Title, item.Title)
		assert.Equal(t, todo.Description, item.Description)
		assert.Equal(t, todo.Completed, item.Completed)

		// assertion mocking
		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// initial mocking db
		db, mock, _ := sqlmock.New()
		defer func() { db.Close() }()

		// runnning mock
		query := fmt.Sprintf("SELECT id, title, description, completed FROM %s WHERE id = ?", todo.TableName())
		mock.ExpectQuery(query).WillReturnError(errors.New("error from db"))

		// testing
		repo := NewTodoRepository(db)
		item, err := repo.FindByID(ctx, "not_found")
		assert.Error(t, err)
		assert.Nil(t, item)

		// assertion mocking
		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
