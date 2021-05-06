package v1

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"webconsole/global"
	"webconsole/internal/dao/webcache"

	"github.com/gin-gonic/gin"
)

type Info struct {
}

func NewInfo() Info {
	return Info{}
}

// GetUpdataInfo 获取数据库原始数据接口 访问后会更新缓存
// @Summary 更新缓存接口
// @Description 获取数据库原始数据接口 访问后会更新缓存
// @Tags 缓存相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param infotype path string true "查询类型"
// @Param level path int true "查询等级"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{msg=string,data=string}
// @Router /api/v1/info/{infotype}/{level} [get]
func (this *Info) GetUpdateInfo(c *gin.Context) {
	info := c.GetString("info")

	c.JSON(http.StatusOK, info) // 向浏览器返回数据

	key := "/" + c.Param("infotype") + "/" + c.Param("count")

	// 如果是缓存在磁盘中
	if global.CacheSetting.CacheType == "disk" {
		go webcache.UpdataCache(key, info)

	} else {
		c.Set("type", "mem")
		c.Request.URL.Path = "/api/v1/cache/update" + key //将请求的URL修改
		c.Request.Method = http.MethodPut
		c.Request.Body = ioutil.NopCloser(bytes.NewReader([]byte(info)))

	}

}
