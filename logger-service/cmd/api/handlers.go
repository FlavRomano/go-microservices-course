package main

import (
	"encoding/json"
	"log"
	"logger-service/cmd/data"
	"net/http"
)

type LogRequest struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var payload LogRequest
	_ = app.readJSON(w, r, &payload)

	event := data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	}

	message, _ := json.Marshal(payload)
	log.Println("Logging payload:", string(message))
	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Log posted",
	}
	log.Println("Log posted")

	app.writeJSON(w, http.StatusAccepted, resp)
}
