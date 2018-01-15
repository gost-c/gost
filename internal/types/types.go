package types

import (
	"errors"
	"regexp"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	validUsername      = regexp.MustCompile(`^[a-zA-Z0-9_]{6,20}$`)
	validaPassword     = regexp.MustCompile(`^[a-zA-Z0-9!"#$%&'()*+,-./:;<=>?@\[\\\]^_{|} ~]{6,20}$`)
	ErrUsernameInvalid = errors.New("Username length should > 6 and < 20, only support character, numbers and '_'")
	ErrPasswordInvalid = errors.New("Password's length should > 6 and < 20")
)

func (u *UserForm) Validate() (bool, error) {
	if !validUsername.MatchString(u.Username) {
		return false, ErrUsernameInvalid
	}
	if !validaPassword.MatchString(u.Password) {
		return false, ErrPasswordInvalid
	}
	return true, nil
}
