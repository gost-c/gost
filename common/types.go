package common

import "github.com/zcong1993/libgo/gin/ginhelper"

type TokenResp struct {
	Token string `json:"token"`
}

type UserResp struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

// Gost is gost model
type GostResp struct {
	ginhelper.Model
	Description string      `json:"description"`
	Private     bool        `json:"private"`
	Version     int         `json:"version"`
	Files       []*FileResp `json:"files"`
	User        *UserResp   `json:"user"`
}

// File is file model
type FileResp struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}
