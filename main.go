package main

import (
	"log"

	"daeuniverse/functions/pkg/web/server"
)

func main() {
	log.Printf("Server listening on port: %d", 5888)
	router := server.NewRouter()
	err := router.Serve(5888)
	if err != nil {
		log.Fatal(err)
	}
}
