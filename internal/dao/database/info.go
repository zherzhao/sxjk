package database

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"webconsole/global"
	"webconsole/internal/model"

	"github.com/impact-eintr/eorm"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

func Info(unit, role, prefix, year string, t interface{}) (string, error) {
	var res interface{}
	var flag, flag2 bool = false, false
	var condition string

	switch reflect.TypeOf(t).String() {
	case "model.L21":
		res = &[]model.L21{}
		condition = "所在行政区划代码"
		flag = true
	case "model.L24":
		res = &[]model.L24{}
		condition = "所在政区代码"
		flag = true
	case "model.L25":
		res = &[]model.L25{}
		condition = "政区代码"
		flag = true
	case "model.F":
		res = &[]model.F{}
	case "model.SM":
		res = &[]model.SM{}
		condition = "所在行政区划代码"
		flag = true
	case "model.SZ":
		res = &[]model.SZ{}
		condition = "所在地级市"
		flag2 = true
	default:
		return "", errors.New("无法找到对应数据模型")
	}

	statement := eorm.NewStatement()
	statement = statement.SetTableName(prefix + year)

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	if flag {
		for _, v := range global.AreaMap[unit] {
			statement = statement.OrEqual(condition, v)
		}
	}
	if flag2 && unit != "交科" {
		statement = statement.AndEqual(condition, unit+"市")
	}

	statement = statement.Select("*")

	err := c.FindAll(context.Background(), statement, res)
	if err != nil {
		zap.L().Error("sql exec failed: ", zap.String("", err.Error()))
		return "", ErrorNotFound
	}

	data, err := json.Marshal(res)
	if err != nil {
		zap.L().Error("marshal failed: ", zap.String("", err.Error()))
		return "", err
	}
	return string(data), nil

}
