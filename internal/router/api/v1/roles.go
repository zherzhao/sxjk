package v1

import (
	"webconsole/global"
	"webconsole/internal/dao/rbac"
	"webconsole/internal/service"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// UpdateRolesHandler 更新权限数据接口
// @Summary 更新权限数据接口
// @Description 请求后可以跟新权限数据 构造一个新的权限json 作为 body
// @Tags 权限相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=map[string][]string}
// @Router /api/v1/home/roles [post]
func UpdateRolesHandler(c *gin.Context) {
	json := make(map[string][]string) //注意该结构接受的内容
	if err := c.BindJSON(&json); err != nil {
		zap.L().Error("序列化失败", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			respcode.ResponseError(c, respcode.CodeInvalidParam)
		} else {
			respcode.ResponseErrorWithMsg(c, respcode.CodeInvalidParam, errs.Translate(global.Trans))
		}
		c.Abort()
		return
	}

	err := service.UpdateRoles(json)
	if err != nil {
		if err == rbac.ErrorInvalidInher {
			zap.L().Error("出现环继承", zap.Error(err))
			respcode.ResponseError(c, respcode.CodeInvalidInher)
			c.Abort()
			return
		} else {
			zap.L().Error("未知错误", zap.Error(err))
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				respcode.ResponseError(c, respcode.CodeServerBusy)
			} else {
				respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, errs.Translate(global.Trans))
			}
			c.Abort()
			return
		}
	}
	respcode.ResponseSuccess(c, nil)

}

// GetRoles 获取权限数据接口
// @Summary 获取权限数据接口
// @Description 请求后可以拿到权限数据
// @Tags 权限相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=map[string][]string}
// @Router /api/v1/home/roles [get]
func GetRoles(c *gin.Context) {
	r, err := service.GetRoles()
	if err != nil {
		if err == rbac.ErrorRBACNotFound {
			zap.L().Error("未找到用户自定义权限文件", zap.Error(err))
			respcode.ResponseError(c, respcode.CodeServerBusy)
			return
		}
		zap.L().Error("未知错误", zap.Error(err))
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}

	respcode.ResponseSuccess(c, r)

}

// DefaultRolesHandler 恢复默认权限接口
// @Summary 恢复默认权限接口
// @Description 请求后可以拿到用户数据
// @Tags 权限相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=string}
// @Router /api/v1/home/roles [head]
func DefaultRolesHandler(c *gin.Context) {
	if err := service.DefaultRoles(); err != nil {
		zap.L().Error("未知错误", zap.Error(err))
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}
	respcode.ResponseSuccess(c, nil)

}
