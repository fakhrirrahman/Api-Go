package main

import (
	"goApi/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	log.Println("Server berjalan pada http://localhost:8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
