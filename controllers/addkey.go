package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/LavHinsu/gokvstore/helpers"
	"github.com/LavHinsu/gokvstore/keystore"
)

type key_struct struct { // used for both adding and expiring keys
	Keyname string
	Value   string
	E_AT    int64
}

// this function is used to add a key to our map/kvstore
func PostKeyController(w http.ResponseWriter, req *http.Request) {
	ipaddr := helpers.ReadUserIP(req) // extract ip address from the req
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest) // return a badrequest response
		return                                                  // do not continue if the values could not be extracted
	} else {
		var kvpair key_struct // extract values from the request

		// unmarshall the request. this can fail, so we return a error if it does.
		err = json.Unmarshal(body, &kvpair)
		if err != nil {
			log.Println(err, "couldn't parse request json, bad data")
			http.Error(w, "bad request, couldn't parse json", http.StatusBadRequest)
			return // exit the function here if json parsing couldn't be completed
		}
		// add our key. the status will say if it was successfull or not.
		status := keystore.Addkey(kvpair.Keyname, kvpair.Value, kvpair.E_AT, ipaddr)

		// check how the operation went and respond accordingly.
		if status == 200 {
			w.WriteHeader(http.StatusOK)
			response := "200 ok"
			w.Write([]byte(response))
		} else if status == 409 {
			http.Error(w, "Key Already exists, key name : "+kvpair.Keyname, http.StatusConflict)
		} else {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
		}
	}
}
