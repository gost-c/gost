package server

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func FindUserByName(name string) *User {
	var user User
	db.First(&user, "username=?", name)
	return &user
}

func GetUserIDByName(name string) uint {
	user := FindUserByName(name)
	return user.Model.ID
}

func CreateDefaultGist() *Gist {
	return &Gist{
		Public:      false,
		Version:     1,
		Hash:        uuid.NewV4().String(),
		Description: "published by zcong1993/gost-cli",
	}
}

func AppendFileToGist(gist *Gist, file *File) {
	db.Model(gist).Association("Files").Append(file)
}

func FindGistByHash(hash string) *Gist {
	var gist Gist
	db.First(&gist, "hash=?", hash)
	return &gist
}

func CreateRes(code, msg string) *gin.H {
	return &gin.H{"code": code, "msg": msg}
}
