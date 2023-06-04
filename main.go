package main

import (
	"log"

	"daeuniverse/functions/pkg/web/server"
)

func main() {
	log.Printf("Server listening on port: %d", 3000)
	router := server.NewRouter()
	err := router.Serve(3000)
	if err != nil {
		log.Fatal(err)
	}
}
