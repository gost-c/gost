package jwt_test

import (
	"github.com/gost-c/gost/internal/jwt"
	"testing"
	"github.com/stretchr/testify/assert"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gost-c/gost/internal/models"
)

var user = &models.User{
	Username: "zcong",
	Password: "password",
	Joined: "2018-01-03",
}

func TestJwt(t *testing.T) {
	assert2 := assert.New(t)
	j, err := jwt.JwtEncode(user)
	assert2.Nil(err, "Encode error should be nil")
	assert2.NotEmpty(j, "Encode token should not empty")
	keyfunc := func(token *jwt2.Token) (interface{}, error) {
		return []byte(jwt.JwtKey), nil
	}
	tk, err := jwt2.Parse(j, keyfunc)
	assert2.Nil(err, "Token parse should not failed")
	u, err := jwt.JwtDecode(tk)
	assert2.Nil(err, "Decode error should be nil")
	assert2.Equal(user, u, "Decode should be equal to encode user")
}
