package database

import (
	"context"
	"fmt"
	"webconsole/global"

	"github.com/impact-eintr/eorm"
)

func DeleteTableHandler(tableName string) error {
	sqlStr := fmt.Sprintf("truncate %s", tableName)
	_, err := global.DB.Exec(sqlStr)
	if err != nil {
		return err
	}

	cli := <-global.DBClients
	defer func() {
		global.DBClients <- cli
	}()

	statement := eorm.NewStatement()
	statement = statement.SetTableName("tableManager").
		AndEqual("tableName", tableName)

	_, err = cli.Delete(context.Background(), statement)
	if err != nil {
		return err
	}
	return nil

}
