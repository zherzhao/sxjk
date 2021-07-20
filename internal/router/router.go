package router

import (
	"log"
	_ "webconsole/docs"
	"webconsole/global"
	"webconsole/internal/middleware"
	"webconsole/pkg/logger"

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

	// 注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// apiv1路由组
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.CorsV1())

	// 注册路由
	apiv1.POST("/signup", middleware.Translations(), v1.SignUpHandler)

	// 登录路由
	apiv1.POST("/login", v1.LoginHandler)

	// 用户家目录路由
	homeGroup := apiv1.Group("/home")
	{
		homeGroup.Use(middleware.JWTAuthMiddleware())
		homeGroup.GET("/menus", v1.MenusHandler)
	}

	// 缓存路由
	cacheGroup := apiv1.Group("/cache")
	{
		// 权限验证
		cacheGroup.Use(middleware.JWTAuthMiddleware())
		// 操作缓存
		cacheGroup.GET("/hit/*key", v1.CacheCheck,
			func(c *gin.Context) {
				c.Request.URL.Path = "/api/v1/data/info" + c.Param("key") // 将请求的URL修改
				r.HandleContext(c)                                        // 继续之后的操作
				c.Abort()
			})

	}

	// 数据路由
	dataGroup := apiv1.Group("/data")
	{
		dataGroup.Use(middleware.JWTAuthMiddleware())
		// 数据导航栏路由
		dataGroup.GET("/menus", v1.DataMenusHandler)

		// 数据操作路由
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

			// 添加记录
			//infoGroup.POST("/:infotype/:year/:count/:id", middleware.PathParse, info.UpdateInfo)
		}

		// 数据添加路由
		// 添加表

		// 数据删除路由
		// 删除表

		// 数据修改路由
		// 修改表
	}

	apiv2 := r.Group("/api/v2")
	apiv2.Use(middleware.CorsV2())

	iserver := apiv2.Group("/iserver")
	{
		iserver.Any("/services/:dataname/rest/data/*result", v2.IServerHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		log.Println("404 page not found")

	})

	return r, nil
}
