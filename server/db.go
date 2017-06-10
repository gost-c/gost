package server

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // import mysql lib
	"log"
	"os"
)

var db *gorm.DB

func init() {
	println("open db...")
	link := os.Getenv("MYSQL_DB_URL")
	d, err := gorm.Open("mysql", link)
	if err != nil {
		log.Fatal(err)
	}
	d.AutoMigrate(&User{}, &Gist{}, &File{})
	db = d
}

// User struct is the model of User table
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100"`
	Gists    []Gist
}

// Gist struct is the model of Gist table
type Gist struct {
	gorm.Model
	User User
	UserID      uint    `gorm:"index"`
	Public      bool    `form:"public" json:"public"`
	Description string  `form:"description" json:"description" binding:"required"`
	Version     uint    `form:"version" json:"version" binding:"required"`
	Hash        string  `gorm:"type:char(100);index;unique" form:"hash" json:"hash"`
	Files       []File `form:"files" json:"files" binding:"required"`
}

// File struct is the model of File table
type File struct {
	gorm.Model
	GistID   uint   `gorm:"index"`
	Filename string `form:"filename" json:"filename" binding:"required"`
	Content  string `gorm:"type:text" form:"content" json:"content" binding:"required"`
}
