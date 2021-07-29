package v1

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
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
// @Tags 注册api
// @Accept application/json
// @Produce application/json
// @Param 注册信息 body model.ParamSignUp true "用户注册信息"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=string}
// @Router /api/v1/signup [post]
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
	// 2. 验证注册码
	if p.VerifyCode != global.VerifyCode.Code {
		respcode.ResponseError(c, respcode.CodeInvalidVerifyCode)
		return
	}

	// 3. 业务处理
	if err := service.SignUp(p); err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		if errors.Is(err, database.ErrorUserExist) {
			respcode.ResponseError(c, respcode.CodeUserExist)
		} else {
			respcode.ResponseError(c, respcode.CodeServerBusy)

		}
		return
	}

	// 4. 返回响应
	respcode.ResponseSuccess(c, nil)

}

// SignUpCode 注册码申请接口
// @Summary 申请一个注册码
// @Description 申请一个注册码 重复请求上一个就会失效 重启后端服务也会失效
// @Tags 注册api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=string}
// @Router /api/v1/signup [get]
func SignUpCode(c *gin.Context) {
	code := genValidateCode(6)
	respcode.ResponseSuccess(c, code)
	global.VerifyCode.Ch <- code
	global.VerifyCode.T.Reset(60 * 10 * time.Second)
}

func genValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// LoginHandler 登录接口
// @Summary 处理登录请求
// @Description 首页登录
// @Tags 登陆api
// @Accept application/json
// @Produce application/json
// @Param 登录信息 body model.ParamLogin true "用户登录信息"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=string}
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
	userID, userRole, aToken, err := service.Login(p)
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

	data := map[string]string{
		"token":    aToken,
		"userrole": userRole,
		"userid":   fmt.Sprintf("%d", userID),
		"username": p.UserName}

	// 3. 返回响应
	respcode.ResponseSuccess(c, data)

}
