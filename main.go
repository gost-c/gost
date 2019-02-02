package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/zcong1993/gost/controllers"
	_ "github.com/zcong1993/gost/db"
	"github.com/zcong1993/gost/middlewares"
	"github.com/zcong1993/libgo/gin/ginerr"
	"github.com/zcong1993/libgo/gin/ginhelper"
	"github.com/zcong1993/libgo/validator"
	"os"
)

func corsConfig() cors.Config {
	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	c.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	c.AllowMethods = []string{"GET", "POST", "PUT", "HEAD", "DELETE", "PATCH", "OPTIONS"}
	c.ExposeHeaders = []string{ginhelper.HEADER_TOTAL_COUNT}
	return c
}

func createGinEngine() *gin.Engine {
	r := gin.Default()

	binding.Validator = new(validator.DefaultValidator)

	r.Use(cors.New(corsConfig()))

	if os.Getenv("GIN_MODE") != "release" {
		// dev and test
	} else {
		// production
		r.Use(secure.New(secure.DefaultConfig()))
	}

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/register", ginerr.CreateGinController(controllers.Register()))
		v1.POST("/login", ginerr.CreateGinController(controllers.Login()))
		v1.DELETE("/revoke/:token", ginerr.CreateGinController(controllers.RevokeToken()))
		v1.GET("/gosts", ginerr.CreateGinController(controllers.GetAllGosts()))
		v1.GET("/users/:id/gosts", ginerr.CreateGinController(controllers.UserGosts()))
	}

	auth := v1.Group("", middlewares.AuthMiddleware())
	{
		auth.GET("/me", ginerr.CreateGinController(controllers.Me()))
		auth.GET("/me/gosts", ginerr.CreateGinController(controllers.MyGosts()))
		auth.POST("/gosts", ginerr.CreateGinController(controllers.CreateGost()))
		auth.DELETE("/gosts/:id", ginerr.CreateGinController(controllers.DeleteGost()))
		auth.GET("/gosts/:id", ginerr.CreateGinController(controllers.RetrieveGost()))
	}

	return r
}

func main() {
	r := createGinEngine()
	r.Run(":8080")
}
