package main

import (
	"project/config"
	"project/middlewares"
	"project/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.NewLogMiddleware(e).LogMiddleWare()
	e.Logger.Fatal(e.Start(":8080"))
}
