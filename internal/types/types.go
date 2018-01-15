package types

import (
	"errors"
	"regexp"
)

// Response is common api response struct
type Response struct {
	// Success is if a success response
	Success bool `json:"success"`
	// Message is error message
	Message string `json:"message"`
	// Data is response data
	Data interface{} `json:"data"`
}

// UserForm is struct for register and login post
type UserForm struct {
	// Username is user name
	Username string `json:"username"`
	// Password is user password
	Password string `json:"password"`
}

var (
	validUsername  = regexp.MustCompile(`^[a-zA-Z0-9_]{6,20}$`)
	validaPassword = regexp.MustCompile(`^[a-zA-Z0-9!"#$%&'()*+,-./:;<=>?@\[\\\]^_{|} ~]{6,20}$`)
	// ErrUsernameInvalid is error message for invalidate username
	ErrUsernameInvalid = errors.New("Username length should > 6 and < 20, only support character, numbers and '_'")
	// ErrPasswordInvalid is error message for invalidate password
	ErrPasswordInvalid = errors.New("Password's length should > 6 and < 20")
)

// Validate is UserForm validator
func (u *UserForm) Validate() (bool, error) {
	if !validUsername.MatchString(u.Username) {
		return false, ErrUsernameInvalid
	}
	if !validaPassword.MatchString(u.Password) {
		return false, ErrPasswordInvalid
	}
	return true, nil
}
