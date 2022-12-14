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
	validator "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
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
	userRole := c.GetString("userRole")
	unit := c.GetString("userUnit")

	var info string
	var err error

	switch infotype {
	case "road":
		info, err = database.Info(unit, userRole, "l21_", year, model.L21{})
	case "bridge":
		info, err = database.Info(unit, userRole, "l24_", year, model.L24{})
	case "tunnel":
		info, err = database.Info(unit, userRole, "l25_", year, model.L25{})
	case "service":
		info, err = database.Info(unit, userRole, "F_", year, model.F{})
	case "portal":
		info, err = database.Info(unit, userRole, "SM_", year, model.SM{})
	case "toll":
		info, err = database.Info(unit, userRole, "SZ_", year, model.SZ{})
	default:
		err = errors.New("查询类型不存在")
	}

	if err != nil {
		if err == database.ErrorNotFound {
			respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, err.Error())
			return
		}
		respcode.ResponseError(c, respcode.CodeServerBusy)
		return
	}

	respcode.ResponseSuccess(c, info)

	key := fmt.Sprintf("%s/%s/%s",
		c.GetString("userUnit"), c.Param("infotype"), c.Param("year"))

	webcache.CacheUpdate(key, info)
}

// QueryInfo 查询数据库数据接口
// @Summary 获取查询数据
// @Description 查询表数据（年报）数据接口
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
	year := c.GetString("year")
	unit := c.GetString("userUnit")

	json := make(map[string][]string) //注意该结构接受的内容
	if err := c.BindJSON(&json); err != nil {
		zap.L().Error("序列化失败", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			respcode.ResponseError(c, respcode.CodeInvalidParam)
		} else {
			respcode.ResponseErrorWithMsg(c, respcode.CodeInvalidParam, errs.Translate(global.Trans))
		}
		c.Abort()
		return
	}

	query := make(map[string]string, len(json["column"]))
	for i, v := range json["column"] {
		query[v] = json["value"][i]
	}

	var info string
	var err error

	switch infotype {
	case "road":
		info, err = database.Query("l21_", year, unit, query, model.L21{})
	case "bridge":
		info, err = database.Query("l24_", year, unit, query, model.L24{})
	case "tunnel":
		info, err = database.Query("l25_", year, unit, query, model.L25{})
	case "service":
		info, err = database.Query("F_", year, unit, query, model.F{})
	case "portal":
		info, err = database.Query("SM_", year, unit, query, model.SM{})
	case "toll":
		info, err = database.Query("SZ_", year, unit, query, model.SZ{})
	default:
		err = errors.New("查询类型不存在")
	}

	if err != nil {
		if err == database.ErrorNotFound {
			respcode.ResponseError(c, respcode.CodeNotFoundData)
			return
		}
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, err.Error())
		return
	}

	respcode.ResponseSuccess(c, info)
}

// UpdateInfo 根据 id 数据库数据接口
// @Summary 修改指定数据
// @Description  更新数据库原始数据接口 访问后会删除缓存
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
	year := c.GetString("year")
	unit := c.GetString("userUnit")
	level := c.GetInt("count")
	var err error

	switch infotype {
	case "road":
		t := new(model.L21)
		c.ShouldBindJSON(t)
		database.UpdateRecordHandler("l21_", year, unit, level, t)
	case "bridge":
		t := new(model.L24)
		c.ShouldBindJSON(t)
		database.UpdateRecordHandler("l24_", year, unit, level, t)
	case "tunnel":
		t := new(model.L25)
		c.ShouldBindJSON(t)
		database.UpdateRecordHandler("l25_", year, unit, level, t)
	case "service":
		t := new(model.F)
		c.ShouldBindJSON(t)
		database.UpdateRecordHandler("F_", year, unit, level, t)
	case "portal":
		t := new(model.SM)
		c.ShouldBindJSON(t)
		database.UpdateRecordHandler("SM_", year, unit, level, t)
	case "toll":
		t := new(model.SZ)
		c.ShouldBindJSON(t)
		database.UpdateRecordHandler("SZ_", year, unit, level, t)
	default:
		err = errors.New("查询类型不存在")
	}

	if err != nil {
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, err.Error())
	}

	respcode.ResponseSuccess(c, nil)

}

// DeleteInfo 根据 id 删除数据库数据接口
// @Summary 删除指定数据
// @Description  删除数据库原始数据接口 访问后会删除缓存
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
func DeleteInfo(c *gin.Context) {
	var err error
	infotype := c.GetString("infotype")
	year := c.GetString("year")
	unit := c.GetString("userUnit")
	level := c.GetInt("count")
	id := c.Param("id")

	switch infotype {
	case "road":
		t := new(model.L21)
		t.ID = id
		err = database.DeleteRecordHandler("l21_", year, unit, level, t)
	case "bridge":
		t := new(model.L24)
		t.ID = id
		err = database.DeleteRecordHandler("l24_", year, unit, level, t)
	case "tunnel":
		t := new(model.L25)
		t.ID = id
		err = database.DeleteRecordHandler("l25_", year, unit, level, t)
	case "service":
		t := new(model.F)
		t.ID = id
		err = database.DeleteRecordHandler("F_", year, unit, level, t)
	case "portal":
		t := new(model.SM)
		t.ID = id
		err = database.DeleteRecordHandler("SM_", year, unit, level, t)
	case "toll":
		t := new(model.SZ)
		t.ID = id
		err = database.DeleteRecordHandler("SZ_", year, unit, level, t)
	default:
		err = errors.New("查询类型不存在")
	}

	if err != nil {
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, err.Error())
	}

	respcode.ResponseSuccess(c, nil)

}
