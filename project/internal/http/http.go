package http

import (
	"github.com/faizalnurrozi/alterra-agmc/internal/app/auth"
	"github.com/faizalnurrozi/alterra-agmc/internal/app/book"
	"github.com/faizalnurrozi/alterra-agmc/internal/app/user"
	"github.com/faizalnurrozi/alterra-agmc/internal/factory"
	"github.com/faizalnurrozi/alterra-agmc/pkg/util"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})
	v1 := e.Group("/api/v1")
	user.NewHandler(f).Route(v1.Group("/users"))
	auth.NewHandler(f).Route(v1.Group("/auth"))
	book.NewHandler(f).Route(v1.Group("/books"))
}
