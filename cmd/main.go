package main

import (
	"go-fiber-app/config"
	"go-fiber-app/internal/router"
	"go-fiber-app/pkg/db"
	"go-fiber-app/pkg/logger"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize logger
	logger.InitLogger()

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize database
	if err := db.InitDB(cfg); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Go Fiber App",
	})

	// Setup routes
	router.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(cfg.ServerPort))
}
