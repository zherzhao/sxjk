package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"webconsole/global"
	"webconsole/internal/model"

	"github.com/impact-eintr/eorm"
)

func GetUsersHandler() (users []model.RespUser, err error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("user").
		Select("user_id, username, username, role, unit")

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
	for idx := range users {
		users[idx].UserIDStr = fmt.Sprintf("%d", users[idx].UserID)
	}
	return

}

var userMap = map[string]string{
	"用户名":  "username",
	"用户id": "user_id",
	"单位":   "unit",
	"角色":   "role",
}

func QueryUsersHandler(column, value string) (users []model.RespUser, err error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("user")
	switch column {
	case "user_id_str":
		id, _ := strconv.ParseInt(value, 10, 64)
		statement = statement.AndEqual("user_id", id)
	default:
		statement = statement.AndLike(column, value)
	}
	statement = statement.Select("user_id, username, username, role, unit")

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
	for idx := range users {
		users[idx].UserIDStr = fmt.Sprintf("%d", users[idx].UserID)
	}
	return

}

func UserPassword(id int64) (string, error) {
	tmp := new(model.User)
	statement1 := eorm.NewStatement()
	statement1 = statement1.SetTableName("user").
		AndEqual("user_id", id).
		Select("user_id, username, password, role, unit")

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindOne(context.Background(), statement1, tmp)
	// 查询没有结果
	if err == sql.ErrNoRows {
		return "", ErrorUserNotExist
	}

	// 查询失败
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tmp.Password, nil

}

func UpdateUsersHandler(users *model.User) (err error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("user").
		AndEqual("user_id", users.UserID).UpdateStruct(users)

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	_, err = c.Update(context.Background(), statement)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}

	if err != nil {
		log.Println(err)
		return
	}
	return

}

func DeleteUsersHandler(users *model.User) (err error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("user").
		AndEqual("user_id", users.UserID)

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
		return
	}
	return

}
