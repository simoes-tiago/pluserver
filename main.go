package main

import (
	"log"
	"net/http"
	"pluserver/handlers"
)

func main() {
	log.Println("This will be a server")

	router := handlers.InitRouter()
	// Start the server
	log.Println("This is my server")
	log.Fatalln(http.ListenAndServe(":8080", router))
}
