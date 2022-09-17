package routes

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"project/config"
	"project/controllers"
)

func New() *echo.Echo {
	contract := config.Contract{DB: config.DB}
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}

	v1 := e.Group("/v1")
	jwtMiddleware := middleware.JWT([]byte(os.Getenv("SECRET_JWT")))

	/**
	 * Routes of books V1
	 */

	bookController := controllers.NewBookController(contract)
	v1.GET("/books", bookController.GetBooks)
	v1.GET("/books/:id", bookController.GetBookByID)
	v1.POST("/books", bookController.Create, jwtMiddleware)
	v1.PUT("/books/:id", bookController.Update, jwtMiddleware)
	v1.DELETE("/books/:id", bookController.Delete, jwtMiddleware)

	/**
	 * Routes of users v1
	 */

	userController := controllers.NewUserController(contract)
	v1.GET("/users", userController.GetUsers, jwtMiddleware)
	v1.GET("/users/:id", userController.GetUserByID, jwtMiddleware)
	v1.POST("/users", userController.Create)
	v1.PUT("/users/:id", userController.Update, jwtMiddleware)
	v1.DELETE("/users/:id", userController.Delete, jwtMiddleware)

	v1.POST("/login", userController.Login)

	return e
}
