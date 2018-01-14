package controllers

import (
	"errors"
	"github.com/gost-c/gost/internal/jwt"
	"github.com/gost-c/gost/internal/models"
	"github.com/gost-c/gost/internal/utils"
	"github.com/kataras/iris"
)

var (
	ErrPasswordError = errors.New("Password mismatch username.")
)

func LoginHandler(ctx iris.Context) {
	user := models.User{}
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
