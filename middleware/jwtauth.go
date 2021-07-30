package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"fiber-web/api"
)

// Protected protect routes
func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return api.Response(c, err, nil)
	} else {
		c.Status(fiber.StatusUnauthorized)
		return api.Response(c, err, nil)
	}
}