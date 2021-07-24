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
	err := tmpRBAC.SaveUserRBAC(
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

func GetRoles(c *gin.Context) {
	global.Auth.RLock()
	f, err := os.Open(global.RBACSetting.CustomerRoleFile)
	if err != nil {
		c.JSON(500, nil)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		c.JSON(500, nil)
	}

	t := new(test)
	err = json.Unmarshal(b, t)
	if err != nil {
		log.Println(err)
		c.JSON(500, nil)
	}

	global.Auth.RUnlock()

	c.JSON(200, t)

}

func DefaultRoles(c *gin.Context) {
	var err error
	global.Auth.Lock()
	defer global.Auth.Unlock()

	global.Auth.RBAC, global.Auth.Permissions, err = erbac.BuildRBAC(
		global.RBACSetting.DefaultRoleFile, global.RBACSetting.DefaultInherFile)
	if err != nil {
		log.Println(err)
		c.JSON(500, nil)
		return
	}

	err = global.Auth.RBAC.SaveUserRBAC(
		global.RBACSetting.CustomerRoleFile, global.RBACSetting.CustomerInherFile)
	if err != nil {
		log.Println(err)
		c.JSON(500, nil)
		return
	}
	respcode.ResponseSuccess(c, nil)

}
