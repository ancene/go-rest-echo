package infrastructure

import (
	"database/sql"
	"time"

	"github.com/ancene/go-rest-echo/internal/application/usecase"
	"github.com/ancene/go-rest-echo/internal/domain"
	"github.com/ancene/go-rest-echo/internal/infrastructure/handler"
	"github.com/ancene/go-rest-echo/internal/infrastructure/middleware"
	"github.com/ancene/go-rest-echo/internal/infrastructure/repository/mysqldb"
	"github.com/labstack/echo/v4"
)

func NewServer(cfg *domain.Configuration, db *sql.DB) *echo.Echo {
	/** +----- initialize server -----+ **/
	e := echo.New()
	e.HideBanner = true
	e.Server.ReadTimeout = 30 * time.Second
	e.Server.WriteTimeout = 30 * time.Second

	/** +----- middlewares -----+ **/
	e.Use(middleware.Logger(), middleware.Recovery())
	e.Use(middleware.CORS(), middleware.Secure())
	e.Use(middleware.BodyLimit(), middleware.Gzip())

	/** +----- ioc -----+ **/
	repoTodo := mysqldb.NewTodoRepository(db)
	useTodo := usecase.NewTodoUsecase(repoTodo)
	handTodo := handler.NewTodoHandler(useTodo)

	/** +----- Resouces routes /todos -----+ **/
	e.GET("/todos/:id", handTodo.GetByID)

	return e
}
