package database

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"reflect"
	"webconsole/global"
	"webconsole/internal/model"

	"github.com/impact-eintr/eorm"
)

func DeleteRecordHandler(prefix, year, unit string, count int, t interface{}) error {
	var err error
	statement := eorm.NewStatement()
	statement = statement.SetTableName(prefix + year)

	switch reflect.TypeOf(t).String() {
	case "*model.L21":
		log.Println(t.(*model.L21))
		statement = statement.AndEqual("id", t.(*model.L21).ID)
	case "*model.L24":
		log.Println(t.(*model.L24))
		statement = statement.AndEqual("id", t.(*model.L24).ID)
	case "*model.L25":
		log.Println(t.(*model.L25))
		statement = statement.AndEqual("id", t.(*model.L25).ID)
	case "*model.F":
		log.Println(t.(*model.F))
		statement = statement.AndEqual("id", t.(*model.F).ID)
	case "*model.SM":
		log.Println(t.(*model.SM))
		statement = statement.AndEqual("id", t.(*model.SM).ID)
	case "*model.SZ":
		log.Println(t.(*model.SZ))
		statement = statement.AndEqual("id", t.(*model.SZ).ID)
	default:
		return errors.New("无法找到对应数据模型")
	}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()
	_, err = c.Delete(context.Background(), statement)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}

	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}
