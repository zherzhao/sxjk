package database

import (
	"context"
	"database/sql"
	"log"
	"webconsole/global"
	"webconsole/internal/model"

	"github.com/impact-eintr/eorm"
)

func GetUsersHandler() (users []model.RespUser, err error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("user").
		Select("user_id, username, role, unit")

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err = c.FindAll(context.Background(), statement, &users)
	if err == sql.ErrNoRows {
		return nil, ErrorUserNotExist
	}

	if err != nil {
		// 查询失败
		log.Println(err)
		return
	}
	return

}
