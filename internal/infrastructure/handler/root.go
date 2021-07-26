package handler

import (
	"github.com/ancene/go-rest-echo/internal/domain"
	"github.com/ancene/go-rest-echo/internal/infrastructure/serialization"
	"github.com/labstack/echo/v4"
)

type RootHanlder struct {
	config *domain.Configuration
}

func NewRootHandler(config *domain.Configuration) *RootHanlder {
	return &RootHanlder{
		config: config,
	}
}

func (rh *RootHanlder) Home(c echo.Context) error {
	return c.JSON(200, serialization.OK("welcome to ancene", []string{}))
}

func (rh *RootHanlder) Health(c echo.Context) error {
	data := map[string]interface{}{
		"http": "good",
		"connection": map[string]string{
			"mysql":    "good",
			"postgres": "good",
		},
		"memory": "good",
	}
	return c.JSON(200, serialization.OK(rh.config.App.Name+" is healhty", data))
}
