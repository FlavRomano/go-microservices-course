package main

import (
	"context"
	"fmt"
	"log"
	"logger-service/cmd/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "8080"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

var mongoClient *mongo.Client

type Config struct {
	Models data.Models
}

func connectToMongo() *mongo.Client {
	connOptions := options.Client().ApplyURI(mongoURL)
	credentials := options.Credential{
		Username: "admin",
		Password: "password",
	}
	connOptions.SetAuth(credentials)

	conn, err := mongo.Connect(context.TODO(), connOptions)
	if err != nil {
		log.Panic("Can't connect to mongodb at", mongoURL, "due to", err)
	}

	log.Println("Connected to mongo at", mongoURL)
	return conn
}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Println("Logger listening on", webPort)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("Can't listen and serve:", err)
	}
}

func main() {
	client := connectToMongo()

	mongoClient = client

	// create context to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close conn
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	app.serve()
}
