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

		if c.GetInt("count") > 0 {
			permission = infotype + "-records"
		} else {
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
