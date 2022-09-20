package main

import (
	"github.com/faizalnurrozi/alterra-agmc/config"
	"github.com/faizalnurrozi/alterra-agmc/middlewares"
	"github.com/faizalnurrozi/alterra-agmc/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.NewLogMiddleware(e).LogMiddleWare()
	e.Logger.Fatal(e.Start(":8080"))
}
