package middleware

import (
	"webconsole/global"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
)

func RBACMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("userRole")
		if global.RBAC.IsGranted(role, global.Permissions["read-road-record"], nil) {
			respcode.ResponseSuccess(c, global.Permissions)
			c.Next()
		} else {
			respcode.ResponseError(c, respcode.CodeUserPermissionDenied)
			c.Abort()
		}
	}
}
