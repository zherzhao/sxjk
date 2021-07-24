package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"webconsole/global"
	"webconsole/internal/model"

	"github.com/impact-eintr/eorm"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

func Info(unit, role, prefix, year string, count int, t interface{}) (string, error) {
	var res interface{}
	var flag bool = true
	var condition string

	switch reflect.TypeOf(t).String() {
	case "model.L21":
		res = &[]model.L21{}
		condition = "所在行政区划代码"
	case "model.L24":
		res = &[]model.L24{}
		condition = "所在政区代码"
	case "model.L25":
		res = &[]model.L25{}
		condition = "政区代码"
	case "model.F":
		res = &[]model.F{}
		flag = false
	case "model.SM":
		res = &[]model.SM{}
		flag = false
	case "model.SZ":
		res = &[]model.SZ{}
		flag = false
	default:
		return "", errors.New("无法找到对应数据模型")
	}

	statement := eorm.NewStatement()
	if sLevel, ok := prefixMap[prefix]; ok {
		level, err := model.Level(count)
		if err != nil {
			return "", err
		}
		statement = statement.SetTableName(fmt.Sprintf(
			"(select * from %s WHERE `%s`='%s' )as res", prefix+year, sLevel, level))
	}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	if flag {
		for _, v := range global.AreaMap[unit] {
			statement = statement.OrEqual(condition, v)
		}
	}
	statement = statement.AndGreaterThan("ID", "2").Select("*")

	err := c.FindAll(nil, statement, res)
	if err != nil {
		zap.L().Error("sql exec failed: ", zap.String("", err.Error()))
		return "", err
	}

	data, err := json.Marshal(res)
	if err != nil {
		zap.L().Error("marshal failed: ", zap.String("", err.Error()))
		return "", err
	}
	return string(data), nil

}
