package router

import (
	"go-fiber/handler"
	"go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	user := api.Group("user")
	user.Get("/", middleware.Protected(), handler.GetUsers)
}
