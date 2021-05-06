package v1

import (
	"errors"
	"fmt"
	"webconsole/global"
	"webconsole/internal/dao/database"
	"webconsole/internal/model"
	"webconsole/internal/service"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const (
	CTXUserID string = "userID"
)

// SignUpHandler 注册接口
// @Summary 处理注册请求
// @Description 首页注册
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{msg=string,data=string}
// @Router /api/v1/sigin [post]
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
	if err := service.SignUp(p); err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		if errors.Is(err, database.ErrorUserExist) {
			respcode.ResponseError(c, respcode.CodeUserExist)
		} else {
			respcode.ResponseError(c, respcode.CodeServerBusy)

		}
		return
	}

	// 3. 返回响应
	respcode.ResponseSuccess(c, nil)

}

// LoginHandler 登录接口
// @Summary 处理登录请求
// @Description 首页登录
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param logininfo body model.ParamLogin true "用户登录信息"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{msg=string,data=string}
// @Router /api/v1/login [post]
func LoginHandler(c *gin.Context) {
	// 获取请求参数以及参数校验

	p := new(model.ParamLogin)

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			respcode.ResponseError(c, respcode.CodeInvalidParam)
			return
		}

		respcode.ResponseErrorWithMsg(c, respcode.CodeInvalidParam, errs.Translate(global.Trans))
		return
	}

	// 2. 业务处理
	userID, aToken, err := service.Login(p)
	if err != nil {
		zap.L().Error("登录失败", zap.Error(err))
		if errors.Is(err, database.ErrorUserNotExist) {
			respcode.ResponseError(c, respcode.CodeUserNotExist)

		} else if errors.Is(err, database.ErrorInvalidPassword) {
			respcode.ResponseError(c, respcode.CodeInvalidPassword)

		} else {
			respcode.ResponseError(c, respcode.CodeServerBusy)

		}
		return
	}

	data := map[string]string{"token": aToken, "userid": fmt.Sprintf("%d", userID), "username": p.UserName}

	// 3. 返回响应
	respcode.ResponseSuccess(c, data)

}
