package database

import (
	"context"
	"fmt"
	"reflect"
	"webconsole/global"
	"webconsole/internal/model"

	"github.com/impact-eintr/eorm"
)

func InsertTableHandler(tableName, year string, inStructPtr interface{}, data *[][]string) (err error) {
	var count int = 1

	cli := <-global.DBClients
	defer func() {
		global.DBClients <- cli
	}()

	menu := new(model.Menus)
	menu.TableName = tableName
	menu.Year = year

	for _, row := range (*data)[1:] {
		rType := reflect.TypeOf(inStructPtr).Elem()
		rVal := reflect.ValueOf(inStructPtr).Elem()
		for i := 0; i < rType.NumField(); i++ {
			f := rVal.Field(i)
			for len(row) < rType.NumField() {
				row = append(row, "")
			}
			if i == 0 {
				v := fmt.Sprintf("%d", count)
				f.Set(reflect.ValueOf(v))
			} else {
				v := row[i-1]
				f.Set(reflect.ValueOf(v))
			}
		}
		count++
		statement := eorm.NewStatement()
		statement = statement.SetTableName(tableName).InsertStruct(inStructPtr)

		_, err = cli.Insert(context.Background(), statement)
		if err != nil {
			return
		}

	}

	statement := eorm.NewStatement()
	statement = statement.SetTableName("tableManager").
		InsertStruct(menu)

	_, err = cli.Insert(context.Background(), statement)
	if err != nil {
		return
	}
	return

}
