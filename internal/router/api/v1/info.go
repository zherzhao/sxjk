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
