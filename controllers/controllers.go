package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/LavHinsu/gokvstore/keystore"
)

type key_struct struct {
	Keyname string
	Value   string
}

func GetKeyController(w http.ResponseWriter, req *http.Request) {
	key := req.PathValue("key")
	value := keystore.Getkey(key)
	if value != "404" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, value)
	} else {
		http.Error(w, "key not found", http.StatusNotFound)
	}
}

func PostKeyController(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	} else {
		var kvpair key_struct
		err = json.Unmarshal(body, &kvpair)
		if err != nil {
			log.Println(err)
		}
		// log.Print(string(body))
		status := keystore.Addkey(kvpair.Keyname, kvpair.Value)

		if status == 200 {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "200 ok")
		} else if status == 409 {
			http.Error(w, "Key Already exists, key name : "+kvpair.Keyname, http.StatusConflict)
		} else {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
		}
	}
}
