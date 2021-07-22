package mysqldb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ancene/go-rest-echo/internal/application/port"
	"github.com/ancene/go-rest-echo/internal/domain"
)

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) port.TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (tr *todoRepository) FindByID(ctx context.Context, id string) (*domain.Todo, error) {
	todo := domain.Todo{}
	query := fmt.Sprintf("SELECT id, title, description, completed FROM %s WHERE id = ?", todo.TableName())

	row := tr.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
