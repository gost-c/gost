package server

import (
	"github.com/codehack/scrypto"
	"github.com/gin-gonic/gin"
	"gopkg.in/appleboy/gin-jwt.v2"
	"os"
	"time"
)

func GetAuthMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      "server",
		Key:        []byte(os.Getenv("JWT_SECRET")),
		Timeout:    time.Hour * 24 * 365,
		MaxRefresh: time.Hour * 24 * 365,
		Authenticator: func(userID string, password string, c *gin.Context) (string, bool) {
			var user User
			db.First(&user, "username=?", userID)
			if user.Username == userID && scrypto.Compare(password, user.Password) {
				return userID, true
			}
			return userID, false
		},
		Authorizator: func(userID string, c *gin.Context) bool {
			var user User
			db.First(&user, "username=?", userID)
			if user.Username == userID {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
}
