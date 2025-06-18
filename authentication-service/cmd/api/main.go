package main

import (
	"authentication-service/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "8080"

var counts int64

type Config struct {
	Repository data.Repository
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *Config) setupRepo(conn *sql.DB) {
	db := data.NewPostgresRepository(conn)
	app.Repository = db
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		conn, err := openDB(dsn)

		if err != nil {
			log.Println("Db not yet ready...")
			counts++
		} else {
			log.Println("Connected to Db")
			return conn
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for 2 sec.")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	log.Println("Starting authentication-service")

	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to DB")
	}

	app := Config{}
	app.setupRepo(conn)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
