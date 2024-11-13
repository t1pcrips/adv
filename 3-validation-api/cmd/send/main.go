package main

import (
	"go-adv/3-validation-api/configs"
	"go-adv/3-validation-api/internal/verify"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	router := http.NewServeMux()
	verify.NewVerifyHandler(router, verify.VerifyServiceDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    conf.Url.Port,
		Handler: router,
	}

	log.Println("Server Sender starts...")
	log.Fatal(server.ListenAndServe())
}
