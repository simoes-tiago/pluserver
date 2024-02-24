package main

import (
	"log"
	"net/http"
	"pluserver/db"
	"pluserver/handlers"
)

func main() {

	router := handlers.InitRouter()
	_, err := db.InitDB()
	if err != nil {
		log.Fatalln("not able to start database")
	}
	// Start the server
	log.Println("this is my server")
	log.Fatalln(http.ListenAndServe(":8080", router))
}
