package utils

import (
	"github.com/codehack/scrypto"
	"github.com/gost-c/gost/internal/types"
	"github.com/kataras/iris"
	"github.com/oklog/ulid"
	"math/rand"
	"os"
	"time"
)

var (
	entropy = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// GetEnvOrDefault return a env value or a default if not exists
func GetEnvOrDefault(key, d string) string {
	v := os.Getenv(key)
	if v == "" {
		return d
	}
	return v
}

// HashPassword is helper function to hash password
func HashPassword(password string) (string, error) {
	return scrypto.Hash(password)
}

// CheckPassword can compare password and hashed password
func CheckPassword(pass, hashed string) bool {
	return scrypto.Compare(pass, hashed)
}

// Uuid return a uuid
func Uuid() string {
	return ulid.MustNew(ulid.Now(), entropy).String()
}

// ResponseErr is a helper function response error message in json format
func ResponseErr(ctx iris.Context, err error) {
	resp := &types.Response{
		Success: false,
		Message: err.Error(),
	}
	ctx.JSON(resp)
}

// ResponseData is a helper function response data in json format
func ResponseData(ctx iris.Context, data interface{}) {
	resp := &types.Response{
		Success: true,
		Data:    data,
	}
	ctx.JSON(resp)
}
