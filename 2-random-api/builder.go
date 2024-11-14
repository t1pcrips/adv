package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

func BuilderRandomNum(router *http.ServeMux, path string) {
	router.HandleFunc(path, GetRandomNum)
}

func GetRandomNum(w http.ResponseWriter, r *http.Request) {
	num := fmt.Sprintf("%d", rand.Int()%6+1)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(num))

	log.Println("Get Random num: ", num)
}
