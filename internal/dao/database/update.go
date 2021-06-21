package database

import (
	"context"
	"errors"
	"reflect"
	"webconsole/global"
	"webconsole/internal/model"

	"github.com/impact-eintr/eorm"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

func Update(prefix, year string, count int, t interface{}) error {
	statement := eorm.NewStatement()
	statement = statement.SetTableName(prefix + year)
	if sLevel, ok := prefixMap[prefix]; ok {
		level, err := model.Level(count)
		if err != nil {
			return err
		}
		statement = statement.AndEqual(sLevel, level).AndGreaterThan("ID", "2")
	}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	var res interface{}
	switch reflect.TypeOf(t).String() {
	case "model.L21":
		res = &model.L21{}
	case "model.L24":
		res = &model.L24{}
	case "model.L25":
		res = &model.L25{}
	case "model.F":
		res = &model.F{}
	case "model.SM":
		res = &model.SM{}
	case "model.SZ":
		res = &model.SZ{}
	default:
		return errors.New("无法找到对应数据模型")
	}

	statement = statement.UpdateStruct(res)
	_, err := c.Update(context.Background(), statement)
	if err != nil {
		zap.L().Error("sql exec failed: ", zap.String("", err.Error()))
		return err
	}

	return nil

}
