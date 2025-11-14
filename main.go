package main

import (
	"goApi/config"
	"goApi/database"
	"goApi/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize database
	if err := database.Init(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Println("Server berjalan pada http://localhost:8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
