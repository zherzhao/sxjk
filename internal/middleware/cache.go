package middleware

import "github.com/gin-gonic/gin"

func ClearCache(r *gin.Engine, m map[string]struct{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		for v := range m {
			c.Request.URL.Path = "/api/v1/cache/hit" + v // 将请求的URL修改
			c.Request.Method = "DELETE"
			r.HandleContext(c)
		}
	}
}
