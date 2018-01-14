package middlewares

import (
	"errors"
	"github.com/gost-c/gost/internal/jwt"
	"github.com/gost-c/gost/internal/utils"
	"github.com/gost-c/gost/logger"
	"github.com/kataras/iris"
)

var (
	// ContextKey is AuthMiddleware user context key
	ContextKey = "gostuser"
	// ErrPasswordMismatch is password mismatch error
	ErrPasswordMismatch = errors.New("Password mismatch user.")
)

// AuthMiddleware auth user, if failed response 401, success set user to context ContextKey
func AuthMiddleware(ctx iris.Context) {
	token := jwt.JwtMiddleware.Get(ctx)
	user, err := jwt.JwtDecode(token)
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		utils.ResponseErr(ctx, err)
		ctx.StopExecution()
		return
	}
	password := user.Password
	err = user.GetUserByName()
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		utils.ResponseErr(ctx, err)
		ctx.StopExecution()
		return
	}
	ok := utils.CheckPassword(password, user.Password)
	if !ok {
		ctx.StatusCode(iris.StatusUnauthorized)
		utils.ResponseErr(ctx, ErrPasswordMismatch)
		ctx.StopExecution()
		return
	}
	logger.Logger.Debugf("%#v", user)
	ctx.Values().Set(ContextKey, user)
	ctx.Next()
}
