package database

import (
	"encoding/json"
	"fmt"
	"log"
	"webconsole/global"
	"webconsole/internal/model"

	_ "github.com/go-sql-driver/mysql"
)

func RoadQuery(count int, column string, value string) string {
	level, _ := Level(count)

	sqlStr := fmt.Sprintf("select * from l21 where `技术等级`=? AND `id`>2 AND `%s` LIKE ?", column)
	// 查询后调用Scan 否则持有的数据库连接不会被释放
	rows, err := global.DB.Query(sqlStr, level, value)
	if err != nil {
		log.Println(err)
		return ""
	}

	// 关闭rows释放持有的数据库连接
	defer rows.Close()

	roads := []model.L21{}

	for rows.Next() {
		var road model.L21
		rows.Scan(
			&road.ID,
			&road.R路线编号,
			&road.R所在行政区划代码,
			&road.R路线名称,
			&road.R起点名称,
			&road.R止点名称,
			&road.R起点桩号,
			&road.R止点桩号,
			&road.R里程公里,
			&road.R技术等级代码,
			&road.R技术等级,
			&road.R是否一幅高速,
			&road.R车道数量个,
			&road.R面层类型代码,
			&road.R面层类型,
			&road.R路基宽度米,
			&road.R路面宽度米,
			&road.R面层厚度厘米,
			&road.R设计时速公里小时,
			&road.R修建年度,
			&road.R改建年度,
			&road.R最近一次修复养护年度,
			&road.R断链类型,
			&road.R是否城管路段,
			&road.R是否断头路段,
			&road.R路段收费性质,
			&road.R重复路段线路编号,
			&road.R重复路段起点桩号,
			&road.R重复路段终点桩号,
			&road.R养护里程,
			&road.R可绿化里程,
			&road.R已绿化里程,
			&road.R地貌代码,
			&road.R地貌汉字,
			&road.R涵洞数量个,
			&road.R管养单位名称,
			&road.R省际出入口,
			&road.R国道调整前路线编号,
			&road.R是否按干线公路管理接养,
			&road.R备注,
		)

		roads = append(roads, road)
	}

	data, err := json.Marshal(roads)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(roads)
	return string(data)
}

func BridgeQuery(count int, column string, value string) string {
	level, _ := Level(count)

	sqlStr := fmt.Sprintf("select * from l24 where `技术等级`=? AND `id`>2 AND `%s` LIKE ?", column)
	// 查询后调用Scan 否则持有的数据库连接不会被释放
	rows, err := global.DB.Query(sqlStr, level, value)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer rows.Close()

	bridges := []model.L24{}

	for rows.Next() {
		var bridge model.L24
		rows.Scan(
			&bridge.ID,
			&bridge.Q桥梁名称,
			&bridge.Q桥梁代码,
			&bridge.Q桥梁中心桩号,
			&bridge.Q路线编号,
			&bridge.Q路线名称,
			&bridge.Q技术等级,
			&bridge.Q桥梁全长米,
			&bridge.Q跨径总长米,
			&bridge.Q单孔最大跨径米,
			&bridge.Q跨径组合孔米,
			&bridge.Q桥梁全宽米,
			&bridge.Q桥面净宽米,
			&bridge.Q按跨径分类代码,
			&bridge.Q按跨径分类类型,
			&bridge.Q按使用年限分类代码,
			&bridge.Q按使用年限分类类型,
			&bridge.Q主桥上部构造结构类型代码,
			&bridge.Q主桥上部构造结构类型类型,
			&bridge.Q主桥上部构造材料代码,
			&bridge.Q主桥上部构造材料名称,
			&bridge.Q桥墩类型代码,
			&bridge.Q桥墩类型类型,
			&bridge.Q设计荷载等级代码,
			&bridge.Q设计荷载等级,
			&bridge.Q抗震等级代码,
			&bridge.Q抗震等级,
			&bridge.Q跨越地物代码,
			&bridge.Q跨越地物类型,
			&bridge.Q跨越地物名称,
			&bridge.Q通航等级,
			&bridge.Q墩台防撞设施类型,
			&bridge.Q是否互通立交,
			&bridge.Q建设单位,
			&bridge.Q设计单位,
			&bridge.Q施工单位,
			&bridge.Q监理单位,
			&bridge.Q修建年度,
			&bridge.Q建成通车日期,
			&bridge.Q管养单位性质代码,
			&bridge.Q管养单位名称,
			&bridge.Q监管单位名称,
			&bridge.Q收费性质代码,
			&bridge.Q收费性质,
			&bridge.Q评定等级代码,
			&bridge.Q评定等级,
			&bridge.Q评定日期,
			&bridge.Q评定单位,
			&bridge.Q最后一次改造年度,
			&bridge.Q最后一次改造完工日期,
			&bridge.Q最后一次改造部位,
			&bridge.Q最后一次改造工程性质,
			&bridge.Q最后一次改造施工单位,
			&bridge.Q最后一次改造是否部补助项目,
			&bridge.Q当前主要病害部位,
			&bridge.Q当前主要病害,
			&bridge.Q交通管制措施代码,
			&bridge.Q交通管制措施,
			&bridge.Q所在政区代码,
			&bridge.Q桥梁所在位置,
			&bridge.Q是否宽路窄桥,
			&bridge.Q是否在长大桥梁目录中,
			&bridge.Q备注,
		)

		bridges = append(bridges, bridge)
	}

	data, err := json.Marshal(bridges)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func TunnelQuery(count int, column string, value string) string {
	level, _ := Level(count)

	sqlStr := fmt.Sprintf("select * from l25 where `所属线路技术等级`=? AND `id`>2 AND `%s` LIKE ?", column)
	// 查询后调用Scan 否则持有的数据库连接不会被释放
	rows, err := global.DB.Query(sqlStr, level, value)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer rows.Close()

	tunnels := []model.L25{}

	for rows.Next() {
		var tunnel model.L25
		rows.Scan(
			&tunnel.ID,
			&tunnel.S隧道名称,
			&tunnel.S隧道代码,
			&tunnel.S隧道中心桩号,
			&tunnel.S所属路线编号,
			&tunnel.S所属路线名称,
			&tunnel.S所属线路技术等级,
			&tunnel.S隧道长度米,
			&tunnel.S隧道净宽米,
			&tunnel.S隧道净高米,
			&tunnel.S隧道按长度分类代码,
			&tunnel.S隧道按长度分类,
			&tunnel.S是否水下隧道,
			&tunnel.S修建年度,
			&tunnel.S建设单位名称,
			&tunnel.S设计单位名称,
			&tunnel.S施工单位名称,
			&tunnel.S监理单位名称,
			&tunnel.S建成通车时间,
			&tunnel.S隧道养护等级,
			&tunnel.S管养单位性质代码,
			&tunnel.S管养名称,
			&tunnel.S监管单位名称,
			&tunnel.S技术状况评定等级,
			&tunnel.S技术状况评定日期,
			&tunnel.S技术状况评定单位,
			&tunnel.S土建结构评定等级,
			&tunnel.S土建结构评定日期,
			&tunnel.S土建结构评定单位,
			&tunnel.S机电设施评定等级,
			&tunnel.S机电设施评定日期,
			&tunnel.S机电设施评定单位,
			&tunnel.S其他工程设施评定等级,
			&tunnel.S其他工程设施评定日期,
			&tunnel.S其他工程设施评定单位,
			&tunnel.S最近一次改造年度,
			&tunnel.S最近一次改造完工日期,
			&tunnel.S最近一次改造部位,
			&tunnel.S最近一次改造工程性质,
			&tunnel.S当前主要病害部位,
			&tunnel.S当前主要病害描述,
			&tunnel.S政区代码,
			&tunnel.S是否在长大隧道目录中,
			&tunnel.S备注,
		)

		tunnels = append(tunnels, tunnel)
	}
	data, err := json.Marshal(tunnels)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func FQuery(column string, value string) string {
	sqlStr := fmt.Sprintf("select * from F where `%s` LIKE ?", column)
	// 查询后调用Scan 否则持有的数据库连接不会被释放
	rows, err := global.DB.Query(sqlStr, value)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer rows.Close()

	services := []model.F{}

	for rows.Next() {
		var f model.F
		rows.Scan(
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

		services = append(services, f)
	}
	data, err := json.Marshal(services)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func MQuery(column string, value string) string {
	sqlStr := fmt.Sprintf("select * from SM where %s` LIKE ?", column)
	// 查询后调用Scan 否则持有的数据库连接不会被释放
	rows, err := global.DB.Query(sqlStr, value)
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

func SQuery(column string, value string) string {
	sqlStr := fmt.Sprintf("select * from SZ where %s` LIKE ?", column)
	// 查询后调用Scan 否则持有的数据库连接不会被释放
	rows, err := global.DB.Query(sqlStr, value)
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
		)
		tolls = append(tolls, toll)
	}
	data, err := json.Marshal(tolls)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}
