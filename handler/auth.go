package handler

import (
	"go-fiber/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Login get user and password
func Login(c *fiber.Ctx) error {

	admin := config.Config("LOCAL_USERNAME")
	password := config.Config("LOCAL_PASSWORD")

	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	identity := input.Identity
	pass := input.Password
	if identity != admin || pass != password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
