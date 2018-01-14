package main

import (
	"github.com/gost-c/gost/internal/controllers"
	"github.com/gost-c/gost/internal/jwt"
	"github.com/gost-c/gost/internal/middlewares"
	"github.com/gost-c/gost/internal/models"
	"github.com/gost-c/gost/logger"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.Post("/register", controllers.RegisterHandler)
	app.Post("/login", controllers.LoginHandler)

	app.Get("/test", jwt.JwtMiddleware.Serve, middlewares.AuthMiddleware, func(ctx iris.Context) {
		logger.Logger.Debug(ctx.Values().Get(middlewares.ContextKey))
		user, ok := ctx.Values().Get(middlewares.ContextKey).(*models.User)
		if !ok {
			ctx.Writef("not a user")
			return
		}
		ctx.Writef("%#v", user)
	})
	app.Run(iris.Addr("localhost:9393"))
}
