package main

import (
	"go-fiber/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

/**
Things todo
- authentication
- middleware to standardise all responses and statuses
*/

func main() {

	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()

	router.SetupRouter(app)
	log.Fatal(app.Listen(":3000"))
}

// Hello handle api status
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
