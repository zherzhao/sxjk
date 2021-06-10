package database

import (
	"encoding/json"
	"fmt"
	"log"
	"webconsole/global"
	"webconsole/internal/model"

	"github.com/impact-eintr/eorm"

	_ "github.com/go-sql-driver/mysql"
)

func Level(level int) string {
	switch level {
	case 0:
		return "高速"
	case 1:
		return "一级"
	case 2:
		return "二级"
	case 3:
		return "三级"
	case 4:
		return "四级"
	case 5:
		return "等外"
	}

	return ""

}

func RoadInfo(count int) string {
	level := Level(count)

	statement := eorm.NewStatement()
	statement = statement.SetTableName("l21").
		AndEqual("技术等级", level).
		AndGreaterThan("ID", "2").
		Select("*")

	roads := []model.L21{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindAll(nil, statement, &roads)
	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(roads)
	if err != nil {
		log.Println(err)
	}
	return string(data)

}

func BridgeInfo(count int) string {
	level := Level(count)

	statement := eorm.NewStatement()
	statement = statement.SetTableName("l24").
		AndEqual("技术等级", level).
		AndGreaterThan("ID", "2").
		Select("*")

	bridges := []model.L24{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindAll(nil, statement, &bridges)
	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(bridges)
	if err != nil {
		log.Println(err)
	}
	return string(data)

}

func TunnelInfo(count int) string {
	level := Level(count)

	statement := eorm.NewStatement()
	statement = statement.SetTableName("l25").
		AndEqual("所属线路技术等级", level).
		AndGreaterThan("ID", "2").
		Select("*")

	tunnels := []model.L25{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindAll(nil, statement, &tunnels)
	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(tunnels)
	if err != nil {
		log.Println(err)
	}
	return string(data)

}

func FInfo(count int) string {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("F").
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

func MInfo(count int) string {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("SM").Select("*")

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

func SInfo(count int) string {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("SZ").Select("*")

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
