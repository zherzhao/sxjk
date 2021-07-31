package service

import (
	"webconsole/internal/dao/database"
	"webconsole/internal/model"
)

func GetUsers() ([]model.UserResp, error) {
	return database.GetUsers()

}

func QueryUsers(c string, v interface{}) ([]model.UserResp, error) {
	return database.QueryUsers(c, v)
}

func UpdateUser(reqUser *model.UserReq, id int64) (err error) {
	user := new(model.User)
	user.UserID = id
	user.Username = reqUser.Username
	if reqUser.Role == "user" && reqUser.Unit == "交科" {
		return database.ErrorInvalidUnit
	}
	user.Unit = reqUser.Unit
	user.Role = reqUser.Role
	user.Password, err = database.UserPassword(id)
	if err != nil {
		return
	}

	err = database.UpdateUsers(user)
	if err != nil {
		return
	}
	return nil

}

func DeleteUser(id int64) error {
	tmpUser := new(model.User)
	tmpUser.UserID = id

	err := database.DeleteUsersHandler(tmpUser)
	if err != nil {
		return err
	}
	return nil

}
