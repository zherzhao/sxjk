package database

import (
	"encoding/json"
	"errors"
	"fmt"
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

func Query(prefix, year, unit string, count int, column, value string, t interface{}) (string, error) {
	var condition, sLevel, level string
	var ok bool
	var err error
	var res interface{}

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
	case "model.SM":
		res = &[]model.SM{}
	case "model.SZ":
		res = &[]model.SZ{}
	default:
		return "", errors.New("无法找到对应数据模型")
	}

	statement := eorm.NewStatement()
	if sLevel, ok = prefixMap[prefix]; ok {
		level, err = model.Level(count)
		if err != nil {
			return "", err
		}
		if unit != "" {
			switch column {
			case "id":
				statement = statement.SetTableName(prefix + year)
				statement = statement.AndEqual(column, value).Select("*")
			default:
				statement = statement.SetTableName(fmt.Sprintf(
					"(select * from %s WHERE `%s`='%s' AND `%s` LIKE '%s')as res",
					prefix+year, sLevel, level, column, "%"+value+"%"))
				for _, v := range global.AreaMap[unit] {
					statement = statement.OrEqual(condition, v)
				}
				statement = statement.Select("*")
			}
		} else {
			statement = statement.SetTableName(prefix + year)
			statement = statement.AndEqual(sLevel, level)

			switch column {
			case "id":
				statement = statement.AndEqual(column, value).Select("*")
			default:
				statement = statement.AndLike(column, "%"+value+"%").Select("*")
			}
		}
	}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err = c.FindAll(nil, statement, res)
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
