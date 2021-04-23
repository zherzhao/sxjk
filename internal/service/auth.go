package service

import (
	"errors"
	"webconsole/global"
	"webconsole/internal/model"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const (
	CTXUserID string = "userID"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数 参数校验
	p := new(model.ParamSignUp)

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			respcode.ResponseError(c, respcode.CodeInvalidParam)
			return
		}

		respcode.ResponseErrorWithMsg(c, respcode.CodeInvalidParam, errs.Translate(global.Trans))
		return
	}

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			respcode.ResponseError(c, respcode.CodeUserExist)
		} else {
			respcode.ResponseError(c, respcode.CodeServerBusy)

		}
		return
	}

	// 3. 返回响应
	respcode.ResponseSuccess(c, nil)

}

func LoginHandler(c *gin.Context) {
	// 获取请求参数以及参数校验

	p := new(model.ParamLogin)

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			respcode.ResponseError(c, respcode.CodeInvalidParam)
			return
		}

		respcode.ResponseErrorWithMsg(c, respcode.CodeInvalidParam, errs.Translate(trans))
		return
	}

	// 2. 业务处理
	aToken, err := logic.Login(p)
	if err != nil {
		zap.L().Error("登录失败", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			respcode.ResponseError(c, respcode.CodeUserNotExist)

		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			respcode.ResponseError(c, respcode.CodeInvalidPassword)

		} else {
			respcode.ResponseError(c, respcode.CodeServerBusy)

		}
		return
	}

	data := map[string]string{"token": aToken}

	// 3. 返回响应
	respcode.ResponseSuccess(c, data)

}
