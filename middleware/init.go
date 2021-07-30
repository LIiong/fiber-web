package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)


func InitMiddleware(app *fiber.App) {
	// 日志处理
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
	// 自定义错误处理
	app.Use(recover.New())
	// 跨域处理
	app.Use(cors.New())
	// Cache
	//app.Use(cache.New())
	app.Use(Timer())
	app.Use(requestid.New())
	app.Use(requestid.New())
}