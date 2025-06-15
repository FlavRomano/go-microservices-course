package main

import (
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestBody mailMessage
	err := app.readJSON(w, r, &requestBody)
	if err != nil {
		log.Println("Error received on readJSON:", err)
		app.errorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestBody.From,
		To:      requestBody.To,
		Subject: requestBody.Subject,
		Data:    requestBody.Message,
	}

	err = app.Mailer.sendSMTPMessage(msg)
	if err != nil {
		log.Println("Error received on sendSMTPMessage:", err)
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "sent to " + requestBody.To,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
