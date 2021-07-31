package v1

import (
	"errors"
	"fmt"
	"log"
	"webconsole/global"
	"webconsole/internal/dao/database"
	"webconsole/internal/model"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func ParseTable(c *gin.Context) {
	var tabletype, year string
	var inStructPtr interface{}
	var tableName string
	var err error

	tabletype = c.Param("tabletype")
	year = c.Param("year")

	fp := c.GetString("tableName")
	f, err := excelize.OpenFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		// 只能上传.xlsx文件，并且只能有一个sheet，名为"Sheet1"
		respcode.ResponseSuccess(c, `文件上传成功，解析失败，只能上传.xlsx文件，并且只能有一个sheet，名为"Sheet1"`)
	}

	cli := <-global.DBClients
	defer func() {
		global.DBClients <- cli
	}()
	// 开始检验数据
	switch tabletype {
	case "road":
		inStructPtr = new(model.L21)
		tableName = "l21_" + year
		err = database.InsertTableHandler(tableName, year, 6, inStructPtr, &rows)
	case "bridge":
		inStructPtr = new(model.L24)
		tableName = "l24_" + year
		err = database.InsertTableHandler(tableName, year, 5, inStructPtr, &rows)
	case "tunnel":
		inStructPtr = new(model.L25)
		tableName = "l25_" + year
		err = database.InsertTableHandler(tableName, year, 6, inStructPtr, &rows)
	case "service":
		inStructPtr = new(model.F)
		tableName = "F_" + year
		err = database.InsertTableHandler(tableName, year, 1, inStructPtr, &rows)
	case "portal":
		inStructPtr = new(model.SM)
		tableName = "SM_" + year
		err = database.InsertTableHandler(tableName, year, 1, inStructPtr, &rows)
	case "toll":
		inStructPtr = new(model.SZ)
		tableName = "SZ_" + year
		err = database.InsertTableHandler(tableName, year, 1, inStructPtr, &rows)
	default:
		err = errors.New("未知类型")
	}

	if err != nil {
		log.Println(err)
		return
	}

	respcode.ResponseSuccess(c, nil)

}

func DeleteTable(c *gin.Context) {

	var tabletype, year string
	var tableName string
	var err error

	tabletype = c.Param("tabletype")
	year = c.Param("year")
	switch tabletype {
	case "road":
		tableName = "l21_" + year
	case "bridge":
		tableName = "l24_" + year
	case "tunnel":
		tableName = "l25_" + year
	case "service":
		tableName = "F_" + year
	case "portal":
		tableName = "SM_" + year
	case "toll":
		tableName = "SZ_" + year
	default:
		err = errors.New("未知类型")
		return
	}

	err = database.DeleteTableHandler(tableName)
	if err != nil {
		log.Println(err)
		return
	}

	respcode.ResponseSuccess(c, nil)

}
