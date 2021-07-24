package v1

import (
	"errors"
	"fmt"
	"webconsole/global"
	"webconsole/internal/dao/database"
	"webconsole/internal/dao/webcache"
	"webconsole/internal/model"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
)

// GetInfo 获取数据库原始数据接口 访问后会更新缓存
// @Summary 更新缓存接口
// @Description 获取数据库原始数据接口 访问后会更新缓存
// @Tags 数据操作api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param infotype path string true "查询类型 : road(路)  bridge(桥) tunnel(隧道) service(服务区) portal(收费门架) toll(收费站)"
// @Param year path string true "查询年份 格式: 202X "
// @Param level path int true "查询等级 : 0(高速) 1(一级) 2(二级) 3(三级) 4(四级) 5(等外)"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=string}
// @Router /api/v1/data/info/{infotype}/{year}/{level} [get]
func GetInfo(c *gin.Context) {
	infotype := c.GetString("infotype")
	year := c.GetString("year")
	countnum := c.GetInt("count")
	userRole := c.GetString("userRole")
	unit := c.GetString("userUnit")

	var info string
	var err error

	switch infotype {
	case "road":
		info, err = database.Info(unit, userRole, "l21_", year, countnum, model.L21{})
	case "bridge":
		info, err = database.Info(unit, userRole, "l24_", year, countnum, model.L24{})
	case "tunnel":
		info, err = database.Info(unit, userRole, "l25_", year, countnum, model.L25{})
	case "service":
		info, err = database.Info(unit, userRole, "F_", year, countnum, model.F{})
	case "portal":
		info, err = database.Info(unit, userRole, "SM_", year, countnum, model.SM{})
	case "toll":
		info, err = database.Info(unit, userRole, "SZ_", year, countnum, model.SZ{})
	default:
		err = errors.New("查询类型不存在")
	}

	if err != nil {
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, err.Error())
		return
	}

	respcode.ResponseSuccess(c, info)

	key := fmt.Sprintf("%s/%s/%s/%s",
		c.GetString("userUnit"), c.Param("infotype"),
		c.Param("year"), c.Param("count"))

	if global.CacheSetting.CacheType == "mem" ||
		global.CacheSetting.CacheType == "disk" {
		webcache.UpdataCache(key, info)
	}
}

// QueryInfo 查询数据库数据接口
// @Summary 获取查询数据
// @Description 获取数据库原始数据接口 访问后会更新缓存
// @Tags 数据操作api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param infotype path string true "查询类型 : road(路)  bridge(桥) tunnel(隧道) service(服务区) portal(收费门架) toll(收费站)"
// @Param year path string true "查询年份 格式: 202X "
// @Param level path int true "查询等级 : 0(高速) 1(一级) 2(二级) 3(三级) 4(四级) 5(等外)"
// @Param 查询信息 body model.L21 true "道路信息（示例）"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{code=int,msg=string,data=string}
// @Router /api/v1/data/info/{infotype}/{year}/{level}/query [post]
func QueryInfo(c *gin.Context) {
	infotype := c.GetString("infotype")
	countnum := c.GetInt("count")
	year := c.GetString("year")
	unit := c.GetString("userUnit")
	column := c.GetString("column")
	value := c.GetString("value")

	var info string
	var err error

	switch infotype {
	case "road":
		info, err = database.Query("l21_", year, unit, countnum, column, value, model.L21{})
	case "bridge":
		info, err = database.Query("l24_", year, unit, countnum, column, value, model.L24{})
	case "tunnel":
		info, err = database.Query("l25_", year, unit, countnum, column, value, model.L25{})
	case "service":
		info, err = database.Query("F_", year, unit, countnum, column, value, model.F{})
	case "portal":
		info, err = database.Query("SM_", year, unit, countnum, column, value, model.SM{})
	case "toll":
		info, err = database.Query("SZ_", year, unit, countnum, column, value, model.SZ{})
	default:
		err = errors.New("查询类型不存在")
	}

	if err != nil {
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, err.Error())
	}

	respcode.ResponseSuccess(c, info)
}

// UpdateInfo 根据 id 数据库数据接口
// @Summary 修改指定数据
// @Description  没完成 别调用!!! 获取数据库原始数据接口 访问后会更新缓存
// @Tags 数据操作api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param infotype path string true "查询类型 : road(路)  bridge(桥) tunnel(隧道) service(服务区) portal(收费门架) toll(收费站)"
// @Param year path string true "查询年份 格式: 202X "
// @Param level path int true "查询等级 : 0(高速) 1(一级) 2(二级) 3(三级) 4(四级) 5(等外)"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{msg=string,data=string}
// @Router /api/v1/data/info/{infotype}/{year}/{level} [post]
func UpdateInfo(c *gin.Context) {
	infotype := c.GetString("infotype")
	var err error

	switch infotype {
	case "road":
	case "bridge":
	case "tunnel":
	case "service":
	case "portal":
	case "toll":
	default:
		err = errors.New("查询类型不存在")
	}

	if err != nil {
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, err.Error())
	}

	respcode.ResponseSuccess(c, nil)

}
