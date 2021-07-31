package database

import "errors"

// 错误码
var (
	ErrorUserExist       = errors.New("用户已经存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	ErrorNoUser          = errors.New("数据库无用户")
	ErrorVerify          = errors.New("验证失败")
	ErrorInvalidUnit     = errors.New("非法单位")
)
