package utils_test

import (
	"github.com/gost-c/gost/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var password = "password"

func TestCheckPassword(t *testing.T) {
	assert2 := assert.New(t)
	h1, err := utils.HashPassword(password)
	assert2.Nil(err, "hash should be success")
	assert2.True(utils.CheckPassword(password, h1), "check should return true")
}
