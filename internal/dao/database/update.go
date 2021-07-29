package database

import (
	"context"
	"errors"
	"log"
	"reflect"
	"webconsole/global"
	"webconsole/internal/model"

	"github.com/impact-eintr/eorm"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

func UpdateRecordHandler(prefix, year, unit string, count int, t interface{}) error {
	//var condition, sLevel, level string
	var err error
	statement := eorm.NewStatement()
	statement = statement.SetTableName(prefix + year)

	switch reflect.TypeOf(t).String() {
	case "*model.L21":
		log.Println(t.(*model.L21))
		statement = statement.AndEqual("id", t.(*model.L21).ID).
			UpdateStruct(t.(*model.L21))
	case "*model.L24":
		log.Println(t.(*model.L24))
		statement = statement.AndEqual("id", t.(*model.L24).ID).
			UpdateStruct(t.(*model.L24))
	case "*model.L25":
		log.Println(t.(*model.L25))
		statement = statement.AndEqual("id", t.(*model.L25).ID).
			UpdateStruct(t.(*model.L25))
	case "*model.F":
		log.Println(t.(*model.F))
		statement = statement.AndEqual("id", t.(*model.F).ID).
			UpdateStruct(t.(*model.F))
	case "*model.SM":
		log.Println(t.(*model.SM))
		statement = statement.AndEqual("id", t.(*model.SM).ID).
			UpdateStruct(t.(*model.SM))
	case "*model.SZ":
		log.Println(t.(*model.SZ))
		statement = statement.AndEqual("id", t.(*model.SZ).ID).
			UpdateStruct(t.(*model.SZ))
	default:
		return errors.New("无法找到对应数据模型")
	}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	_, err = c.Update(context.Background(), statement)
	if err != nil {
		zap.L().Error("sql exec failed: ", zap.String("", err.Error()))
		return err
	}
	return nil

}
