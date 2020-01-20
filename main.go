package main

import (
	"log"
	"net/http"

	"./routers"
)

func main() {
	router := routers.InitRoutes()
	port := ":3000"
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Println("Listening on port ", port)

	log.Fatal(server.ListenAndServe())
}
