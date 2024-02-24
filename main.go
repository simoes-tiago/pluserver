package main

import (
	"log"
	"net/http"
	"pluserver/db"
	"pluserver/handlers"
	"pluserver/service"
)

func main() {

	database, err := db.InitDB()
	if err != nil {
		log.Fatalln("not able to start database")
	}
	svc := service.NewService(
		database,
	)
	router := handlers.InitRouter(*svc)

	// Start the server
	log.Println("this is my server")
	log.Fatalln(http.ListenAndServe(":8080", router))
}
