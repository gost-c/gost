package server

import (
	"github.com/codehack/scrypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerHandler(c *gin.Context) {
	type Login struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	var json Login
	if c.BindJSON(&json) == nil {
		if len(json.Username) < 6 || len(json.Username) > 20 || len(json.Password) < 6 {
			c.JSON(http.StatusOK, gin.H{
				"code": "400",
				"msg":  "Username's len should > 6 and < 20 and password's len should > 6!",
			})
			return
		}
		var user User
		db.First(&user, "name=?", json.Username)
		if user.Username == json.Username {
			c.JSON(http.StatusOK, gin.H{
				"code": "400",
				"msg":  "Username exists!",
			})
			return
		}
		hash, _ := scrypto.Hash(json.Password)
		db.Create(&User{Username: json.Username, Password: hash})
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "register success, please login",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "500",
		"msg":  "register error, username and password is required!",
	})
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hello world"})
}

func GinEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	authMiddleware := GetAuthMiddleware()
	r.POST("/register", registerHandler)
	r.POST("/login", authMiddleware.LoginHandler)
	api := r.Group("/api")
	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/refresh_token", authMiddleware.RefreshHandler)
		api.GET("/hello", hello)
	}
	return r
}
