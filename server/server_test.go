package server

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/appleboy/gofight.v2"
	"net/http"
	"testing"
)

func TestGinEngine(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gofight.New()
	r.GET("/gist/1").
		SetDebug(true).
		Run(GinEngine(), func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, res.Code, http.StatusOK)
		})
	r.POST("/api/create").
		SetDebug(true).
		Run(GinEngine(), func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, res.Code, http.StatusUnauthorized)
		})
	r.POST("/register").
		SetDebug(true).
		SetJSON(gofight.D{"username": "zcxsxs?", "password": "xsaxsaxsx"}).
		Run(GinEngine(), func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, res.Code, http.StatusOK)
			assert.JSONEq(t, res.Body.String(), `{"code": "400", "msg": "Username length should > 6 and < 20, only support character, numbers and '_'"}`)
		})
	r.POST("/register").
		SetDebug(true).
		SetJSON(gofight.D{"username": "zcxsxs", "password": "xsa"}).
		Run(GinEngine(), func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, res.Code, http.StatusOK)
			assert.JSONEq(t, res.Body.String(), `{"code": "400", "msg": "Password's length should > 6 and < 20"}`)
		})
	r.POST("/login").
		SetDebug(true).
		SetJSON(gofight.D{"username": "zcxsxs", "password": "xsa"}).
		Run(GinEngine(), func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, res.Code, http.StatusUnauthorized)
			assert.JSONEq(t, res.Body.String(), `{"code": 401, "message": "Incorrect Username / Password"}`)
		})
}
