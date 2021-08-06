package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"webconsole/global"
	"webconsole/internal/model"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecordRBAC(c *gin.Context) {
	role := c.GetString("userRole")
	infotype := c.GetString("infotype")

	permission := infotype + "-record"

	if global.Auth.RBAC.IsGranted(role, global.Auth.Permissions[permission], nil) {
		c.Next()
	} else {
		respcode.ResponseError(c, respcode.CodeUserPermissionDenied)
		c.Abort()
	}
}

func QueryRBAC(c *gin.Context) {
	role := c.GetString("userRole")
	var permission string = "query-data"

	if global.Auth.RBAC.IsGranted(role, global.Auth.Permissions[permission], nil) {
		c.Next()
	} else {
		respcode.ResponseError(c, respcode.CodeUserPermissionDenied)
		c.Abort()
	}
}

func IServerRBAC(c *gin.Context) {
	role := c.GetString("userRole")
	permission := []string{"query-data", "query-datas"}

	if global.Auth.RBAC.IsGranted(role, global.Auth.Permissions[permission[1]], nil) {
		c.Next()

	} else if global.Auth.RBAC.IsGranted(role, global.Auth.Permissions[permission[0]], nil) {
		if c.Request.Method == "POST" {
			test := new(model.IServerReq)
			b, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				return
			}

			tmp := make([]byte, len(b))
			count := 0
			for i, v := range b {
				if v == 39 && count < 8 {
					tmp[i] = 34
					count++
					continue
				}
				tmp[i] = v
			}
			err = json.Unmarshal(tmp, test)
			if err != nil {
				return
			}
			attr := test.QueryParameter.AttributeFilter
			if strings.HasPrefix(attr, "路线名称") {
				unit := c.GetString("userUnit")
				extra := fmt.Sprintf("交控单位名称 like '%s' and ", "%25"+unit+"%25")
				test.QueryParameter.AttributeFilter = extra + attr

			} else if strings.HasPrefix(attr, "桥梁名称") ||
				strings.HasPrefix(attr, "收费站名称") ||
				strings.HasPrefix(attr, "隧道名称") ||
				strings.HasPrefix(attr, "门架名称") {

				unit := c.GetString("userUnit")
				extra := fmt.Sprintf("所在地级市 like '%s' and ", "%25"+unit+"%25")
				test.QueryParameter.AttributeFilter = extra + attr

			} else if strings.HasPrefix(attr, "名称") {
				zap.L().Warn("没有权限")
				respcode.ResponseError(c, respcode.CodeUserPermissionDenied)
				c.Abort() // 这里应该处理成 返回没有结果
				return
			} else {
				zap.L().Warn("小心sql注入")

			}
			count = 0
			tmp, _ = json.Marshal(test)
			for i, v := range tmp {
				if v == 34 && count < 12 {
					if count%4 == 0 || count%4 == 1 || count == 10 || count == 11 {
						tmp[i] = 39
						count++
						continue
					}
					count++
				}
				tmp[i] = v
			}
			c.Request.Header.Set("Content-Length", fmt.Sprintf("%d", len(tmp)))
			c.Request.Body = ioutil.NopCloser(bytes.NewReader([]byte(tmp)))
		}
		c.Next()

	} else {
		respcode.ResponseError(c, respcode.CodeUserPermissionDenied)
		c.Abort()

	}
}

func RoleRBAC(p string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("userRole")
		if global.Auth.RBAC.IsGranted(role, global.Auth.Permissions[p], nil) {
			c.Next()
		} else {
			respcode.ResponseError(c, respcode.CodeUserPermissionDenied)
			c.Abort()
		}
	}
}

func UserRBAC(p string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("userRole")
		if global.Auth.RBAC.IsGranted(role, global.Auth.Permissions[p], nil) {
			c.Next()
		} else {
			respcode.ResponseError(c, respcode.CodeUserPermissionDenied)
			c.Abort()
		}
	}
}

func TableRBAC(p string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("userRole")
		if global.Auth.RBAC.IsGranted(role, global.Auth.Permissions[p], nil) {
			c.Next()
		} else {
			respcode.ResponseError(c, respcode.CodeUserPermissionDenied)
			c.Abort()
		}
	}
}
