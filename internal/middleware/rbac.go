package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
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

func IServerRBAC(p string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("userRole")
		permission := []string{"query-data", "query-datas"}

		if global.Auth.RBAC.IsGranted(role, global.Auth.Permissions[permission[1]], nil) {
			c.Next()

		} else if global.Auth.RBAC.IsGranted(role, global.Auth.Permissions[permission[0]], nil) {
			defer c.Abort()
			if c.Request.Method == "POST" {
				if p == "search" {
					search := new(model.IServerSearchReq)
					b, err := ioutil.ReadAll(c.Request.Body)
					if err != nil {
						return
					}

					reg, err := regexp.Compile(`'([a-zA-Z]+?)':`)
					index := reg.FindAllIndex(b, 7)
					for _, r := range index {
						log.Println(string(b[r[0]:r[1]]))
						b[r[0]] = 34
						b[r[1]-2] = 34
					}

					json.Unmarshal(b, search)
					log.Println(search)

					attr := search.QueryParameter.AttributeFilter
					if strings.HasPrefix(attr, "路线名称") {
						unit := c.GetString("userUnit")
						extra := fmt.Sprintf("交控单位名称 like '%s' and ", "%25"+unit+"%25")
						search.QueryParameter.AttributeFilter = extra + attr

					} else if strings.HasPrefix(attr, "桥梁名称") ||
						strings.HasPrefix(attr, "收费站名称") ||
						strings.HasPrefix(attr, "隧道名称") ||
						strings.HasPrefix(attr, "门架名称") {

						unit := c.GetString("userUnit")
						extra := fmt.Sprintf("所在地级市 like '%s' and ", "%25"+unit+"%25")
						search.QueryParameter.AttributeFilter = extra + attr

					} else if strings.HasPrefix(attr, "名称") {
						zap.L().Warn("没有权限")
						respcode.ResponseError(c, respcode.CodeUserPermissionDenied)
						return
					} else {
						zap.L().Warn("小心sql注入")

					}
					b, _ = json.Marshal(search)

					reg, err = regexp.Compile(`"([a-zA-Z]+?)":`)
					index = reg.FindAllIndex(b, 7)
					for _, r := range index {
						b[r[0]] = 39
						b[r[1]-2] = 39
					}

					c.Request.Header.Set("Content-Length", fmt.Sprintf("%d", len(string(b))))
					c.Request.Body = ioutil.NopCloser(bytes.NewReader(b))

				} else if p == "query" {
					query := new(model.IServerQueryReq)
					b, err := ioutil.ReadAll(c.Request.Body)
					if err != nil {
						return
					}
					reg, err := regexp.Compile(`'([a-zA-Z]+?)':`)
					index := reg.FindAllIndex(b, 7)
					for _, r := range index {
						b[r[0]] = 34
						b[r[1]-2] = 34
					}

					json.Unmarshal(b, query)

					attr := query.QueryParameter.AttributeFilter
					if strings.HasPrefix(attr, "路线名称") {
						unit := c.GetString("userUnit")
						extra := fmt.Sprintf("交控单位名称 like '%s' and ", "%25"+unit+"%25")
						query.QueryParameter.AttributeFilter = extra + attr

					} else if strings.HasPrefix(attr, "桥梁名称") ||
						strings.HasPrefix(attr, "收费站名称") ||
						strings.HasPrefix(attr, "隧道名称") ||
						strings.HasPrefix(attr, "门架名称") {

						unit := c.GetString("userUnit")
						extra := fmt.Sprintf("所在地级市 like '%s' and ", "%25"+unit+"%25")
						query.QueryParameter.AttributeFilter = extra + attr

					} else if strings.HasPrefix(attr, "名称") {
						zap.L().Warn("没有权限")
						respcode.ResponseError(c, respcode.CodeUserPermissionDenied)
						return
					} else {
						zap.L().Warn("小心sql注入")

					}
					b, _ = json.Marshal(query)

					reg, err = regexp.Compile(`"([a-zA-Z]+?)":`)
					index = reg.FindAllIndex(b, 7)
					for _, r := range index {
						b[r[0]] = 39
						b[r[1]-2] = 39
					}
					log.Println(string(b))

					c.Request.Header.Set("Content-Length", fmt.Sprintf("%d", len(string(b))))
					c.Request.Body = ioutil.NopCloser(bytes.NewReader(b))

				}
			}
			c.Next()

		} else {
			respcode.ResponseError(c, respcode.CodeUserPermissionDenied)
			c.Abort()

		}
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
