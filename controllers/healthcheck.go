package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LavHinsu/gokvstore/keystore"
)

// healthcheck function for confirming that the server is running and return the number of keys present
func HealthCheck(w http.ResponseWriter, req *http.Request) {
	keycount := keystore.GetKeyCount()
	response, err := json.Marshal(keycount)
	if err != nil {
		log.Panic("failed marshalling a response in healtcheck, time to panic and exit, UwU!", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	}
}
