package v1

import (
	"sort"
	"strconv"
	"webconsole/global"
	"webconsole/internal/model"
	"webconsole/pkg/respcode"

	"github.com/impact-eintr/eorm"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

type Node2 struct {
	Id       string  `json:"id"`
	AuthName string  `json:"authName"`
	Query    string  `json:"query"`
	Path     string  `json:"path"`
	Children []Node2 `json:"children"`
}

// DtaMenusHandler 获取导航栏数据接口
// @Summary 获取数据导航栏
// @Description 用于前端渲染数据管理侧边栏
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{msg=string,data=string}
// @Router /api/v1/data/menus [get]
func DataMenusHandler(ctx *gin.Context) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("tableManager").
		Select("tableName,year,levelCount")

	menus := []model.Menus{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindAll(nil, statement, &menus)
	if err != nil {
		zap.L().Error("sql exec failed: ", zap.String("", err.Error()))
	}

	nodeMap := make(map[string][]Node2)
	for _, v := range menus {
		node := Node2{}
		node.Id, node.Path, node.AuthName = model.Menu(&v)
		node.Children = append(node.Children, func() (childNodes []Node2) {
			for i := 0; i < v.Levels; i++ {
				childnode := Node2{}
				childnode.Id = node.Id + strconv.Itoa(i)
				if v.Levels == 1 {
					childnode.AuthName = node.AuthName
				} else {
					childnode.AuthName, _ = model.Level(i)
				}
				childnode.Path = model.Leveltag(node.Path) + childnode.Id
				childnode.Query = strconv.Itoa(i)
				childNodes = append(childNodes, childnode)
			}
			return childNodes
		}()...)

		nodeMap[v.Year] = append(nodeMap[v.Year], node)
	}

	lists := []Node2{}
	for k, v := range nodeMap {
		node := Node2{}
		node.Id = k
		node.AuthName = k
		node.Path = k
		sort.Slice(v, func(i, j int) bool {
			return v[i].Id < v[j].Id
		})
		node.Children = append(node.Children, v...)
		lists = append(lists, node)
	}

	respcode.ResponseSuccess(ctx, lists)

}
