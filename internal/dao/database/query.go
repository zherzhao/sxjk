package database

import (
	"context"
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

func Query(prefix, year, unit string, query map[string]string, t interface{}) (string, error) {
	var condition string
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
	if unit == "交科" || condition == "" {
		statement = statement.SetTableName(prefix + year)
		for k, v := range query {
			statement = statement.AndLike(k, "%"+v+"%")
		}
		statement = statement.Select("*")
	} else {
		// 构建多级查询
		sql := "(select * from " + prefix + year + " WHERE "
		count := 1

		for k, v := range query {
			if count != len(query) {
				sql += fmt.Sprintf("`%s` LIKE '%s' AND ", k, "%"+v+"%")
			} else {
				sql += fmt.Sprintf("`%s` LIKE '%s'", k, "%"+v+"%")
			}
			count++
		}
		sql += ") AS res "
		statement = statement.SetTableName(sql)
		for _, v := range global.AreaMap[unit] {
			statement = statement.OrEqual(condition, v)
		}
		statement = statement.Select("*")

	}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err = c.FindAll(context.Background(), statement, res)
	if err != nil {
		zap.L().Error("sql exec failed: ", zap.String("", err.Error()))
		return "", err
	}

	if size(res) == 0 {
		return "", ErrorNotFound
	}

	data, err := json.Marshal(res)
	if err != nil {
		zap.L().Error("marshal failed: ", zap.String("", err.Error()))
		return "", err
	}
	return string(data), nil

}

func size(a interface{}) int {
	if reflect.TypeOf(a).Elem().Kind() != reflect.Slice {
		return -1
	}
	ins := reflect.ValueOf(a).Elem()
	return ins.Len()
}
