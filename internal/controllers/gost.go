package controllers

import (
	"errors"
	"fmt"
	"github.com/gost-c/gost/internal/middlewares"
	"github.com/gost-c/gost/internal/models/gost"
	"github.com/gost-c/gost/internal/models/user"
	"github.com/gost-c/gost/internal/utils"
	"github.com/gost-c/gost/logger"
	"github.com/kataras/iris"
)

var (
	ErrBadUser      = errors.New("Bad user.")
	ErrBadGost      = errors.New("Bad gost.")
	ErrGostNotFound = errors.New("Gost not exists.")
	ErrNotYourOwn   = errors.New("This gost is not owned to you.")
	log             = logger.Logger
)

func PublishHandler(ctx iris.Context) {
	user, ok := ctx.Values().Get(middlewares.ContextKey).(*user.User)
	if !ok {
		utils.ResponseErr(ctx, ErrBadUser)
		return
	}
	var g gost.Gost
	err := ctx.ReadJSON(&g)
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	ok = g.Validate()
	if !ok {
		utils.ResponseErr(ctx, ErrBadGost)
		return
	}
	g.WithUser(*user)
	err = g.Create()
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	utils.ResponseData(ctx, g.ID)
}

func GetController(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var g gost.Gost
	err := g.GetGostById(id)
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	log.Debugf("find gost %#v", g)
	utils.ResponseData(ctx, g)
}

func DeleteController(ctx iris.Context) {
	user, ok := ctx.Values().Get(middlewares.ContextKey).(*user.User)
	if !ok {
		utils.ResponseErr(ctx, ErrBadUser)
		return
	}
	id := ctx.Params().Get("id")
	var g gost.Gost
	err := g.GetGostById(id)
	if err != nil {
		utils.ResponseErr(ctx, ErrGostNotFound)
		return
	}
	if user.Username != g.User.Username {
		utils.ResponseErr(ctx, ErrNotYourOwn)
		return
	}
	err = g.Remove(true)
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	utils.ResponseData(ctx, fmt.Sprintf("Gost remove success %s!", g.ID))
}

func UserOwnGostsController(ctx iris.Context) {
	user, ok := ctx.Values().Get(middlewares.ContextKey).(*user.User)
	if !ok {
		utils.ResponseErr(ctx, ErrBadUser)
		return
	}
	var g gost.Gost
	gosts, err := g.GetGostsByUsername(user.Username)
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	utils.ResponseData(ctx, gosts)
}

func UserGostsController(ctx iris.Context) {
	username := ctx.Params().Get("username")
	var g gost.Gost
	gosts, err := g.GetGostsByUsername(username)
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	utils.ResponseData(ctx, gosts)
}
