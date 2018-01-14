package debug

import (
	"fmt"
	"github.com/gost-c/gost/internal/jwt"
	"github.com/gost-c/gost/internal/middlewares"
	"github.com/gost-c/gost/internal/models/user"
	"github.com/gost-c/gost/internal/utils"
	"github.com/gost-c/gost/logger"
	"github.com/kataras/iris"
)

var log = logger.Logger

func LoadDebugRouters(app *iris.Application) {
	app.Get("/debug/delete/{username:string}", func(ctx iris.Context) {
		username := ctx.Params().Get("username")
		u := user.User{Username: username}
		err := u.Remove()
		if err != nil {
			utils.ResponseErr(ctx, err)
			return
		}
		utils.ResponseData(ctx, fmt.Sprintf("remove user success %s", username))
	})

	app.Get("/debug/addtoken/{token:string}", func(ctx iris.Context) {
		token := ctx.Params().Get("token")
		u := user.User{Username: "zc1993"}
		err := u.AddToken(token)
		if err != nil {
			utils.ResponseErr(ctx, err)
			return
		}
		err = u.GetUserByName()
		log.Debugf("add token, %#v %#v", u, err)
		utils.ResponseData(ctx, fmt.Sprintf("add token success %s", token))
	})

	app.Get("/debug/removetoken/{token:string}", func(ctx iris.Context) {
		token := ctx.Params().Get("token")
		u := user.User{Username: "zc1993"}
		err := u.RemoveToken(token)
		if err != nil {
			utils.ResponseErr(ctx, err)
			return
		}
		err = u.GetUserByName()
		log.Debugf("remove token, %#v %#v", u, err)
		utils.ResponseData(ctx, fmt.Sprintf("remove token success %s", token))
	})

	app.Get("/debug/test", jwt.JwtMiddleware.Serve, middlewares.AuthMiddleware, func(ctx iris.Context) {
		logger.Logger.Debug(ctx.Values().Get(middlewares.ContextKey))
		user, ok := ctx.Values().Get(middlewares.ContextKey).(*user.User)
		if !ok {
			ctx.Writef("not a user")
			return
		}
		ctx.Writef("%#v", user)
	})
}
