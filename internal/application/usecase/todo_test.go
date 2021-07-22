package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/ancene/go-rest-echo/internal/domain"
	"github.com/ancene/go-rest-echo/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewTodoUsecase(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	use := NewTodoUsecase(mockRepo)
	assert.NotNil(t, use)
}

func TestTodo_GetByID(t *testing.T) {
	ctx := context.Background()
	id := "id80"

	t.Run("success", func(t *testing.T) {
		mockRepo := new(mocks.TodoRepository)
		use := NewTodoUsecase(mockRepo)

		mockRepo.On("FindByID", ctx, id).Return(&domain.Todo{}, nil)

		todo, err := use.GetByID(ctx, id)
		assert.NoError(t, err)
		assert.NotNil(t, todo)

		mockRepo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		mockRepo := new(mocks.TodoRepository)
		use := NewTodoUsecase(mockRepo)

		mockRepo.On("FindByID", ctx, id).Return(nil, errors.New("error"))

		todo, err := use.GetByID(ctx, id)
		assert.Error(t, err)
		assert.NotNil(t, err)
		assert.Nil(t, todo)

		mockRepo.AssertExpectations(t)
	})
}
