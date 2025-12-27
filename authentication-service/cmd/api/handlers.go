package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {

	// requestPayload defn
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// marshal request into requestpayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}
	//validate user and password
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJson(w, errors.New("Invalid credentials"), http.StatusUnauthorized)
		return
	}
	match, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !match {
		app.errorJson(w, errors.New("Invalid credentials"), http.StatusUnauthorized)
		return
	}
	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}
	app.writeJson(w, http.StatusAccepted, payload)
}
