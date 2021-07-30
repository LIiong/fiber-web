package router

import (
	"fiber-web/api/login"
	"fiber-web/api/task"
	_ "fiber-web/docs"
	"fiber-web/middleware"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func InitSysRouter(app *fiber.App) fiber.Router {
	router := app.Group("")
	router.Get("/docs/*", swagger.Handler)
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Hello, World 👋!",
		})
	})
	router.Post("/login", login.Login)
	api := router.Group("/api", middleware.Protected())
	taskApi := api.Group("/task")
	taskApi.Get("/list", task.FindAll)
	taskApi.Post("/save", task.Save)
	taskApi.Put("/:id", task.ChangeStatus)
	taskApi.Delete("/:id", task.Remove)
	return router
}
