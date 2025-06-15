package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	webPort = 8080
)

type Config struct {
	Mailer Mail
}

func createMail() Mail {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	m := Mail{
		Domain:     os.Getenv("MAIL_DOMAIN"),
		Host:       os.Getenv("MAIL_HOST"),
		Port:       port,
		Username:   os.Getenv("MAIL_USERNAME"),
		Password:   os.Getenv("MAIL_PASSWORD"),
		Encryption: os.Getenv("MAIL_ENC"),
		FromName:   os.Getenv("MAIL_FROM_NAME"),
		FromAddr:   os.Getenv("MAIL_FROM_ADDR"),
	}

	return m
}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: app.routes(),
	}

	log.Println("Mail service listening on", webPort)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("Can't listen and serve:", err)
	}
}

func main() {
	app := Config{
		Mailer: createMail(),
	}

	app.serve()
}
