package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/LavHinsu/gokvstore/keystore"
)

func getKey(w http.ResponseWriter, req *http.Request) {
	key := req.PathValue("id")
	value := keystore.Getkey(key)
	fmt.Fprint(w, value)
}

func postKey(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	} else {
		log.Print(string(body))
	}
	keystore.Addkey()
	fmt.Fprint(w, "200 ok")
}

func healthCheck(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "200 ok")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthcheck", healthCheck) // healthcheck for a probe if this is run in kubernetes
	mux.HandleFunc("GET /key/{id}", getKey)
	mux.HandleFunc("POST /writekey", postKey)

	http.ListenAndServe("localhost:8080", mux)
}
