package controllers

import (
	"fmt"
	"net/http"

	"github.com/LavHinsu/gokvstore/helpers"
	"github.com/LavHinsu/gokvstore/keystore"
)

// this function is used to delete a key in our map/kvstore
func DeleteKeyController(w http.ResponseWriter, req *http.Request) {
	ipaddr := helpers.ReadUserIP(req)
	key := req.PathValue("key")
	value := keystore.DeleteKey(key, ipaddr)
	if value != 404 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "200 ok")
	} else {
		http.Error(w, "key not found", http.StatusNotFound)
	}
}
