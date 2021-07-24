package router

import (
	"fmt"
	"log"
	_ "webconsole/docs"
	"webconsole/global"
	"webconsole/internal/middleware"
	"webconsole/pkg/logger"

	v1 "webconsole/internal/router/api/v1"
	v2 "webconsole/internal/router/api/v2"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() (r *gin.Engine, err error) {

	if global.ServerSetting.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r = gin.New()

	r.Use(logger.GinLogger())
	r.Use(logger.GinRecovery(true))
	r.Use(middleware.Cors())

	// 注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// apiv1路由组
	apiv1 := r.Group("/api/v1")

	// 注册路由
	apiv1.POST("/signup", middleware.Translations(), v1.SignUpHandler)

	// 登录路由
	apiv1.POST("/login", v1.LoginHandler)

	cacheKey := make(map[string]struct{})
	// 缓存路由
	cacheGroup := apiv1.Group("/cache")
	{
		cacheGroup.Use(middleware.JWTAuthMiddleware())
		// 操作缓存
		cacheGroup.GET("/hit/*key", v1.CacheCheck,
			func(c *gin.Context) {
				c.Request.URL.Path = "/api/v1/data/info" + c.Param("key") // 将请求的URL修改
				cacheKey[c.Param("key")] = struct{}{}
				r.HandleContext(c) // 继续之后的操作
				c.Abort()
			})
		cacheGroup.DELETE("/hit/*key", v1.CacheDelete)
	}

	// 用户家目录路由
	homeGroup := apiv1.Group("/home")
	{
		homeGroup.Use(middleware.JWTAuthMiddleware())
		homeGroup.GET("/menus", v1.MenusHandler)
		homeGroup.POST("/roles", v1.UpdateRoles,
			func(c *gin.Context) {
				for v := range cacheKey {
					c.Request.URL.Path = "/api/v1/cache/hit" + v // 将请求的URL修改
					c.Request.Method = "DELETE"
					r.HandleContext(c) // 继续之后的操作
					//c.Abort()
				}
			})
		homeGroup.GET("/roles", v1.GetRoles)
		homeGroup.HEAD("/roles", v1.DefaultRoles)

		homeGroup.GET("/users", v1.GetUsers)
	}

	// 数据路由
	dataGroup := apiv1.Group("/data")
	{
		dataGroup.Use(middleware.JWTAuthMiddleware())
		// 数据导航栏路由
		dataGroup.GET("/menus", v1.DataMenusHandler)

		// 数据(记录)操作路由
		infoGroup := dataGroup.Group("/info")
		{
			// 获取记录
			infoGroup.GET("/:infotype/:year/:count", middleware.PathParse,
				middleware.RBACMiddleware, v1.GetInfo)
			// 查询记录
			infoGroup.GET("/:infotype/:year/:count/query", middleware.PathParse,
				middleware.QueryParse, middleware.RBACMiddleware,
				v1.QueryInfo)
			// 修改记录
			infoGroup.POST("/:infotype/:year/:count/:id", middleware.PathParse,
				middleware.QueryParse, middleware.RBACMiddleware,
				v1.UpdateInfo)
			// 删除记录
			infoGroup.DELETE("/:infotype/:year/:count/:id",
				middleware.PathParse, v1.UpdateInfo)
		}

		// 数据(表)操作路由
		tableGroup := dataGroup.Group("/table")
		{
			// 添加表
			tableGroup.POST("/:tabletype/:year", v1.UploadTable,
				func(c *gin.Context) {
					fp := c.GetString("tableName")
					f, err := excelize.OpenFile(fp)
					if err != nil {
						fmt.Println(err)
						return
					}

					rows, _ := f.GetRows("Sheet1")
					for _, r := range rows {
						fmt.Println(r)
					}
				})
			// 删除表

			// 修改表
		}
	}

	apiv2 := r.Group("/api/v2")
	apiv2.Use(middleware.JWTAuthMiddleware())

	iserver := apiv2.Group("/iserver")
	{
		iserver.Any("/services/:dataname/rest/data/*result",
			middleware.IServerRBACMiddleware,
			v2.IServerHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		log.Println("404 page not found")
	})

	return r, nil
}
