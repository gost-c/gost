package server

import (
	"github.com/codehack/scrypto"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"gopkg.in/appleboy/gin-jwt.v2"
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
			c.JSON(http.StatusOK, CreateRes("400", "Username's len should > 6 and < 20 and password's len should > 6!"))
			return
		}
		hash, _ := scrypto.Hash(json.Password)
		if err := db.Save(&User{Username: json.Username, Password: hash}).Error; err != nil {
			c.JSON(http.StatusOK, CreateRes("400", err.Error()))
			return
		}
		c.JSON(http.StatusOK, CreateRes("200", "register success, please login"))
		return
	}
	c.JSON(http.StatusOK, CreateRes("400", "register error, username and password is required!"))
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hello world"})
}

func createHandler(c *gin.Context) {
	jwtClaims := jwt.ExtractClaims(c)
	username, _ := jwtClaims["id"].(string)
	user := FindUserByName(username)
	var gist Gist
	if c.BindJSON(&gist) != nil {
		c.JSON(http.StatusOK, CreateRes("400", "post data error!"))
		return
	}
	gist.UserID = user.Model.ID
	gist.Hash = uuid.NewV4().String()
	if err := db.Save(&gist).Error; err != nil {
		c.JSON(http.StatusOK, CreateRes("400", err.Error()))
		return
	}
	c.JSON(http.StatusOK, CreateRes("200", "create success"))
}

func showGistHandler(c *gin.Context) {
	id := c.Param("hash")
	var files []*File
	gist := FindGistByHash(id)
	db.Model(&gist).Related(&files)
	gist.Files = files
	c.JSON(http.StatusOK, gist)
}

func deleteHandler(c *gin.Context) {
	jwtClaims := jwt.ExtractClaims(c)
	username, _ := jwtClaims["id"].(string)
	userID := GetUserIDByName(username)
	hash := c.Param("hash")
	var gist Gist
	db.First(&gist, "hash=?", hash)
	if gist.Hash != hash {
		c.JSON(http.StatusOK, CreateRes("400", "gist not exists!"))
		return
	}
	if userID != gist.UserID {
		c.JSON(http.StatusOK, CreateRes("400", "permission denied!"))
		return
	}
	if err := db.Delete(&gist).Error; err != nil {
		c.JSON(http.StatusOK, CreateRes("400", "delete failed!"))
		return
	}
	c.JSON(http.StatusOK, CreateRes("200", "delete success!"))
}

func GinEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	authMiddleware := GetAuthMiddleware()
	r.POST("/register", registerHandler)
	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/gist/:hash", showGistHandler)
	r.GET("/mock", MockHandler)
	r.POST("/test", testHandler)
	api := r.Group("/api")
	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/refresh_token", authMiddleware.RefreshHandler)
		api.GET("/hello", hello)
		api.POST("/create", createHandler)
		api.GET("/delete/:hash", deleteHandler)
	}
	return r
}

func testHandler(c *gin.Context) {
	var json Gist
	if c.BindJSON(&json) != nil {
		c.JSON(400, gin.H{"code": 400, "msg": "post data error!"})
		return
	}
	c.JSON(200, CreateRes("200", "hello"))
}
