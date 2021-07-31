package v1

import (
	"strconv"
	"webconsole/global"
	"webconsole/internal/dao/database"
	"webconsole/internal/model"
	"webconsole/internal/service"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// GetUesrsHandler 获取用户数据接口
// @Summary 获取用户数据接口
// @Description 请求后可以拿到用户数据
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=[]model.RespUser}
// @Router /api/v1/home/users [get]
func GetUsersHandler(c *gin.Context) {
	users, err := service.GetUsers()
	if err != nil {
		zap.L().Error("未知错误", zap.Error(err))
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}
	respcode.ResponseSuccess(c, users)

}

// QueryUesrsHandler 查询用户数据接口
// @Summary 查询用户数据接口
// @Description 请求后可以获取指定的用户数据
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param column query string true "属性"
// @Param value query string true "值"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=[]model.RespUser}
// @Router /api/v1/home/users/query [get]
func QueryUsersHandler(c *gin.Context) {
	column := c.GetString("column")
	value := c.GetString("value")
	users := []model.UserResp{}
	var err error

	switch column {
	case "user_id_str":
		id, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			zap.L().Error("解析错误", zap.Error(err))
			respcode.ResponseError(c, respcode.CodeInvalidPath)
			return
		}
		users, err = service.QueryUsers("user_id", id)
	default:
		users, err = service.QueryUsers(column, value)
	}

	if err != nil {
		zap.L().Error("未知错误", zap.Error(err))
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}
	respcode.ResponseSuccess(c, users)

}

// UpdateUesrsHandler 更新用户数据接口
// @Summary 更新用户数据接口
// @Description 请求后可以将post请求body中提供的用户新数据替换原来的用户数据
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param 用户信息 body model.RespUser true "用户信息"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=string}
// @Router /api/v1/home/users [post]
func UpdateUsersHandler(c *gin.Context) {
	var err error
	reqUser := new(model.UserReq)
	if err := c.ShouldBindJSON(reqUser); err != nil {
		zap.L().Error("update user with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			respcode.ResponseError(c, respcode.CodeInvalidParam)
			return
		}
		respcode.ResponseErrorWithMsg(c, respcode.CodeInvalidParam, errs.Translate(global.Trans))
		return
	}

	id, err := strconv.ParseInt(reqUser.UserID, 10, 64)
	if err != nil {
		zap.L().Error("解析错误", zap.Error(err))
		respcode.ResponseError(c, respcode.CodeInvalidPath)
		return
	}

	err = service.UpdateUser(reqUser, id)
	if err != nil {
		if err == database.ErrorInvalidUnit {
			zap.L().Error("user 想要获取交科权限", zap.Error(err))
			respcode.ResponseError(c, respcode.CodeInvalidUnit)
			return
		}
		zap.L().Error("未知错误", zap.Error(err))
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}
	respcode.ResponseSuccess(c, nil)

}

// DeleteUesrsHandler 删除用户接口
// @Summary 删除用户数据接口
// @Description 请求后可以删除用户数据（直接从数据库删除）
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path string true "用户id"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=string}
// @Router /api/v1/home/users/{id} [delete]
func DeleteUsersHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		zap.L().Error("解析错误", zap.Error(err))
		respcode.ResponseError(c, respcode.CodeInvalidPath)
		return
	}

	err = service.DeleteUser(id)
	if err != nil {
		zap.L().Error("未知错误", zap.Error(err))
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}
	respcode.ResponseSuccess(c, nil)

}
