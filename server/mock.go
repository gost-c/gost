package server

import (
	"github.com/codehack/scrypto"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func MockHandler(c *gin.Context) {
	passwordHash, _ := scrypto.Hash("gostmock")
	user := &User{Username: "gostmockuser", Password: string(passwordHash)}
	db.Where(User{Username: "gostmockuser"}).FirstOrCreate(&user)
	gist := CreateDefaultGist()
	gist.UserID = user.Model.ID
	file1, _ := ioutil.ReadFile("main.go")
	file2, _ := ioutil.ReadFile(".travis.yml")
	gist.Files = []*File{{Filename: "main.go", Content: string(file1)}, {Filename: ".travis.yml", Content: string(file2)}}
	db.Where(Gist{UserID: user.Model.ID}).FirstOrCreate(&gist)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "mock success",
	})
}
