package rbac

import "errors"

// 错误码
var (
	ErrorInvalidInher = errors.New("非法的继承关系")
	ErrorRBACNotFound = errors.New("未找到用户自定义权限表")
)
