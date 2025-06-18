package main

import (
	"errors"
	"fmt"
	"net/http"
)

type UserDTO struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"Password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.log("authenticator", "Bad request on reading")
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Repository.GetByEmail(requestPayload.Email)
	if err != nil {
		app.log("authenticator", "Invalid credentials")
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
	}

	valid, err := app.Repository.PasswordMatches(requestPayload.Password, *user)
	if !valid || err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		app.log("authenticator", "Invalid credentials")
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data: UserDTO{
			Email:     user.Email,
			Firstname: user.FirstName,
			Lastname:  user.LastName,
		},
	}
	app.log("authenticator", "User logged successfully")
	app.writeJSON(w, http.StatusAccepted, payload)
}
