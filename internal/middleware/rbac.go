package middleware

import (
	"webconsole/global"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
)

func RBACMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("userRole")
		infotype := c.GetString("infotype")
		var permission string

		switch {
		case c.GetInt("count") > 0:
			permission = infotype + "-records"
		case c.GetInt("count") == 0:
			permission = infotype + "-record"
		}

		if global.RBAC.IsGranted(role, global.Permissions[permission], nil) {
			c.Next()
		} else {
			respcode.ResponseError(c, respcode.CodeUserPermissionDenied)
			c.Abort()
		}

	}
}

func MapRBACMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("mapname", c.Param("mapname"))
		c.Set("year", c.Param("year"))
		role := c.GetString("userRole")
		var permission string = "read-maps"

		if global.RBAC.IsGranted(role, global.Permissions[permission], nil) {
			c.Next()
		} else {
			respcode.ResponseError(c, respcode.CodeUserPermissionDenied)
			c.Abort()
		}

	}
}
