package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8088",
		Handler: router,
	}

	BuilderRandomNum(router, "/random")

	log.Println("Server starts...")
	log.Fatal(server.ListenAndServe())
}
