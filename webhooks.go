package main

import (
	"encoding/json"
	"net/http"
)

//Webhook stores all information about a given webhook
type Webhook struct {
	WebookURL       string `json:"webhookURL"`
	MinTriggerValue string `json:"minTriggerValue"`
	Timestamp       int64  `json: "timestamp"`
}

//HandlerWebhook handles webhook requests.
func HandlerWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var newHook Webhook
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&newHook)

		newHook.Timestamp = Millisec()

		db.AddHook(newHook)
	} else if r.Method == "GET" {

	} else if r.Method == "DELETE" {

	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

}
