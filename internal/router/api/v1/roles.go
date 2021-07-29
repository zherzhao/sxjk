package v1

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"webconsole/global"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
	"github.com/impact-eintr/WebKits/erbac"
)

// UpdateRoles 更新权限数据接口
// @Summary 更新权限数据接口
// @Description 请求后可以跟新权限数据 构造一个新的权限json 作为 body
// @Tags 权限相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=map[string][]string}
// @Router /api/v1/home/roles [post]
func UpdateRoles(c *gin.Context) {
	json := make(map[string][]string) //注意该结构接受的内容
	c.BindJSON(&json)
	tmpRBAC := erbac.NewRBAC()
	tmpPermissions := make(erbac.Permissions)

	// Build roles and add them to eRBAC instance
	for rid, pids := range json {
		role := erbac.NewStdRole(rid)
		for _, pid := range pids {
			_, ok := tmpPermissions[pid]
			if !ok {
				tmpPermissions[pid] = erbac.NewStdPermission(pid)
			}
			role.Assign(tmpPermissions[pid])
		}
		tmpRBAC.Add(role)
	}
	// Load inheritance information
	var jsonInher map[string][]string
	if err := erbac.LoadJson(
		global.RBACSetting.DefaultInherFile, &jsonInher); err != nil {
	}

	// Assign the inheritance relationship
	for rid, parents := range jsonInher {
		if err := tmpRBAC.SetParents(rid, parents); err != nil {
			respcode.ResponseError(c, respcode.CodeInvalidInher)
		}
	}

	global.Auth.Lock()
	global.Auth.RBAC = tmpRBAC
	global.Auth.Permissions = tmpPermissions

	// 保存权限文件
	err := tmpRBAC.SaveUserRBACWithSort(
		global.RBACSetting.CustomerRoleFile, global.RBACSetting.CustomerInherFile)
	if err != nil {
		log.Fatalln(err)
	}

	global.Auth.Unlock()

}

type test struct {
	Root    []string `json:"root"`
	Manager []string `json:"manager"`
	User    []string `json:"user"`
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
	global.Auth.RLock()
	f, err := os.Open(global.RBACSetting.CustomerRoleFile)
	if err != nil {
		log.Println(err)
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}

	t := new(test)
	err = json.Unmarshal(b, t)
	if err != nil {
		log.Println(err)
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}

	global.Auth.RUnlock()

	respcode.ResponseSuccess(c, t)

}

// DefaultRoles 恢复默认权限接口
// @Summary 恢复默认权限接口
// @Description 请求后可以拿到用户数据
// @Tags 权限相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=string}
// @Router /api/v1/home/roles [head]
func DefaultRoles(c *gin.Context) {
	var err error
	global.Auth.Lock()
	defer global.Auth.Unlock()

	global.Auth.RBAC, global.Auth.Permissions, err = erbac.BuildRBAC(
		global.RBACSetting.DefaultRoleFile, global.RBACSetting.DefaultInherFile)
	if err != nil {
		log.Println(err)
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}

	err = global.Auth.RBAC.SaveUserRBACWithSort(
		global.RBACSetting.CustomerRoleFile, global.RBACSetting.CustomerInherFile)
	if err != nil {
		log.Println(err)
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}
	respcode.ResponseSuccess(c, nil)

}
