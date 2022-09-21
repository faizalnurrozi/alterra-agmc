package book

import (
	"github.com/faizalnurrozi/alterra-agmc/internal/dto"
	"github.com/faizalnurrozi/alterra-agmc/internal/middleware"
	"github.com/faizalnurrozi/alterra-agmc/internal/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.Use(middleware.JWTMiddleware(dto.JWTClaims{}, util.JWT_SECRET))
	g.GET("", h.Get)
	g.GET("/:id", h.GetById)
	g.PUT("/:id", h.UpdateById)
	g.DELETE("/:id", h.DeleteById)
	g.POST("", h.Create)
}
