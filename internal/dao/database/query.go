package database

import (
	"encoding/json"
	"webconsole/global"
	"webconsole/internal/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/impact-eintr/eorm"
	"go.uber.org/zap"
)

func RoadQuery(year string, count int, column string, value string) (string, error) {
	level, err := model.Level(count)
	if err != nil {
		return "", err
	}

	statement := eorm.NewStatement()
	statement = statement.SetTableName("l21_"+year).
		AndEqual("技术等级", level).
		AndGreaterThan("ID", "2").
		AndLike(column, "%"+value).
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

func BridgeQuery(year string, count int, column string, value string) (string, error) {
	level, err := model.Level(count)
	if err != nil {
		return "", err
	}

	statement := eorm.NewStatement()
	statement = statement.SetTableName("l24_"+year).
		AndEqual("技术等级", level).
		AndGreaterThan("ID", "2").
		AndLike(column, "%"+value).
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

func TunnelQuery(year string, count int, column string, value string) (string, error) {
	level, err := model.Level(count)
	if err != nil {
		return "", err
	}

	statement := eorm.NewStatement()
	statement = statement.SetTableName("l25_"+year).
		AndEqual("所属线路技术等级", level).
		AndGreaterThan("ID", "2").
		AndLike(column, "%"+value).
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

func FQuery(year string, column string, value string) (string, error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("F_"+year).
		AndGreaterThan("ID", "2").
		AndLike(column, "%"+value).
		Select("*")

	services := []model.F{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindAll(nil, statement, &services)
	if err != nil {
		zap.L().Error("sql exec failed: ", zap.String("", err.Error()))
		return "", err
	}

	data, err := json.Marshal(services)
	if err != nil {
		zap.L().Error("marshal failed: ", zap.String("", err.Error()))
		return "", err
	}
	return string(data), nil

}

func MQuery(year, column, value string) (string, error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("SM_"+year).
		AndLike(column, "%"+value).
		Select("*")

	portals := []model.SM{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindAll(nil, statement, &portals)
	if err != nil {
		zap.L().Error("sql exec failed: ", zap.String("", err.Error()))
		return "", err
	}

	data, err := json.Marshal(portals)
	if err != nil {
		zap.L().Error("marshal failed: ", zap.String("", err.Error()))
		return "", err
	}

	return string(data), nil

}

func SQuery(year, column, value string) (string, error) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("SZ_"+year).
		AndLike(column, "%"+value).
		Select("*")

	portals := []model.SZ{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindAll(nil, statement, &portals)
	if err != nil {
		zap.L().Error("sql exec failed: ", zap.String("", err.Error()))
		return "", err
	}

	data, err := json.Marshal(portals)
	if err != nil {
		zap.L().Error("marshal failed: ", zap.String("", err.Error()))
		return "", err
	}
	return string(data), nil

}
