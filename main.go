package main

import (
	"github.com/gost-c/gost/internal/controllers"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.Post("/register", controllers.RegisterHandler)
	app.Run(iris.Addr("localhost:9393"))
}
