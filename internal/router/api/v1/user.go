package v1

import (
	"log"
	"strconv"
	"webconsole/internal/dao/database"
	"webconsole/internal/model"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
)

// GetUesrs 获取用户数据接口
// @Summary 获取用户数据接口
// @Description 请求后可以拿到用户数据
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=[]model.RespUser}
// @Router /api/v1/home/users [get]
func GetUsers(c *gin.Context) {
	users, err := database.GetUsersHandler()
	if err != nil {
		log.Println(err)
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}
	respcode.ResponseSuccess(c, users)

}

// QueryUesrs 查询用户数据接口
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
func QueryUsers(c *gin.Context) {
	column := c.GetString("column")
	value := c.GetString("value")
	users, err := database.QueryUsersHandler(column, value)
	if err != nil {
		log.Println(err)
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}
	respcode.ResponseSuccess(c, users)

}

// UpdateUesrs 更新用户数据接口
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
func UpdateUsers(c *gin.Context) {
	var err error
	reqUser := new(model.RespUser)
	user := new(model.User)

	c.ShouldBindJSON(reqUser)

	id, _ := strconv.ParseInt(reqUser.UserIDStr, 10, 64)
	user.UserID = id
	user.Username = reqUser.Username
	user.Unit = reqUser.Unit
	user.Role = reqUser.Role
	user.Password, err = database.UserPassword(user.UserID)
	if err != nil {
		log.Println(err)
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}

	err = database.UpdateUsersHandler(user)
	if err != nil {
		log.Println(err)
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}
	respcode.ResponseSuccess(c, nil)

}

// DeleteUesrs 删除用户接口
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
func DeleteUsers(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, err)
		return
	}
	tmpUser := new(model.User)
	tmpUser.UserID = id

	err = database.DeleteUsersHandler(tmpUser)
	if err != nil {
		log.Println(err)
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}
	respcode.ResponseSuccess(c, nil)

}
