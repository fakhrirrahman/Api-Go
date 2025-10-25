package main

import (
	"goApi/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.SetupRoutes()

	log.Println("Server berjalan pada http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
