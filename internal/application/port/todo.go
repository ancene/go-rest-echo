package port

import (
	"context"

	"github.com/ancene/go-rest-echo/internal/domain"
)

type TodoUsecase interface {
	GetByID(ctx context.Context, id string) (*domain.Todo, error)
}

type TodoRepository interface {
	FindByID(ctx context.Context, id string) (*domain.Todo, error)
}
