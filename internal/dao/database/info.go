package database

import (
	"encoding/json"
	"errors"
	"reflect"
	"webconsole/global"
	"webconsole/internal/model"

	"github.com/impact-eintr/eorm"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

func Info(unit, role, prefix, year string, count int, t interface{}) (string, error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName(prefix + year)
	if sLevel, ok := prefixMap[prefix]; ok {
		level, err := model.Level(count)
		if err != nil {
			return "", err
		}
		statement = statement.AndEqual(sLevel, level).AndGreaterThan("ID", "2")
	}

	if role == "user" {
		statement = statement.AndLike("管养单位名称", unit+"%")
	}
	statement = statement.Select("*")

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	var res interface{}
	switch reflect.TypeOf(t).String() {
	case "model.L21":
		res = &[]model.L21{}
	case "model.L24":
		res = &[]model.L24{}
	case "model.L25":
		res = &[]model.L25{}
	case "model.F":
		res = &[]model.F{}
	case "model.SM":
		res = &[]model.SM{}
	case "model.SZ":
		res = &[]model.SZ{}
	default:
		return "", errors.New("无法找到对应数据模型")
	}

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
