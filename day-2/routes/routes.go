package routes

import (
	"day2-crud/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")
	v1.GET("/books", controllers.GetBooksController)
	v1.GET("/books/:id", controllers.GetBookByIDController)
	v1.POST("/books", controllers.StoreBook)
	v1.PUT("/books/:id", controllers.UpdateBook)
	v1.DELETE("/books/:id", controllers.DeleteBook)
	return e
}
