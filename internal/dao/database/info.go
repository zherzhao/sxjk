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

	err := global.DBClient.FindAll(nil, statement, &roads)
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

	err := global.DBClient.FindAll(nil, statement, &bridges)
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

	err := global.DBClient.FindAll(nil, statement, &tunnels)
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

	err := global.DBClient.FindAll(nil, statement, &services)
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

	err := global.DBClient.FindAll(nil, statement, &portals)
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
	rows, err := global.DB.Query("select * from SZ")
	if err != nil {
		log.Println(err)
		return ""
	}
	defer rows.Close()

	tolls := []model.SZ{}

	for rows.Next() {
		var toll model.SZ
		rows.Scan(
			&toll.S序号,
			&toll.S收费站编号,
			&toll.S收费站名称,
			&toll.S收费广场数量,
			&toll.S收费站HEX,
			&toll.S线路类型,
			&toll.S网络所属运营商,
			&toll.S数据汇聚点,
			&toll.SIMEI号,
			&toll.S接入设备ip,
			&toll.Ssnmp协议版本号,
			&toll.Ssnmp端口,
			&toll.S团体名称,
			&toll.S用户名,
			&toll.S安全级别,
			&toll.S认证协议,
			&toll.S认证密钥,
			&toll.S加密算法,
			&toll.S加密密钥,
		)
		tolls = append(tolls, toll)
	}
	data, err := json.Marshal(tolls)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}
