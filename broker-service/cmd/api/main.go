package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "8080"

type Config struct {
	Rabbit *amqp.Connection
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			log.Println("Still can't connect to rabbitmq")
			counts++
		} else {
			connection = c
			break
		}

		if counts > 5 {
			log.Println("Can't connect, give up after 5 tries. Caused by", err)
			return nil, err
		}

		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		time.Sleep(backoff)

	}
	return connection, nil
}

func main() {

	// try to connect to rabbitmq
	log.Println("Connecting to rabbitmq")
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		return
	}
	defer rabbitConn.Close()
	log.Println("Connected to rabbitmq")

	app := Config{
		Rabbit: rabbitConn,
	}

	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start server
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
