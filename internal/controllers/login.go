package controllers

import (
	"errors"
	"github.com/gost-c/gost/internal/jwt"
	"github.com/gost-c/gost/internal/models/user"
	"github.com/gost-c/gost/internal/types"
	"github.com/gost-c/gost/internal/utils"
	"github.com/kataras/iris"
)

var (
	// ErrPasswordError is password error message
	ErrPasswordError = errors.New("Password mismatch username.")
)

// LoginHandler is http handler for login router
func LoginHandler(ctx iris.Context) {
	ctx.SetMaxRequestBodySize(1024)
	userform := types.UserForm{}
	err := ctx.ReadJSON(&userform)
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	password := userform.Password
	u := user.NewUser(userform.Username, userform.Password)
	err = u.GetUserByName()
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	ok := utils.CheckPassword(password, u.Password)
	if !ok {
		utils.ResponseErr(ctx, ErrPasswordError)
		return
	}
	u.Password = password
	token, err := jwt.JwtEncode(u)
	if err != nil {
		utils.ResponseErr(ctx, ErrPasswordError)
		return
	}
	utils.ResponseData(ctx, token)
}
