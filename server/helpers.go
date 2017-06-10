package server

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func findUserByName(name string) User {
	var user User
	db.First(&user, "username=?", name)
	return user
}

func getUserIDByName(name string) uint {
	user := findUserByName(name)
	return user.Model.ID
}

func createDefaultGist() *Gist {
	return &Gist{
		Public:      false,
		Version:     1,
		Hash:        uuid.NewV4().String(),
		Description: "published by zcong1993/gost-cli",
	}
}

func appendFileToGist(gist Gist, file File) {
	db.Model(gist).Association("Files").Append(file)
}

func findGistByHash(hash string) Gist {
	var gist Gist
	db.First(&gist, "hash=?", hash)
	return gist
}

func createRes(code, msg string) *gin.H {
	return &gin.H{"code": code, "msg": msg}
}
