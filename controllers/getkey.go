package controllers

import (
	"fmt"
	"net/http"

	"github.com/LavHinsu/gokvstore/keystore"
)

// this function is used to get a key from our map/kvstore
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
