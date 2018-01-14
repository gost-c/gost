package controllers

import (
	"github.com/gost-c/gost/internal/models"
	"github.com/gost-c/gost/internal/utils"
	"github.com/kataras/iris"
)

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
	ctx.Writef("%#v", user)
}
