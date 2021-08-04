package model

//func Level(level int) (string, error) {
//	switch level {
//	case 0:
//		return "高速", nil
//	case 1:
//		return "一级", nil
//	case 2:
//		return "二级", nil
//	case 3:
//		return "三级", nil
//	case 4:
//		return "四级", nil
//	case 5:
//		return "等外", nil
//	default:
//		return "", errors.New("查询不匹配")
//	}
//
//}

func Menu(m *Menus) (string, string, string) {
	tablename := m.TableName[:len(m.TableName)-5]
	switch tablename {
	case "l21":
		return "0", "road", "公路信息"
	case "l24":
		return "1", "bridge", "公路桥梁"
	case "l25":
		return "2", "tunnel", "公路隧道"
	case "F":
		return "3", "service", "服务区和停车区"
	case "SZ":
		return "4", "toll", "收费站信息"
	case "SM":
		return "5", "portal", "收费门架信息"
	default:
		return "", "", ""
	}
}

func Leveltag(tag string) string {
	switch tag {
	case "road":
		return "ro"
	case "bridge":
		return "br"
	case "tunnel":
		return "tu"
	case "service":
		return "se"
	case "portal":
		return "po"
	case "toll":
		return "to"
	default:
		return ""
	}

}
