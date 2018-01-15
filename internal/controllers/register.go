package controllers

import (
	"fmt"
	"github.com/gost-c/gost/internal/models/user"
	"github.com/gost-c/gost/internal/types"
	"github.com/gost-c/gost/internal/utils"
	"github.com/kataras/iris"
)

// RegisterSuccess is message for register success
var RegisterSuccess = "Register success, your username is %s. You can login later."

// RegisterHandler is handler for register
func RegisterHandler(ctx iris.Context) {
	userform := types.UserForm{}
	err := ctx.ReadJSON(&userform)
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	ok, err := userform.Validate()
	if !ok {
		utils.ResponseErr(ctx, err)
		return
	}
	u := user.NewUser(userform.Username, userform.Password)
	err = u.Create()
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	utils.ResponseData(ctx, fmt.Sprintf(RegisterSuccess, u.Username))
}
