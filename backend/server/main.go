package main

import (
	"SuperHack2023/server/router"
	"log"
	"net/http"
)

func main() {
	rtr := router.New()

	log.Println("Server listening on http://localhost:3000/")

	err := http.ListenAndServe("0.0.0.0:3000", rtr)

	if err != nil {
		log.Fatalln("There was an error with the http server: %v", err)
	}
}
