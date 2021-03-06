package connection

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ancene/go-rest-echo/internal/domain"
)

func NewPostgresConnection(cfg *domain.Configuration) (*sql.DB, error) {
	if cfg == nil {
		return nil, errors.New("configuration is nil")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Database)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
