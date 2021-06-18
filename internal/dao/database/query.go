package database

import (
	"encoding/json"
	"reflect"
	"webconsole/global"
	"webconsole/internal/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/impact-eintr/eorm"
	"go.uber.org/zap"
)

var (
	prefixMap = map[string]string{
		"l21_": "技术等级",
		"l24_": "技术等级",
		"l25_": "所属线路技术等级",
	}
)

func Query(prefix, year string, count int, column, value string, t interface{}) (string, error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName(prefix + year)
	if sLevel, ok := prefixMap[prefix]; ok {
		level, err := model.Level(count)
		if err != nil {
			return "", err
		}
		statement = statement.AndEqual(sLevel, level)
	}

	switch column {
	case "id":
		statement = statement.AndEqual(column, value).Select("*")
	default:
		statement = statement.AndLike(column, "%"+value).Select("*")
	}

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
		return "", nil
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
