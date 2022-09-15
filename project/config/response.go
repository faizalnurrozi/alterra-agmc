package config

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HTTPResponse struct {
	Status string      `json:"message"`
	Data   interface{} `json:"data"`
}

func (res HTTPResponse) ResponseOk(ctx echo.Context, data interface{}) error {
	res.Status = "Success"
	res.Data = data
	return ctx.JSON(http.StatusOK, res)
}
