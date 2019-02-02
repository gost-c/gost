package models

import "github.com/zcong1993/libgo/gin/ginhelper"

// Gost is gost model
type Gost struct {
	ginhelper.Model
	Private     bool    `json:"private"`
	Description string  `json:"description" gorm:"type:varchar(255)"`
	Version     int     `json:"version" gorm:"default:1"`
	Files       []*File `json:"files" binding:"gt=0,dive"`
	User        *User   `json:"-"`
	UserID      uint    `json:"-"`
}

// File is file model
type File struct {
	ginhelper.Model
	Filename string `json:"filename" gorm:"type:varchar(100);unique_index:file_gost_index" binding:"required"`
	Content  string `json:"content" gorm:"type:text" binding:"required"`
	GostID   uint   `json:"-" gorm:"unique_index:file_gost_index"`
}
