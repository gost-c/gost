package server

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	d.AutoMigrate(&User{}, &Gist{})
	db = d
}

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100;unique"`
	Gists    []Gist
}

type Gist struct {
	gorm.Model
	UserID      uint `gorm:"index"`
	Pubilc      bool
	Description string
	Version     uint
	Hash        string `gorm:"type:char(100);index"`
	Files       []File
}

type File struct {
	Filename string
	Content  string
}
