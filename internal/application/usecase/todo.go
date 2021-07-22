package usecase

import (
	"context"
	"errors"

	"github.com/ancene/go-rest-echo/internal/application/port"
	"github.com/ancene/go-rest-echo/internal/domain"
)

type todoUsecase struct {
	repo port.TodoRepository
}

func NewTodoUsecase(repo port.TodoRepository) port.TodoUsecase {
	return &todoUsecase{
		repo: repo,
	}
}

func (tu *todoUsecase) GetByID(ctx context.Context, id string) (*domain.Todo, error) {
	todo, err := tu.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("something wrong")
	}
	return todo, nil
}
