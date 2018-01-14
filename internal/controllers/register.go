package controllers

import (
	"fmt"
	"github.com/gost-c/gost/internal/models"
	"github.com/gost-c/gost/internal/utils"
	"github.com/kataras/iris"
)

var RegisterSuccess = "Register success, your username is %s. You can login later."

// RegisterHandler is handler for register
func RegisterHandler(ctx iris.Context) {
	user := models.User{}
	err := ctx.ReadJSON(&user)
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	ok, err := user.Validate()
	if !ok {
		utils.ResponseErr(ctx, err)
		return
	}
	u := user.New()
	err = u.Create()
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	utils.ResponseData(ctx, fmt.Sprintf(RegisterSuccess, user.Username))
}
