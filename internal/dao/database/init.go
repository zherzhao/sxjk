package database

import (
	"database/sql"
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
	global.DB, err = sql.Open(setting.DriverName, setting.DataSourceName())
	if err != nil {
		return err
	}

	global.DBClients = make(chan *eorm.Client, 50)
	for i := 0; i < 50; i++ {
		c, err := eorm.NewClientWithDBconn(global.DB)
		if err != nil {
			return err
		}
		global.DBClients <- c
	}
	return nil

}
