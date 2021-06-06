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
	rows, err := global.DB.Query("select * from F where `ID`>2")
	if err != nil {
		log.Println(err)
		return ""
	}
	defer rows.Close()

	services := []model.F{}

	for rows.Next() {
		var f model.F
		err := rows.Scan(
			&f.ID,
			&f.F路线编号,
			&f.F路线名称,
			&f.F桩号,
			&f.F服务设施类型,
			&f.F服务设施名称,
			&f.F初始运营时间,
			&f.F布局形式,
			&f.F经营模式,
			&f.F占地面积,
			&f.F停车场面积,
			&f.F停车位数量,
			&f.F是否有厕所,
			&f.F是否有加油设施,
			&f.F是否有加气设施,
			&f.F是否有车辆充电设施,
			&f.F是否有餐饮服务设施,
			&f.F是否有超市,
			&f.F是否有住宿设施,
			&f.F是否有汽车维修,
			&f.F备注,
		)
		if err != nil {
			fmt.Println(err)
		}

		services = append(services, f)
	}
	data, err := json.Marshal(services)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func MInfo(count int) string {
	rows, err := global.DB.Query("select * from SM ")
	if err != nil {
		log.Println(err)
		return ""
	}
	defer rows.Close()

	portals := []model.SM{}

	for rows.Next() {
		var portal model.SM
		rows.Scan(
			&portal.M序号,
			&portal.M收费门架编号,
			&portal.M门架名称,
			&portal.M门架类型,
			&portal.M门架种类,
			&portal.M门架标志,
			&portal.M省界入出口标识,
			&portal.M收费单元编码组合,
			&portal.M车道数,
			&portal.M纬度,
			&portal.M经度,
			&portal.M桩号,
			&portal.M使用状态,
			&portal.M起始日期,
			&portal.M终止日期,
			&portal.M门架HEX字符串,
			&portal.M反向门架HEX字符串,
			&portal.M代收门架编号,
			&portal.M门架的RSU厂商代码,
			&portal.M门架的RSU型号,
			&portal.M门架的RSU编号,
			&portal.M门架的高清车牌识别设备厂商代码,
			&portal.M门架的高清车牌识别设备型号,
			&portal.M门架的高清车牌识别设备编号,
			&portal.M门架的高清摄像机设备厂商代码,
			&portal.M门架的高清摄像机设备型号,
			&portal.M门架的高清摄像机设备编号,
			&portal.M门架控制机设备厂商代码,
			&portal.M门架控制机设备型号,
			&portal.M门架控制机设备编号,
			&portal.M门架控制机操作系统软件版本,
			&portal.M门架服务器设备厂商代码,
			&portal.M门架服务器设备型号,
			&portal.M门架服务器设备编号,
			&portal.M门架服务器操作系统软件版本,
			&portal.M门架服务器数据库系统软件版本,
			&portal.M门架的车辆检测器设备厂商代码,
			&portal.M门架的车辆检测器设备型号,
			&portal.M门架的车辆检测器设备编号,
			&portal.M门架的车辆气象检测设备厂商代码,
			&portal.M门架的气象检测设备型号,
			&portal.M门架的气象检测设备编号,
			&portal.M门架的车型检测设备厂商代码,
			&portal.M门架的车型检测设备型号,
			&portal.M门架的车型检测设备编号,
			&portal.M门架的断面称重检测设备厂商代码,
			&portal.M门架的断面称重检测设备型号,
			&portal.M门架的断面称重检测设备编号,
			&portal.M门架的温控设备厂商代码,
			&portal.M门架的温控设备型号,
			&portal.M门架的温控设备编号,
			&portal.M门架的供电设备厂商代码,
			&portal.M门架的供电设备型号,
			&portal.M门架的供电设备编号,
			&portal.M门架的安全接入设备厂商代码,
			&portal.M门架的安全接入设备型号,
			&portal.M门架的安全接入设备编号,
			&portal.M线路类型,
			&portal.M网络所属运营商,
			&portal.M数据汇聚点,
			&portal.MIMEI号,
			&portal.M接入设备ip,
			&portal.Msnmp协议版本号,
			&portal.Msnmp端口,
			&portal.M团体名称,
			&portal.M用户名,
			&portal.M安全级别,
			&portal.M认证协议,
			&portal.M认证密钥,
			&portal.M加密算法,
			&portal.M加密密钥,
		)

		portals = append(portals, portal)
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
