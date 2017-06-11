package server

import (
	"github.com/codehack/scrypto"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"gopkg.in/appleboy/gin-jwt.v2"
	"gopkg.in/gin-contrib/cors.v1"
	"net/http"
	"os"
)

func registerHandler(c *gin.Context) {
	type Login struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	var json Login
	if c.BindJSON(&json) == nil {
		if len(json.Username) < 6 || len(json.Username) > 20 || len(json.Password) < 6 {
			c.JSON(http.StatusOK, createRes("400", "Username's len should > 6 and < 20 and password's len should > 6!"))
			return
		}
		hash, _ := scrypto.Hash(json.Password)
		if err := db.Save(&User{Username: json.Username, Password: hash}).Error; err != nil {
			c.JSON(http.StatusOK, createRes("400", err.Error()))
			return
		}
		c.JSON(http.StatusOK, createRes("200", "register success, please login"))
		return
	}
	c.JSON(http.StatusOK, createRes("400", "register error, username and password is required!"))
}

func createHandler(c *gin.Context) {
	jwtClaims := jwt.ExtractClaims(c)
	username, _ := jwtClaims["id"].(string)
	user := findUserByName(username)
	var gist Gist
	if c.BindJSON(&gist) != nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusBadRequest, createRes("400", "post data error!"))
		return
	}
	gist.UserID = user.Model.ID
	gist.User = user
	gist.Hash = uuid.NewV4().String()
	if err := db.Save(&gist).Error; err != nil {
		c.JSON(http.StatusOK, createRes("400", err.Error()))
		return
	}
	c.JSON(http.StatusOK, createRes("200", gist.Hash))
}

func showGistHandler(c *gin.Context) {
	hash := c.Param("hash")
	var files []File
	var user User
	gist := findGistByHash(hash)
	db.Model(&gist).Related(&files)
	if gist.Hash != hash {
		c.JSON(http.StatusOK, createRes("400", "gist not exists!"))
		return
	}
	db.First(&user, gist.UserID)
	gist.Files = files
	gist.User = user
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  gist,
	})
}

func deleteHandler(c *gin.Context) {
	jwtClaims := jwt.ExtractClaims(c)
	username, _ := jwtClaims["id"].(string)
	userID := getUserIDByName(username)
	hash := c.Param("hash")
	var gist Gist
	db.First(&gist, "hash=?", hash)
	if gist.Hash != hash {
		c.JSON(http.StatusOK, createRes("400", "gist not exists!"))
		return
	}
	if userID != gist.UserID {
		c.JSON(http.StatusOK, createRes("400", "permission denied!"))
		return
	}
	if err := db.Delete(&gist).Error; err != nil {
		c.JSON(http.StatusOK, createRes("400", "delete failed!"))
		return
	}
	c.JSON(http.StatusOK, createRes("200", "delete success!"))
}

// GinEngine provide gin Engine instance
func GinEngine() *gin.Engine {
	r := gin.New()
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	authMiddleware := getAuthMiddleware()
	r.POST("/register", registerHandler)
	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/gist/:hash", showGistHandler)
	if os.Getenv("GIN_MODE") == "debug" {
		r.GET("/mock", mockHandler)
	}
	api := r.Group("/api")
	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/refresh_token", authMiddleware.RefreshHandler)
		api.POST("/create", createHandler)
		api.GET("/delete/:hash", deleteHandler)
	}
	return r
}
