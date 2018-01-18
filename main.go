package main

import (
	"fmt"
	"github.com/gost-c/gost/debug"
	"github.com/gost-c/gost/internal/controllers"
	"github.com/gost-c/gost/internal/jwt"
	"github.com/gost-c/gost/internal/middlewares"
	"github.com/gost-c/gost/internal/utils"
	"github.com/gost-c/gost/logger"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	app.Use(crs)

	// public router
	app.PartyFunc("/api", func(r iris.Party) {
		r.Post("/register", controllers.RegisterHandler)
		r.Post("/login", controllers.LoginHandler)
		r.Get("/gost/{id:string}", controllers.GetController)
		r.Get("/gosts/{username:string}", controllers.UserGostsController)
		r.Get("/raw/gost/{id:string}/{file:string}", controllers.RawGostHandler)
	})

	// private router
	app.PartyFunc("/api", func(r iris.Party) {
		r.Use(jwt.JwtMiddleware.Serve, middlewares.AuthMiddleware)
		r.Post("/gost", controllers.PublishHandler)
		r.Delete("/gost/{id:string}", controllers.DeleteController)
		r.Get("/user/gosts", controllers.UserOwnGostsController)
	})

	// if debug mode, load debug routers
	if utils.GetEnvOrDefault("ENV", "prod") == "debug" {
		logger.Logger.Debug("debug mode, load debug routers")
		debug.LoadDebugRouters(app)
	}

	host := utils.GetEnvOrDefault("HOST", "")
	port := utils.GetEnvOrDefault("PORT", "9393")

	app.Run(
		iris.Addr(fmt.Sprintf("%s:%s", host, port)),
		// disables updates:
		iris.WithoutVersionChecker,
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}
