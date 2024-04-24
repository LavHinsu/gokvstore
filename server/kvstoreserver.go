package kvstoreserver

import (
	"log"
	"net/http"

	"github.com/LavHinsu/gokvstore/controllers"
)

// start our server
func Kvstoreserver() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthcheck", controllers.HealthCheck)                // healthcheck for a probe if this is run in kubernetes
	mux.HandleFunc("GET /key/{key}", controllers.GetKeyController)             // get a key
	mux.HandleFunc("POST /writekey", controllers.PostKeyController)            // add a key
	mux.HandleFunc("PATCH /updatekey", controllers.UpdateKeyController)        // update a key
	mux.HandleFunc("DELETE /deletekey/{key}", controllers.DeleteKeyController) // delete a key

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Println("unable to start server,", err)
	}
}
