package respcode

type RespCode int

const (
	CodeSuccess RespCode = 1000 + iota

	CodeInvalidParam
	CodeInvalidVerifyCode

	CodeUserExist
	CodeUserPermissionDenied
	CodeUserNotExist
	CodeInvalidPassword

	CodeNullAuth
	CodeInvalidToken
	CodeInvalidInher

	CodeServerBusy

	CodeNotFound
)

var codeMsgMap = map[RespCode]string{
	CodeSuccess:              "success",
	CodeInvalidParam:         "请求参数错误",
	CodeInvalidVerifyCode:    "无效的注册码",
	CodeUserExist:            "用户名已经存在",
	CodeUserNotExist:         "该用户不存在",
	CodeUserPermissionDenied: "用户没有权限",
	CodeInvalidPassword:      "用户密码错误",

	CodeNullAuth:     "请头中auth为空",
	CodeInvalidToken: "无效的Token",
	CodeInvalidInher: "发现环继承",

	CodeServerBusy: "服务器繁忙，通知后端查看日志",

	CodeNotFound: "路由不存在",
}

func (rescode RespCode) Msg() string {
	return codeMsgMap[rescode]
}
