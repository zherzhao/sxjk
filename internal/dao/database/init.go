package database

import (
	"webconsole/global"

	"github.com/impact-eintr/eorm"
)

func Init() error {

	setting := eorm.Settings{
		DriverName: "mysql",
		User:       global.DatabaseSetting.User,
		Password:   global.DatabaseSetting.Password,
		Database:   global.DatabaseSetting.DBname,
		Host:       global.DatabaseSetting.Host,
		Options:    map[string]string{"charset": "utf8mb4"},
	}

	var err error
	global.DBClient, err = eorm.NewClient(setting)
	if err != nil {
		return err
	}

	return nil

}
