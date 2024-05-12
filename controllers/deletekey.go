package controllers

import (
	"net/http"

	"github.com/LavHinsu/gokvstore/helpers"
	"github.com/LavHinsu/gokvstore/keystore"
)

// this function is used to delete a key in our map/kvstore (this is used for the httpServer)
func DeleteKeyController(w http.ResponseWriter, req *http.Request) {
	ipaddr := helpers.ReadUserIP(req) // extract ip address from the req
	key := req.PathValue("key")
	value := keystore.DeleteKey(key, ipaddr)
	if value != 404 {
		w.WriteHeader(http.StatusOK)
		response := "200 ok"
		w.Write([]byte(response))
	} else {
		http.Error(w, "key not found", http.StatusNotFound)
	}
}
