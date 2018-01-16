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
	// ErrBadUser is error message for bad user
	ErrBadUser = errors.New("Bad user. ")
	// ErrBadGost is error message for bad gost
	ErrBadGost = errors.New("Bad gost, maybe you push too many files or file is too big. ")
	// ErrGostNotFound is error message for gost not found
	ErrGostNotFound = errors.New("Gost not exists. ")
	// ErrNotYourOwn is error message for user try to delete others gost
	ErrNotYourOwn = errors.New("This gost is not owned to you. ")
	log           = logger.Logger
)

// PublishHandler is handler for create new gost
func PublishHandler(ctx iris.Context) {
	println(ctx.Request().ContentLength)
	ctx.SetMaxRequestBodySize(32 << 20)
	u, ok := ctx.Values().Get(middlewares.ContextKey).(*user.User)
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
	g.WithUser(*u)
	err = g.Create()
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	utils.ResponseData(ctx, g.ID)
}

// GetController is handler for get gost by id
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

// DeleteController is handler for delete a gost
func DeleteController(ctx iris.Context) {
	u, ok := ctx.Values().Get(middlewares.ContextKey).(*user.User)
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
	if u.Username != g.User.Username {
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

// UserOwnGostsController is handler show user own gosts (private)
func UserOwnGostsController(ctx iris.Context) {
	u, ok := ctx.Values().Get(middlewares.ContextKey).(*user.User)
	if !ok {
		utils.ResponseErr(ctx, ErrBadUser)
		return
	}
	var g gost.Gost
	gosts, err := g.GetGostsByUsername(u.Username)
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	utils.ResponseData(ctx, gosts)
}

// UserGostsController is handler show a user's own gosts, (public)
func UserGostsController(ctx iris.Context) {
	username := ctx.Params().Get("username")
	u := user.User{Username: username}
	err := u.GetUserByName()
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	var g gost.Gost
	gosts, err := g.GetGostsByUsername(username)
	if err != nil {
		utils.ResponseErr(ctx, err)
		return
	}
	utils.ResponseData(ctx, gosts)
}

// RawGostHandler is handler for raw file
func RawGostHandler(ctx iris.Context) {
	id := ctx.Params().Get("id")
	file := ctx.Params().Get("file")
	var g gost.Gost
	err := g.GetGostById(id)
	if err != nil {
		ctx.Text("Not found")
		return
	}
	f := g.FindFile(file)
	if f == nil {
		ctx.Text("Not found")
		return
	}
	ctx.Text(f.Content)
}
