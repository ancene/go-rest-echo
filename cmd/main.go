package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ancene/go-rest-echo/internal/domain"
	"github.com/ancene/go-rest-echo/internal/infrastructure"
	"github.com/ancene/go-rest-echo/internal/infrastructure/connection"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	// _ "github.com/go-sql-driver/mysql" // uncomment this if using mysql
)

func main() {
	/********** ********** ********** ********** **********/
	/* Configuration
	/********** ********** ********** ********** **********/
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	cfg := domain.NewConfiguration()

	/********** ********** ********** ********** **********/
	/* Connection
	/* ex: mysql, etc., anything resouces what your app need
	/********** ********** ********** ********** **********/
	dbPG, err := connection.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	/********** ********** ********** ********** **********/
	/* HTTP Server on goroutine
	/********** ********** ********** ********** **********/
	e := infrastructure.NewServer(cfg, dbPG)
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	/********** ********** ********** ********** **********/
	/* make channel to receive signal
	/********** ********** ********** ********** **********/
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	/********** ********** ********** ********** **********/
	/* defer resources
	/********** ********** ********** ********** **********/
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer dbPG.Close()

	/********** ********** ********** ********** **********/
	/* shutdown server
	/********** ********** ********** ********** **********/
	println("💥 shutdown server ...")
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
