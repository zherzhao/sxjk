package middleware

import (
	"errors"
	"strconv"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
)

// 获取路径的中间件
func PathParse(c *gin.Context) {
	infotype := c.Param("infotype")
	year := c.Param("year")
	count, err := strconv.Atoi(c.Param("count"))
	if err != nil || count > 6 {
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy,
			errors.New("参数非法").Error())
		c.Abort()
	}

	c.Set("infotype", infotype)
	c.Set("year", year)
	c.Set("count", count)

	c.Next()
}

// 获取查询参数的中间件
func QueryParse(c *gin.Context) {
	column := c.Query("column")
	value := c.Query("value")

	c.Set("column", column)
	c.Set("value", value)

	c.Next()
}
