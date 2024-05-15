package kvstoreserver

import (
	"log"
	"net/http"
	"strconv"

	"github.com/LavHinsu/gokvstore/controllers"
)

// start our server
func Kvstoreserver(port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthcheck", controllers.HealthCheck)                // healthcheck for a probe if this is run in kubernetes
	mux.HandleFunc("GET /key/{key}", controllers.GetKeyController)             // get a key
	mux.HandleFunc("POST /writekey", controllers.PostKeyController)            // add a key
	mux.HandleFunc("PATCH /updatekey", controllers.UpdateKeyController)        // update a key
	mux.HandleFunc("DELETE /deletekey/{key}", controllers.DeleteKeyController) // delete a key

	serve := "localhost:" + strconv.Itoa(port)

	err := http.ListenAndServe(serve, mux)
	if err != nil {
		log.Fatal("failed to start server, ", err)
	}
}
