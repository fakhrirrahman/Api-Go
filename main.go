// @title Booking API
// @version 1.0
// @description API for booking sports fields
// @host localhost:8080
// @BasePath /
package main

import (
	_ "goApi/docs"
	"goApi/config"
	"goApi/database"
	"goApi/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize database
	if err := database.Init(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	// Fiber app
	app := fiber.New()

	// Setup all routes
	routes.SetupRoutes(app)

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Println("Server berjalan pada http://localhost:8080")
	log.Println("Swagger tersedia di http://localhost:8080/swagger/index.html")

	// Start server
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
