package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LavHinsu/gokvstore/helpers"
	"github.com/LavHinsu/gokvstore/keystore"
)

// this function is used to get a key from our map/kvstore
func GetKeyController(w http.ResponseWriter, req *http.Request) {
	ipaddr := helpers.ReadUserIP(req) // extract ip address from the req
	key := req.PathValue("key")
	value := keystore.Getkey(key, ipaddr)
	if value != nil {
		response, err := json.Marshal(value)
		if err != nil {
			log.Panic(err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}

	} else {
		http.Error(w, "key not found", http.StatusNotFound)
	}
}
