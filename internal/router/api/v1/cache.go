package v1

import (
	"webconsole/internal/dao/webcache"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
	"github.com/impact-eintr/WebKits/encoding"
)

// CacheCheck 检查缓存命中接口
// @Summary 检查缓存命中接口
// @Description 检查缓存中是否有请求的值 有就返回没有将请求转发
// @Tags 缓存相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param infotype path string true "查询类型"
// @Param year path string true "查询年份 格式: 202X "
// @Param level path string true "查询等级"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=string}
// @Router /api/v1/cache/hit/{infotype}/{year}/{level} [get]
func CacheCheck(c *gin.Context) {
	key := c.GetString("userUnit") + c.Param("key")
	if key == "" {
		respcode.ResponseError(c, respcode.CodeInvalidParam)
		return
	}

	b := webcache.CacheCheck(key)
	if len(b) > 0 {
		respcode.ResponseSuccess(c, encoding.Bytes2str(b))
		c.Abort()
	} else {
		c.Next()
	}

}

func CacheDelete(c *gin.Context) {
	key := c.GetString("userUnit") + c.Param("key")
	if key == "" {
		respcode.ResponseError(c, respcode.CodeInvalidParam)
		return
	}
	webcache.CacheDelete(key)

}
