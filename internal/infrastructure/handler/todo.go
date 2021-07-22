package handler

import (
	"github.com/ancene/go-rest-echo/internal/application/port"
	"github.com/labstack/echo/v4"
)

type TodoHanlder struct {
	usecase port.TodoUsecase
}

func NewTodoHandler(usecase port.TodoUsecase) *TodoHanlder {
	return &TodoHanlder{
		usecase: usecase,
	}
}

func (th *TodoHanlder) GetByID(c echo.Context) error {
	paramID := c.Param("id")

	todo, err := th.usecase.GetByID(c.Request().Context(), paramID)
	if err != nil {
		return c.JSON(404, "not found")
	}

	return c.JSON(200, todo)
}
