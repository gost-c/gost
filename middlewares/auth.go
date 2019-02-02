package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/zcong1993/gost/models"
	"github.com/zcong1993/gost/services"
	"github.com/zcong1993/libgo/gin/ginerr"
	"net/http"
	"strings"
)

const USER_CTX_KEY = "USER_CTX_KEY"

var errResp = ginerr.NewDefaultError(http.StatusUnauthorized, "TOKEN_ERROR", "TOKEN_ERROR", nil)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		a := c.GetHeader("Authorization")
		token := strings.Replace(a, "Bearer ", "", 1)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
			return
		}
		u, err := services.GetUserByToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
			return
		}

		c.Set(USER_CTX_KEY, u)
		c.Next()
	}
}

func MustGetUser(ctx *gin.Context) *models.User {
	user, ok := ctx.Get(USER_CTX_KEY)
	if !ok {
		panic("can't get user from ctx")
	}
	u, ok := user.(*models.User)
	if !ok {
		panic("can't get correct user model from ctx")
	}
	return u
}
