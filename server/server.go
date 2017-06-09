package server

import (
	"github.com/codehack/scrypto"
	"github.com/gin-gonic/gin"
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
			c.JSON(http.StatusOK, gin.H{
				"code": "400",
				"msg":  "Username's len should > 6 and < 20 and password's len should > 6!",
			})
			return
		}
		//user := FindUserByName(json.Username)
		//if user.Username == json.Username {
		//	c.JSON(http.StatusOK, gin.H{
		//		"code": "400",
		//		"msg":  "Username exists!",
		//	})
		//	return
		//}
		hash, _ := scrypto.Hash(json.Password)
		if err := db.Save(&User{Username: json.Username, Password: hash}).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": "500",
				"msg":  err.Error(),
			})
			return
		}
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

func createHandler(c *gin.Context) {
	jwtClaims := jwt.ExtractClaims(c)
	username, _ := jwtClaims["id"].(string)
	user := FindUserByName(username)
	gist := CreateDefaultGist()
	gist.UserID = user.Model.ID
	gist.Files = []*File{{Filename: "test.txt", Content: "xsxsxs"}, {Filename: "test1.txt", Content: "xsxsxsxsxs"}}
	if err := db.Save(&gist).Error; err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": user.Username, "id": user.Model.ID})
}

func showGistHandler(c *gin.Context) {
	id := c.Param("hash")
	var files []*File
	gist := FindGistByHash(id)
	db.Model(&gist).Related(&files)
	gist.Files = files
	//AppendFileToGist(&gist, &File{Filename:"xsxsxsxs", Content:"xsxsxsxsxs"})
	c.JSON(200, gist)
}

func GinEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	authMiddleware := GetAuthMiddleware()
	r.POST("/register", registerHandler)
	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/gist/:hash", showGistHandler)
	api := r.Group("/api")
	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/refresh_token", authMiddleware.RefreshHandler)
		api.GET("/hello", hello)
		api.POST("/create", createHandler)
	}
	return r
}
