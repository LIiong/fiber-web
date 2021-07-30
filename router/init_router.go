package router

import (
	"fiber-web/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitRouter() *fiber.App {
	// Create a new fiber instance with custom config
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Default 500 statuscode
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				// Override status code if fiber.Error type
				code = e.Code
			}
			// Set Content-Type: text/plain; charset=utf-8
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			// Return statuscode with error message
			return c.Status(code).SendString(err.Error())
		},
	})

	middleware.InitMiddleware(app)
	// 注册系统路由
	InitSysRouter(app)
	return app
}
