package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gost-c/gost/internal/models"
	"github.com/gost-c/gost/internal/utils"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

var (
	// ErrJwtDecode is error message for get jwt claims failed
	ErrJwtDecode = errors.New("JWT Claims Failed")
	// DefaultKey is default jwt key we used
	DefaultKey = "you never know zcong"
	// JwtKey is jwt key we actually used
	JwtKey = utils.GetEnvOrDefault("JWTKEY", DefaultKey)
	// JwtMiddleware is jwt middleware for iris
	JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(ctx iris.Context, err string) {
			ctx.StatusCode(iris.StatusUnauthorized)
			utils.ResponseErr(ctx, errors.New(err))
		},
	})
)

// JwtEncode create a token from user
func JwtEncode(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
		"joined":   user.Joined,
	})

	return token.SignedString([]byte(JwtKey))
}

// JwtDecode get user from a jwt token
func JwtDecode(token *jwt.Token) (*models.User, error) {
	var user models.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Username = claims["username"].(string)
		user.Password = claims["password"].(string)
		user.Joined = claims["joined"].(string)
	} else {
		return nil, ErrJwtDecode
	}
	return &user, nil
}
