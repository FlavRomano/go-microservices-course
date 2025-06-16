package main

import (
	"context"
	"log"
	"logger-service/cmd/data"
	"time"
)

type RPCServer struct {
}

type RPCPayload struct {
	Name string
	Data string
}

// any method that I want to expose via RPC needs to be public (capital letter)
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	log.Println("Logging through RPC now")
	collection := mongoClient.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo")
		return err
	}

	*resp = "Processed payload via RPC"
	return nil
}
