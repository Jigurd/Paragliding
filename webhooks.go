package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "strings"
)

//Webhook stores all information about a given webhook
type Webhook struct {
	WebookURL       string `json:"webhookURL"`
	MinTriggerValue string `json:"minTriggerValue"`
}

//WebhookWrapper stores timestamp data used to ping webhooks, as well as the ID (the timestamp for creation)
type WebhookWrapper struct {
	Timestamp int64 `json:"timestamp"` //the timestamp is also used as an ID
	HookStop  int64 `json:"hookstop"`   //is used for finding out what tracks are new to the hook
	Webhook
}

//HandlerWebhook handles webhook requests.
func HandlerWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var newWrapper WebhookWrapper
		var newHook Webhook
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&newHook)

		if newHook.WebookURL == "" { //if the json supplied no webURL, throw bad request error
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if newHook.MinTriggerValue == "" { //set mintriggervalue if none was set in the request
			newHook.MinTriggerValue = "1"
		}

		newWrapper.Webhook = newHook //add the new hook to a wrapper

		newWrapper.Timestamp = Millisec() //set both times in the wrapper to the current time
		newWrapper.HookStop = newWrapper.Timestamp

		db.AddHook(newWrapper)
		json.NewEncoder(w).Encode(newWrapper.Timestamp) //return the ID of the hook
	} else if r.Method == "GET" || r.Method == "DELETE" {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) == 6 || (len(parts) == 7 && parts[6] == "") {
			requestedID, err := strconv.ParseInt(parts[5], 10, 64)

			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			webhook, err := db.GetHook(requestedID)

			if r.Method == "DELETE" {
				err = db.DeleteHook(requestedID)
			}

			if err != nil {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

				return
			}
			json.NewEncoder(w).Encode(webhook.Webhook)

		} else { //if the URL is invalid, return not found
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}

	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

}

//PingWebhooks notifies appropriate webhooks. It is called when a track is added.
func PingWebhooks() {
    webhooks := db.GetAllHooks()
    for i:=0;i<len(webhooks);i++{
        str:=fmt.Sprintf(`{"content": "A new track was added!"}`)
        http.Post(webhooks[i].Webhook.WebookURL, "application/json", bytes.NewBufferString(str))
    }
}

//func PingWebhooks() {
//	webhooks := db.GetAllHooks()
//	tracks := db.GetAll()
//
//	for i := 0; i < len(webhooks); i++ {
//		var count int64
//		var ids []int64
//		startTime := Millisec()
//		count = 0
//		curHookStop := webhooks[i].HookStop //the HookStop of the Webhook currently being treated
//
//		for j := 0; j < (len(tracks) - 1); j++ {
//			if curHookStop < tracks[i].Timestamp {
//				count++
//				ids = append(ids, tracks[j].Timestamp)
//			}
//		}
//
//		triggerval, _ := strconv.ParseInt(webhooks[i].Webhook.MinTriggerValue, 10, 64)
//
//		curHookStop = tracks[len(tracks)-1].Timestamp
//
//		if count >= triggerval {
//			stopTime := Millisec() - startTime
//			str := fmt.Sprintf(
//				`{"content":"%v is the latest timestamp. %v are new tracks. Processing time: %v ms. Count:%v"}`,
//				curHookStop, ids, stopTime, count)
//
//			http.Post(webhooks[i].Webhook.WebookURL, "application/json", bytes.NewBufferString(str))
//		}
//
//	}
//
//}
