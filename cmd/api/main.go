package main

import (
	"github.com/kenesparta/golang-solid/cmd/api/handlers"
	"log"
	"net/http"
)

func main() {
	r := handlers.GetRouter()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
		return
	}
}
