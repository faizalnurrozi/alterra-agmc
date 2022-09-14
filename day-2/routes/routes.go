package routes

import (
	"day2-crud/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")

	/**
	 * Routes of books V1
	 */

	bookController := controllers.BookController{}
	v1.GET("/books", bookController.GetBooks)
	v1.GET("/books/:id", bookController.GetBookByID)
	v1.POST("/books", bookController.Create)
	v1.PUT("/books/:id", bookController.Update)
	v1.DELETE("/books/:id", bookController.Delete)

	/**
	 * Routes of users v1
	 */

	userController := controllers.UserController{}
	v1.GET("/users", userController.GetUsers)
	v1.GET("/users/:id", userController.GetUserByID)
	v1.POST("/users", userController.Create)
	v1.PUT("/users/:id", userController.Update)
	v1.DELETE("/users/:id", userController.Delete)

	return e
}
