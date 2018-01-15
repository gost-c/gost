package main

import (
	"github.com/gost-c/gost/debug"
	"github.com/gost-c/gost/internal/controllers"
	"github.com/gost-c/gost/internal/jwt"
	"github.com/gost-c/gost/internal/middlewares"
	"github.com/gost-c/gost/internal/utils"
	"github.com/gost-c/gost/logger"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	// public router
	app.Post("/api/register", controllers.RegisterHandler)
	app.Post("/api/login", controllers.LoginHandler)
	app.Get("/api/gost/{id:string}", controllers.GetController)
	app.Get("/api/gosts/{username:string}", controllers.UserGostsController)

	// private router
	app.Post("/api/gost", jwt.JwtMiddleware.Serve, middlewares.AuthMiddleware, controllers.PublishHandler)
	app.Delete("/api/gost/{id:string}", jwt.JwtMiddleware.Serve, middlewares.AuthMiddleware, controllers.DeleteController)
	app.Get("/api/user/gosts", jwt.JwtMiddleware.Serve, middlewares.AuthMiddleware, controllers.UserOwnGostsController)

	// if debug mode, load debug routers
	if utils.GetEnvOrDefault("ENV", "prod") == "debug" {
		logger.Logger.Debug("debug mode, load debug routers")
		debug.LoadDebugRouters(app)
	}

	app.Run(iris.Addr("localhost:9393"))
}
