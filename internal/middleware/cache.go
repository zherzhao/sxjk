package middleware

import (
	"sync"

	"github.com/gin-gonic/gin"
)

func ClearCache(r *gin.Engine, m sync.Map) gin.HandlerFunc {
	return func(c *gin.Context) {
		m.Range(func(k, v interface{}) bool {
			c.Request.URL.Path = "/api/v1/cache/hit" + v.(string) // 将请求的URL修改
			c.Request.Method = "DELETE"
			r.HandleContext(c)
			return true
		})

	}
}
