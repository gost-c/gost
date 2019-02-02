package models

import (
	"github.com/zcong1993/libgo/gin/ginhelper"
)

// User is user model
type User struct {
	ginhelper.Model
	Username string  `json:"username" gorm:"type:varchar(50);unique_index" binding:"required,min=6,max=12"`
	Password string  `json:"password" gorm:"type:varchar(100)" binding:"required,min=6,max=20"`
	Gosts    []*Gost `json:"-"`
}

// Token is token model
type Token struct {
	ginhelper.Model
	Token  string `json:"token" gorm:"type:varchar(100);unique_index"`
	UserID uint   `json:"-"`
	User   *User  `json:"-"`
}
