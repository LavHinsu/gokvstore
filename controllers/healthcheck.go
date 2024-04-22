package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LavHinsu/gokvstore/keystore"
)

// healthcheck function for confirming that the server is running and return the number of keys present
func HealthCheck(w http.ResponseWriter, req *http.Request) {
	keycount := keystore.GetKeyCount()
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "total keys: "+strconv.Itoa(keycount))
}
