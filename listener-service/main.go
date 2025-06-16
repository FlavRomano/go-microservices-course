package main

import (
	"listener-service/event"
	"log"
	"math"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func connect() (*amqp.Connection, error) {
	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	for {
		rabbitmqUrl := os.Getenv("RABBITMQ_URL")
		c, err := amqp.Dial(rabbitmqUrl)
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

func readTopics() []string {
	topicsString := os.Getenv("TOPICS")
	return strings.Split(topicsString, ",")
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

	// start listening for messages
	log.Println("Listening for and consuming rabbitmq messages...")

	// create a consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	// observe the queue and consume the event
	err = consumer.Listen(readTopics())
	if err != nil {
		log.Println(err)
	}
}
