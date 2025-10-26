package router

import (
	"go-fiber-app/internal/middleware"
	"go-fiber-app/internal/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(middleware.LoggerMiddleware)

	// API routes
	api := app.Group("/api")

	// Setup feature routes
	user.SetupUserRoutes(api)
}
