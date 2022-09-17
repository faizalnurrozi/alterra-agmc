package config

import (
	"github.com/labstack/echo/v4"
)

type HTTPResponse struct {
	Status     string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

func (res HTTPResponse) ResponseOk(ctx echo.Context, statusCode int, data interface{}) error {
	res.Status = "Success"
	res.Data = data
	return ctx.JSON(statusCode, res)
}
