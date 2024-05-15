package main

import (
	"flag"

	kvstoreserver "github.com/LavHinsu/gokvstore/server"
)

func main() {
	port := flag.Int("port", 8080, "specify port for the server to listen on")

	flag.Parse()
	kvstoreserver.Kvstoreserver(*port)
}
