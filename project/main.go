package main

import (
	"day2-crud/config"
	"day2-crud/middlewares"
	"day2-crud/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.NewLogMiddleware(e).LogMiddleWare()
	e.Logger.Fatal(e.Start(":8080"))
}
