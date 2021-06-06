package database

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"webconsole/global"
	"webconsole/internal/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/impact-eintr/eorm"
)

// 加密salt
const salt string = `impact-eintr`

// 错误码
var (
	ErrorUserExist       = errors.New("用户已经存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

// 检查注册时用户是否已经存在
func CheckUserExist(username string) error {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("user").
		AndEqual("username", username).
		Select("count(user_id)")

	var count int

	err := global.DBClient.FindOne(nil, statement, &count)
	if err != nil {
		log.Println(err)
		return err
	}

	if count > 0 {
		return ErrorUserExist
	}

	return nil
}

// 注册一个新的用户
func InsertUser(user *model.User) (err error) {
	// 密码加密
	user.Password = encryptPassword(user.Password)

	// 执行SQL语句入库
	statement := eorm.NewStatement()
	statement = statement.SetTableName("user").InsertStruct(user)

	_, err = global.DBClient.Insert(nil, statement)
	if err != nil {
		log.Println(err)
		return err
	}

	// 成功将用户注册后 返回用户的UID
	return nil
}

// 加密函数 (md5)
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// 用户登录
func UserLogin(user *model.User) (err error) {
	oPassword := user.Password

	statement := eorm.NewStatement()
	statement = statement.SetTableName("user").
		AndEqual("username", user.Username).
		Select("user_id, username, password")

	err = global.DBClient.FindOne(nil, statement, user)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}

	if err != nil {
		// 查询失败
		log.Println(err)
		return err
	}

	// 判断密码是否匹配
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}

	return
}
