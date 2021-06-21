package middleware

import (
	"errors"
	"strconv"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
)

// 获取路径的中间件
func PathParse(c *gin.Context) {

	year := c.Param("year")
	count, err := strconv.Atoi(c.Param("count"))
	if err != nil || count >= 6 {
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy,
			errors.New("参数非法").Error())
		c.Abort()
	}
	infotype := c.Param("infotype")
	var flag bool
	for _, v := range []string{"road", "bridge", "tunnel", "service", "protal", "toll"} {
		if infotype == v {
			flag = true
			break
		}
	}

	if flag {
		c.Set("infotype", infotype)
		c.Set("year", year)
		c.Set("count", count)
		c.Next()
	} else {
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, "路径错误")
		c.Abort()
	}

}

// 获取查询参数的中间件
func QueryParse(c *gin.Context) {
	column := c.Query("column")
	value := c.Query("value")

	if column == "" || value == "" {
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy,
			errors.New("缺失查询参数").Error())
		c.Abort()

	} else {
		c.Set("column", column)
		c.Set("value", value)
		c.Next()

	}

}
