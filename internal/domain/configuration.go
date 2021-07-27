package domain

import (
	"os"
	"strconv"
)

type app struct {
	Host string
	Port int
	Name string
}

type pg struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type Configuration struct {
	App      app
	Postgres pg
	Mysql    mysql
}

func NewConfiguration() *Configuration {
	ap, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		ap = 8080
	}

	pgp, err := strconv.Atoi(os.Getenv("PG_PORT"))
	if err != nil {
		ap = 5432
	}

	myp, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		ap = 3306
	}

	return &Configuration{
		App: app{
			Host: os.Getenv("APP_HOST"),
			Port: ap,
			Name: os.Getenv("APP_NAME"),
		},
		Postgres: pg{
			Host:     os.Getenv("PG_HOST"),
			Port:     pgp,
			Username: os.Getenv("PG_USERNAME"),
			Password: os.Getenv("PG_PASSWORD"),
			Database: os.Getenv("PG_DATABASE"),
		},
		Mysql: mysql{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     myp,
			Username: os.Getenv("MYSQL_USERNAME"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Database: os.Getenv("MYSQL_DATABASE"),
		},
	}
}
