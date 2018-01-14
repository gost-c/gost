package main

import (
	"github.com/gost-c/gost/debug"
	"github.com/gost-c/gost/internal/controllers"
	"github.com/gost-c/gost/internal/utils"
	"github.com/gost-c/gost/logger"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.Post("/register", controllers.RegisterHandler)
	app.Post("/login", controllers.LoginHandler)

	// if debug mode, load debug routers
	if utils.GetEnvOrDefault("ENV", "prod") == "debug" {
		logger.Logger.Debug("debug mode, load debug routers")
		debug.LoadDebugRouters(app)
	}

	app.Run(iris.Addr("localhost:9393"))
}
