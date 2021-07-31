package middleware

import (
	"strconv"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
)

// 获取路径的中间件
func PathParse(c *gin.Context) {
	var countflag, typeflag bool

	year := c.Param("year")
	count, err := strconv.Atoi(c.Param("count"))
	if err == nil && count < 6 {
		countflag = true
	}

	infotype := c.Param("infotype")
	for _, v := range []string{"road", "bridge", "tunnel", "service", "portal", "toll"} {
		if infotype == v {
			typeflag = true
			break
		}
	}

	if countflag && typeflag {
		c.Set("infotype", infotype)
		c.Set("year", year)
		c.Set("count", count)
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
