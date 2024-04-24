package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/LavHinsu/gokvstore/helpers"
	"github.com/LavHinsu/gokvstore/keystore"
)

// this function is used to update a key in our map/kvstore
func UpdateKeyController(w http.ResponseWriter, req *http.Request) {
	ipaddr := helpers.ReadUserIP(req) // extract ip address from the req
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest) // if bad body is sent
		return
	} else {
		var kvpair key_struct
		err = json.Unmarshal(body, &kvpair)
		if err != nil { // if bad json is sent
			log.Println("couldn't parse request json, bad data")
			http.Error(w, "bad request, couldn't parse json", http.StatusBadRequest)
			return // exit the function here if json parsing couldn't be completed
		}
		status := keystore.UpdateKey(kvpair.Keyname, kvpair.Value, ipaddr)

		if status == 200 {
			w.WriteHeader(http.StatusOK)
			response := "200 ok"
			w.Write([]byte(response))
		} else if status == 404 {
			http.Error(w, "Key doesn't exist, key: "+kvpair.Keyname, http.StatusNotFound)
		} else {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
		}
	}
}
