package controllers

import (
	"errors"
	"github.com/gost-c/gost/internal/jwt"
	"github.com/gost-c/gost/internal/models/user"
	"github.com/gost-c/gost/internal/utils"
	"github.com/kataras/iris"
)

var (
	// ErrPasswordError is password error message
	ErrPasswordError = errors.New("Password mismatch username.")
)

// LoginHandler is http handler for login router
func LoginHandler(ctx iris.Context) {
	user := user.User{}
	err := ctx.ReadJSON(&user)
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	password := user.Password
	err = user.GetUserByName()
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	ok := utils.CheckPassword(password, user.Password)
	if !ok {
		utils.ResponseErr(ctx, ErrPasswordError)
		return
	}
	user.Password = password
	token, err := jwt.JwtEncode(&user)
	if err != nil {
		utils.ResponseErr(ctx, ErrPasswordError)
		return
	}
	utils.ResponseData(ctx, token)
}
