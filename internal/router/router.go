package router

import (
	_ "webconsole/docs"
	"webconsole/global"
	"webconsole/internal/middleware"
	"webconsole/pkg/cache"
	"webconsole/pkg/cache/tcp"
	"webconsole/pkg/logger"

	v1 "webconsole/internal/router/api/v1"

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

	// 配置缓存服务
	c := cache.New(global.CacheSetting.CacheType, global.CacheSetting.TTL)
	// 开启TCP缓存服务 监听TCP请求
	go tcp.New(c).Listen()

	s := v1.NewServer(c)
	info := v1.NewInfo()

	// 注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// api路由组
	apiv1 := r.Group("/api/v1")

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
		// 操作缓存
		cacheGroup.Use(middleware.JWTAuthMiddleware(), middleware.PathParse)

		cacheGroup.GET("/hit/*key", s.CacheCheck, func(c *gin.Context) {
			miss := c.GetBool("miss") // 检查是否命中缓存
			if miss {
				c.Request.URL.Path = "/api/v1/info" + c.Param("key") // 将请求的URL修改
				r.HandleContext(c)                                   // 继续之后的操作

			}
		})

		// 获取缓存状态
		cacheGroup.GET("/status/", s.StatusHandler)
	}

	// 数据查询路由
	infoGroup := apiv1.Group("/info")
	{
		infoGroup.Use(middleware.JWTAuthMiddleware(), middleware.PathParse)

		infoGroup.GET("/:infotype/:count",
			middleware.QueryRouter,
			info.GetUpdateInfo,
			func(c *gin.Context) {
				if c.GetString("type") == "mem" {
					r.HandleContext(c) //继续之后的操作
				}
			})
	}
	return r, nil
}
