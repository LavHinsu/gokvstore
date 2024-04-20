package kvstoreserver

import (
	"fmt"
	"net/http"

	"github.com/LavHinsu/gokvstore/controllers"
)

func healthCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "200 ok")
}

func Kvstoreserver() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthcheck", healthCheck)                     // healthcheck for a probe if this is run in kubernetes
	mux.HandleFunc("GET /key/{key}", controllers.GetKeyController)      // get a key
	mux.HandleFunc("POST /writekey", controllers.PostKeyController)     // add a key
	mux.HandleFunc("PATCH /updatekey", controllers.UpdateKeyController) // update a key

	http.ListenAndServe("localhost:8080", mux)
}
