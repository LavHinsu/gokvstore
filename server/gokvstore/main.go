package main

import (
	"fmt"
	"net/http"

	"github.com/LavHinsu/gokvstore/keystore"
)

func getKey(w http.ResponseWriter, req *http.Request) {
	idString := req.PathValue("id")
	keystore.Keystore()
	fmt.Fprint(w, "Got ID ", idString)
}

func postKey(w http.ResponseWriter, req *http.Request) {
	// idString := req.PathValue("id") todo: add mechanism to extract data from post req and put in memory
	fmt.Fprint(w, "200 ok")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /key/{id}", getKey)
	mux.HandleFunc("POST /writekey/{id}", postKey)
	http.ListenAndServe("localhost:8080", mux)
}
