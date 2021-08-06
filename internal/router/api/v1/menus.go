package v1

import (
	"sort"
	"strconv"
	"sync"
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

var L sync.RWMutex

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

var rootMenu = []Node{
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
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{msg=string,data=string}
// @Router /api/v1/home/menus [get]
func MenusHandler(c *gin.Context) {
	L.RLock()
	respcode.ResponseSuccess(c, rootMenu)
	L.RUnlock()

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
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{msg=string,data=string}
// @Router /api/v1/data/menus [get]
func DataMenusHandler(ctx *gin.Context) {
	statement := eorm.NewStatement()
	statement = statement.SetTableName("tableManager").
		Select("tableName,year")

	menus := []model.Menus{}

	c := <-global.DBClients
	defer func() {
		global.DBClients <- c
	}()

	err := c.FindAll(nil, statement, &menus)
	if err != nil {
		zap.L().Error("sql exec failed: ", zap.String("", err.Error()))
	}

	M := make(map[string][]string, len(menus))
	for _, v := range menus {
		M[v.TableName[:len(v.TableName)-5]] =
			append(M[v.TableName[:len(v.TableName)-5]], v.Year)
	}

	nodeMap := make(map[*Node2][]Node2)
	for _, v := range menus {
		node := Node2{}
		node.Id, node.Path, node.AuthName = model.Menu(&v)
		if _, ok := M[v.TableName[:len(v.TableName)-5]]; !ok {
			continue
		}
		node.Children = append(node.Children, func() (childNodes []Node2) {
			m := M[v.TableName[:len(v.TableName)-5]]
			for i := 0; i < len(m); i++ {
				childnode := Node2{}
				childnode.Id = node.Id + strconv.Itoa(i)
				childnode.AuthName = m[i]
				childnode.Path = model.Leveltag(node.Path) + childnode.Id
				childnode.Query = strconv.Itoa(i)
				childNodes = append(childNodes, childnode)
			}
			return childNodes
		}()...)
		nodeMap[&node] = append(nodeMap[&node], node)
		delete(M, v.TableName[:len(v.TableName)-5])
	}

	lists := []Node2{}
	for k := range nodeMap {
		node := *k
		lists = append(lists, node)
	}
	sort.Slice(lists, func(i, j int) bool {
		return lists[i].Id < lists[j].Id
	})

	respcode.ResponseSuccess(ctx, lists)

}
