package main

import (
	"encoding/json"
	"net/http"
    "strconv"
    "strings"
)

//Webhook stores all information about a given webhook
type Webhook struct {
	WebookURL       string `json:"webhookURL"`
	MinTriggerValue string `json:"minTriggerValue"`
	Timestamp       int64  `json: "timestamp"` //the timestamp is also used as an ID
}

//HandlerWebhook handles webhook requests.
func HandlerWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var newHook Webhook
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&newHook)

		if newHook.WebookURL==""{ //if the json supplied no webURL, throw bad request error
            http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
            return
        }

        if newHook.MinTriggerValue==""{ //set mintriggervalue if none was set in the request
            newHook.MinTriggerValue="1"
        }

		newHook.Timestamp = Millisec()
		db.AddHook(newHook)
		json.NewEncoder(w).Encode(newHook.Timestamp) //return the ID of the hook
	} else if r.Method == "GET" || r.Method == "DELETE" {
        parts := strings.Split(r.URL.Path, "/")
        if len(parts) == 6 || (len(parts)==7&&parts[6]=="") {
            requestedID, err := strconv.ParseInt(parts[5], 10, 64)

            if err != nil {
                http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
                return
            }

                webhook, err := db.GetHook(requestedID)


             if r.Method == "DELETE"{
                 err = db.DeleteHook(requestedID)
                }

             if err != nil{
                 http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

                 return
             }
            json.NewEncoder(w).Encode(webhook)

        }else{ //if the URL is invalid, return not found
            http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
        }


	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

}


//PingWebhook notifies appropriate webhooks. It is called when a track is added.