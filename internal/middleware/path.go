package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 获取路径的中间件
func PathParse(c *gin.Context) {
	infotype := c.Param("infotype")
	count := c.Param("count")

	c.Set("infotype", infotype)
	c.Set("count", count)

	c.Next()
}

// 获取查询参数的中间件
func QueryParse(c *gin.Context) {
	column := c.Query("column")
	value := c.Query("value")

	fmt.Println(column, value)

	c.Set("column", column)
	c.Set("value", value)

	c.Next()
}
