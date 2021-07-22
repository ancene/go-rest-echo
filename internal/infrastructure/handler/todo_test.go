package handler

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/ancene/go-rest-echo/internal/domain"
	"github.com/ancene/go-rest-echo/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewTodoHanlder(t *testing.T) {
	mockRepo := new(mocks.TodoUsecase)
	hand := NewTodoHandler(mockRepo)
	assert.NotNil(t, hand)
}

func TestTodo_GetByID(t *testing.T) {
	id := "id888"

	t.Run("success", func(t *testing.T) {
		// setup data
		todo := &domain.Todo{
			ID:          id,
			Title:       "test title",
			Description: "example test decription",
			Completed:   false,
		}
		temp := domain.Todo{}

		// setup test server environment
		e := echo.New()
		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/todos/:id")
		c.SetParamNames("id")
		c.SetParamValues(id)

		// setup mocking
		mockUc := new(mocks.TodoUsecase)
		mockUc.On("GetByID", c.Request().Context(), id).Return(todo, nil)

		// testing
		hand := NewTodoHandler(mockUc)
		err := hand.GetByID(c)
		assert.NoError(t, err)

		err = json.Unmarshal(rec.Body.Bytes(), &temp)
		assert.NoError(t, err)

		assert.NoError(t, err)
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, todo.ID, temp.ID)

		mockUc.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		// setup test server environment
		e := echo.New()
		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/todos/:id")
		c.SetParamNames("id")
		c.SetParamValues(id)

		// setup mocking
		mockUc := new(mocks.TodoUsecase)
		mockUc.On("GetByID", c.Request().Context(), id).Return(nil, errors.New("any error"))

		// testing
		hand := NewTodoHandler(mockUc)
		err := hand.GetByID(c)
		assert.NoError(t, err)

		assert.Equal(t, 404, rec.Code)
		assert.Equal(t, "\"not found\"\n", rec.Body.String())

		mockUc.AssertExpectations(t)
	})
}
