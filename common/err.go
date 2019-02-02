package common

import (
	"errors"
	"github.com/zcong1993/libgo/gin/ginerr"
	"net/http"
)

func createErr(statusCode int, code string, errors interface{}) ginerr.ApiError {
	return ginerr.NewDefaultError(statusCode, code, code, errors)
}

var (
	DUPLICATE_USER               = ginerr.NewDefaultError(http.StatusBadRequest, "DUPLICATE_USER", "DUPLICATE_USER", nil)
	INTERNAL_ERROR               = ginerr.NewDefaultError(http.StatusInternalServerError, "", "", nil)
	INVALID_USERNAME_OR_PASSWORD = createErr(http.StatusUnauthorized, "INVALID_USERNAME_OR_PASSWORD", nil)
	NOT_FOUND_ERROR              = createErr(http.StatusNotFound, "NOT_FOUND", nil)
	NOT_UUID                     = createErr(http.StatusBadRequest, "NOT_A_UUID", nil)
)

var TOKEN_ERROR = map[string]string{"code": "TOKEN_ERR_OR_EXPIRED", "message": "TOKEN_ERR_OR_EXPIRED"}

var (
	ErrExpired = errors.New("TOKEN_EXPIRED")
)

type ErrResp struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func CreateInvalidErr(errors interface{}) ginerr.ApiError {
	return ginerr.NewDefaultError(http.StatusBadRequest, "INVALID_PARAMS", "INVALID_PARAMS", errors)
}
