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
	r.GET("/api/hello").
		SetDebug(true).
		Run(GinEngine(), func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, res.Code, http.StatusUnauthorized)
		})
}
