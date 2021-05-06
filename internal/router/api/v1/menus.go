package v1

import (
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
)

// Menus渲染节点序号
type NodeId int

const (
	User NodeId = 100 + iota
	UserList

	Right
	RightList
	RoleList

	System
	Map
	Data
)

var idMap = map[NodeId]string{
	User:     "用户管理",
	UserList: "用户列表",

	Right:     "权限管理",
	RightList: "权限列表",
	RoleList:  "角色列表",

	System: "系统管理",
	Map:    "地图管理",
	Data:   "数据管理",
}

type Node struct {
	Id       NodeId `json:"id"`
	AuthName string `json:"authName"`
	Path     string `json:"path"`
	Children []Node `json:"children"`
}

var menu = []Node{
	Node{
		Id:       User,
		AuthName: idMap[User],
		Path:     "users",
		Children: []Node{
			Node{
				Id:       UserList,
				AuthName: idMap[UserList],
				Path:     "users",
				Children: []Node{},
			},
		},
	},

	Node{
		Id:       Right,
		AuthName: idMap[Right],
		Path:     "rights",
		Children: []Node{
			Node{
				Id:       RightList,
				AuthName: idMap[RightList],
				Path:     "rights",
				Children: []Node{},
			},
			Node{
				Id:       RoleList,
				AuthName: idMap[RoleList],
				Path:     "roles",
				Children: []Node{},
			},
		},
	},

	Node{
		Id:       System,
		AuthName: idMap[System],
		Path:     "sysctl",
		Children: []Node{
			Node{
				Id:       Map,
				AuthName: idMap[Map],
				Path:     "mapctl",
				Children: []Node{},
			},
			Node{
				Id:       Data,
				AuthName: idMap[Data],
				Path:     "datactl",
				Children: []Node{},
			},
		},
	},
}

// MenusHandler 获取导航栏数据接口
// @Summary 获取导航栏数据
// @Description 用于前端渲染用户家目录侧边栏
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{msg=string,data=string}
// @Router /api/v1/home/menus [get]
func MenusHandler(c *gin.Context) {
	respcode.ResponseSuccess(c, menu)

}
