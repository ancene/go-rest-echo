package connection

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ancene/go-rest-echo/internal/domain"
)

func NewMysqlConnection(cfg *domain.Configuration) (*sql.DB, error) {
	if cfg == nil {
		return nil, errors.New("configuration is nil")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Mysql.Username, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
