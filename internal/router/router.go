package router

import (
	"sync"
	_ "webconsole/docs"
	"webconsole/global"
	"webconsole/internal/middleware"
	"webconsole/pkg/logger"
	"webconsole/pkg/respcode"

	v1 "webconsole/internal/router/api/v1"
	v2 "webconsole/internal/router/api/v2"

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
	apiv1.POST("/signup", v1.SignUpHandler)
	apiv1.GET("/signup",
		middleware.JWTAuthMiddleware(),
		middleware.UserRBAC("application"),
		v1.SignUpVerifyHandler)

	// 登录路由
	apiv1.POST("/login", v1.LoginHandler)

	//cacheKey := make(map[string]struct{})
	cacheKey := sync.Map{}
	// 缓存路由
	cacheGroup := apiv1.Group("/cache")
	{
		cacheGroup.Use(middleware.JWTAuthMiddleware())
		// 获取缓存
		cacheGroup.GET("/hit/*key", v1.CacheCheckHandler(r, cacheKey))
		// 清除缓存
		cacheGroup.DELETE("/hit/*key", v1.CacheDeleteHandler)
	}

	// 用户家目录路由
	homeGroup := apiv1.Group("/home")
	{
		homeGroup.Use(middleware.JWTAuthMiddleware())
		// 获取导航栏数据
		homeGroup.GET("/menus", v1.MenusHandler)

		// 权限
		homeGroup.POST("/roles", middleware.RoleRBAC("edit-role"),
			v1.UpdateRolesHandler, v1.CacheClearHandler(r, cacheKey))
		homeGroup.GET("/roles", middleware.RoleRBAC("check-role"), v1.GetRoles)
		homeGroup.HEAD("/roles", middleware.RoleRBAC("default-role"),
			v1.DefaultRolesHandler)

		// 用户
		homeGroup.GET("/users", middleware.UserRBAC("check-user"), v1.GetUsersHandler)
		homeGroup.GET("/users/query", middleware.UserRBAC("query-user"),
			middleware.QueryParse, v1.QueryUsersHandler)
		homeGroup.POST("/users", middleware.UserRBAC("edit-user"),
			v1.UpdateUsersHandler)
		homeGroup.DELETE("/users/:id", middleware.UserRBAC("del-user"),
			v1.DeleteUsersHandler)
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
			infoGroup.GET("/:infotype/:year", middleware.PathParse,
				middleware.RecordRBAC, v1.GetInfo)
			// 查询记录
			infoGroup.POST("/:infotype/:year/query", middleware.PathParse,
				middleware.RecordRBAC, v1.QueryInfo)
			// 修改记录
			infoGroup.POST("/:infotype/:year",
				middleware.PathParse, middleware.RecordRBAC,
				v1.UpdateInfo, middleware.ClearCache(r, cacheKey))
			// 删除记录
			infoGroup.DELETE("/:infotype/:year/:id",
				middleware.PathParse, v1.DeleteInfo,
				middleware.ClearCache(r, cacheKey))
		}

		// 数据(表)操作路由
		tableGroup := dataGroup.Group("/table")
		{
			// 添加表
			tableGroup.POST("/:tabletype/:year", middleware.TableRBAC("add-table"),
				v1.UploadTable, v1.ParseTable)
			// 删除表
			tableGroup.DELETE("/:tabletype/:year", middleware.TableRBAC("del-table"),
				v1.DeleteTable)
		}
	}

	apiv2 := r.Group("/api/v2")
	apiv2.Use(middleware.JWTAuthMiddleware())

	apiv2.POST("/query/iserver/services/:dataname/rest/data/*result",
		middleware.IServerRBAC("query"),
		v2.IServerQueryHandler)
	apiv2.POST("/search/iserver/services/:dataname/rest/data/*result",
		middleware.IServerRBAC("search"),
		v2.IServerSearchHandler)

	r.NoRoute(func(c *gin.Context) {
		respcode.ResponseNotFound(c)
	})

	return r, nil
}
