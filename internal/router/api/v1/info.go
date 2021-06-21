package v1

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"webconsole/global"
	"webconsole/internal/dao/database"
	"webconsole/internal/dao/webcache"
	"webconsole/internal/model"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Info struct {
}

func NewInfo() Info {
	return Info{}
}

// GetInfo 获取数据库原始数据接口 访问后会更新缓存
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
func (this *Info) GetInfo(c *gin.Context) {
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
	year := c.GetString("year")
	column := c.GetString("column")
	value := c.GetString("value")

	var info string
	var err error

	switch infotype {
	case "road":
		info, err = database.Query("l21_", year, countnum, column, value, model.L21{})
	case "bridge":
		info, err = database.Query("l24_", year, countnum, column, value, model.L24{})
	case "tunnel":
	case "service":
		info, err = database.Query("F_", year, countnum, column, value, model.F{})
	case "portal":
		info, err = database.Query("SM_", year, countnum, column, value, model.SM{})
	case "toll":
		info, err = database.Query("SZ_", year, countnum, column, value, model.SZ{})
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
func (this *Info) UpdateInfo(c *gin.Context) {
	infotype := c.GetString("infotype")
	countnum := c.GetInt("count")
	year := c.GetString("year")
	column := c.GetString("column")
	value := c.GetString("value")

	var err error

	switch infotype {
	case "road":
		p := new(model.L21)

		if err := c.ShouldBindJSON(&p); err != nil {
			zap.L().Error("Update Error: ", zap.Error(err))
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				respcode.ResponseError(c, respcode.CodeInvalidParam)
				return
			}

			respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, errs.Translate(global.Trans))
			return
		}
		err = database.Update("l21_", year, countnum, p)

	case "bridge":
		_, err = database.Query("l24_", year, countnum, column, value, model.L24{})
	case "tunnel":
		_, err = database.Query("l25_", year, countnum, column, value, model.L25{})
	case "service":
		_, err = database.Query("F_", year, countnum, column, value, model.F{})
	case "portal":
		_, err = database.Query("SM_", year, countnum, column, value, model.SM{})
	case "toll":
		_, err = database.Query("SZ_", year, countnum, column, value, model.SZ{})
	default:
		err = errors.New("查询类型不存在")
	}

	if err != nil {
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, err.Error())
	}

	respcode.ResponseSuccess(c, nil)

}
