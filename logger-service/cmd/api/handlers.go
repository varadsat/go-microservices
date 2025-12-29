package main

import (
	"logger/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPayload
	_ = app.readJSON(w, r, &requestPayload)

	logEvent := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(logEvent)
	if err != nil {
		app.errorJson(w, err)
		return
	}
	resp := jsonResponse{
		Error:   false,
		Message: "logged by logger service",
	}
	app.writeJson(w, http.StatusAccepted, resp)
}
