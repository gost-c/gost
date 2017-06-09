package server

import "github.com/satori/go.uuid"

func FindUserByName(name string) *User {
	var user User
	db.First(&user, "username=?", name)
	return &user
}

func CreateDefaultGist() *Gist {
	return &Gist{
		Pubilc:false,
		Version: 1,
		Hash: uuid.NewV4().String(),
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
