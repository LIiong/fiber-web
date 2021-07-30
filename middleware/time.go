package middleware

import (
	"fiber-web/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"time"
)

// Timer will measure how long it takes before a response is returned
func Timer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// start timer
		start := time.Now()
		// next routes
		err := c.Next()
		// stop timer
		stop := time.Now()
		logger.Info("request use time ", stop.Sub(start).String())
		// Do something with response
		//c.Append("Server-Timing", fmt.Sprintf("app;dur=%v", stop.Sub(start).String()))
		// return stack error if exist
		return err
	}
}
