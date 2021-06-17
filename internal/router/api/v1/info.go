package v1

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"webconsole/global"
	"webconsole/internal/dao/database"
	"webconsole/internal/dao/webcache"
	"webconsole/pkg/respcode"

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
// @Param infotype path string true "查询类型 : road(路)  bridge(桥) tunnel(隧道) service(服务区) portal(收费门架) toll(收费站)"
// @Param level path int true "查询等级 : 0(高速) 1(一级) 2(二级) 3(三级) 4(四级) 5(等外)"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{msg=string,data=string}
// @Router /api/v1/info/{infotype}/{year}/{level} [get]
func (this *Info) GetUpdateInfo(c *gin.Context) {
	infotype := c.GetString("infotype")
	year := c.GetString("year")
	countnum := c.GetInt("count")

	var info string
	var err error
	switch infotype {
	case "road":
		info, err = database.RoadInfo(year, countnum)
	case "bridge":
		info, err = database.BridgeInfo(year, countnum)
	case "tunnel":
		info, err = database.TunnelInfo(year, countnum)
	case "service":
		info = database.FInfo(year)
	case "portal":
		info = database.MInfo(year)
	case "toll":
		info = database.SInfo(year)
	default:
		err = errors.New("请求路径错误")
	}

	if err != nil {
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, err.Error())
		return
	}

	respcode.ResponseSuccess(c, info)

	key := "/" + c.Param("infotype") + "/" + c.Param("year") + "/" + c.Param("count")

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

// QueryInfo 查询数据库数据接口
// @Summary 获取查询数据
// @Description 获取数据库原始数据接口 访问后会更新缓存
// @Tags 缓存相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param infotype path string true "查询类型 : road(路)  bridge(桥) tunnel(隧道) service(服务区) portal(收费门架) toll(收费站)"
// @Param level path int true "查询等级 : 0(高速) 1(一级) 2(二级) 3(三级) 4(四级) 5(等外)"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{msg=string,data=string}
// @Router /api/v1/info/{infotype}/{level} [get]
func (this *Info) QueryInfo(c *gin.Context) {
	infotype := c.GetString("infotype")
	countnum := c.GetInt("count")
	column := c.GetString("column")
	value := c.GetString("value")

	var info string

	switch infotype {
	case "road":
		info = database.RoadQuery(countnum, column, value)
	case "bridge":
		info = database.BridgeQuery(countnum, column, value)
	case "tunnel":
		info = database.TunnelQuery(countnum, column, value)
	case "service":
		info = database.FQuery(column, value)
	case "portal":
		info = database.MQuery(column, value)
	case "toll":
		info = database.SQuery(column, value)
	}

	respcode.ResponseSuccess(c, info)
}
