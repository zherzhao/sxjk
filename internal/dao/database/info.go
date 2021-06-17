package database

import (
	"encoding/json"
	"fmt"
	"log"
	"webconsole/global"
	"webconsole/internal/model"

	"github.com/impact-eintr/eorm"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

func RoadInfo(year string, count int) (string, error) {
	level, err := model.Level(count)
	if err != nil {
		return "", err
	}

	statement := eorm.NewStatement()
	statement = statement.SetTableName("l21_"+year).
		AndEqual("技术等级", level).
		AndGreaterThan("ID", "2").
		Select("*")

	roads := []model.L21{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err = c.FindAll(nil, statement, &roads)
	if err != nil {
		zap.L().Error("sql exec failed: ", zap.String("", err.Error()))
		return "", err
	}

	data, err := json.Marshal(roads)
	if err != nil {
		zap.L().Error("marshal failed: ", zap.String("", err.Error()))
		return "", err
	}
	return string(data), nil

}

func BridgeInfo(year string, count int) (string, error) {
	level, err := model.Level(count)
	if err != nil {
		return "", err
	}

	statement := eorm.NewStatement()
	statement = statement.SetTableName("l24_"+year).
		AndEqual("技术等级", level).
		AndGreaterThan("ID", "2").
		Select("*")

	bridges := []model.L24{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err = c.FindAll(nil, statement, &bridges)
	if err != nil {
		zap.L().Error("sql exec failed: ", zap.String("", err.Error()))
		return "", err
	}

	data, err := json.Marshal(bridges)
	if err != nil {
		zap.L().Error("marshal failed: ", zap.String("", err.Error()))
		return "", err
	}
	return string(data), nil

}

func TunnelInfo(year string, count int) (string, error) {
	level, err := model.Level(count)
	if err != nil {
		return "", err
	}

	statement := eorm.NewStatement()
	statement = statement.SetTableName("l25_"+year).
		AndEqual("所属线路技术等级", level).
		AndGreaterThan("ID", "2").
		Select("*")

	tunnels := []model.L25{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err = c.FindAll(nil, statement, &tunnels)
	if err != nil {
		zap.L().Error("sql exec failed: ", zap.String("", err.Error()))
		return "", err
	}

	data, err := json.Marshal(tunnels)
	if err != nil {
		zap.L().Error("marshal failed: ", zap.String("", err.Error()))
		return "", err
	}
	return string(data), nil

}

func FInfo(year string) string {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("F_"+year).
		AndGreaterThan("ID", "2").
		Select("*")

	services := []model.F{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindAll(nil, statement, &services)
	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(services)
	if err != nil {
		log.Println(err)
	}
	return string(data)

}

func MInfo(year string) string {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("SM_" + year).Select("*")

	portals := []model.SM{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindAll(nil, statement, &portals)
	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(portals)
	if err != nil {
		log.Println(err)
	}
	return string(data)

}

func SInfo(year string) string {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("SZ_" + year).Select("*")

	tolls := []model.SZ{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindAll(nil, statement, &tolls)
	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(tolls)
	if err != nil {
		log.Println(err)
	}
	return string(data)

}
