package middleware

import (
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
)

// 获取路径的中间件
func PathParse(c *gin.Context) {
	var typeflag bool

	year := c.Param("year")

	infotype := c.Param("infotype")
	for _, v := range []string{"road", "bridge", "tunnel", "service", "portal", "toll"} {
		if infotype == v {
			typeflag = true
			break
		}
	}

	if typeflag {
		c.Set("infotype", infotype)
		c.Set("year", year)
		c.Next()
	} else {
		respcode.ResponseError(c, respcode.CodeInvalidPath)
		c.Abort()
	}
}

// 获取查询参数的中间件
func QueryParse(c *gin.Context) {
	column := c.Query("column")
	value := c.Query("value")

	if column == "" || value == "" {
		respcode.ResponseError(c, respcode.CodeInvalidPath)
		c.Abort()
	} else {
		c.Set("column", column)
		c.Set("value", value)
		c.Next()
	}
}
