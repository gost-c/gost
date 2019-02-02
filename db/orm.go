package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zcong1993/gost/models"
	"log"
	"os"
)

var ORM *gorm.DB

func InitDB(fns ...func(db *gorm.DB)) {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxIdleConns(50)

	if os.Getenv("ENV") == "debug" {
		db.LogMode(true)
	}

	ORM = db

	db2 := db.Set("gorm:table_options", "charset=utf8mb4")

	for _, fn := range fns {
		fn(db2)
	}
}

func init() {
	InitDB(func(db *gorm.DB) {
		// auto migrate here
		db.AutoMigrate(new(models.User), new(models.Token), new(models.Gost), new(models.File))
	})
}
