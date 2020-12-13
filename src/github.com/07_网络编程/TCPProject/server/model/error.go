package model

import "errors"

var (
	ErrorUserNotExists = errors.New("用户不存在")
	ErrorUserExists    = errors.New("用户已存在")
	ErrorUserPwd       = errors.New("用户ID或密码错误")
)
