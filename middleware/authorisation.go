package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte("secret")},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
			"data":    nil,
		})
	}
}
