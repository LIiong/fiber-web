package api

import (
	"fiber-web/pkg/errno"
	"github.com/gofiber/fiber/v2"
)

type PageData struct {
	PageNo     int         `json:"pageNo"`
	Data       interface{} `json:"data"`
	TotalCount int         `json:"totalCount"`
}

// ResponseHTTP represents response body of this API
type ResponseHTTP struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *fiber.Ctx, err error, data interface{}) error {

	code, message := errno.DecodeErr(err)
	// always return http.StatusOK
	return c.JSON(ResponseHTTP{
		Success: true,
		Code:    code,
		Message: message,
		Data:    data,
	})
}
