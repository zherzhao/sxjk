package router

import (
	_ "webconsole/docs"
	"webconsole/global"
	"webconsole/internal/middleware"
	"webconsole/internal/router/api/v2/objects"
	"webconsole/internal/service/heartbeat"
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

	// 注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// apiv1路由组
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
		// 配置缓存服务
		c := cache.New(global.CacheSetting.CacheType, global.CacheSetting.TTL)
		// 开启TCP缓存服务 监听TCP请求
		go tcp.New(c).Listen()

		s := v1.NewServer(c)

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
		cacheGroup.PUT("/update/*key", s.UpdateHandler)
	}

	// 数据查询路由
	infoGroup := apiv1.Group("/info")
	{
		info := v1.NewInfo()

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

	// apiv2路由组
	apiv2 := r.Group("/api/v2")
	ossGroup := apiv2.Group("")
	{
		go heartbeat.ListenHeartBeat()
		//OSS存储服务
		ossGroup.PUT("/OSS/objects/:file", objects.Put)
		//ossGroup.POST("/OSS/objects/:file", objects.Post)
		ossGroup.GET("/OSS/objects/:file", objects.Get)
		//ossGroup.DELETE("/OSS/objects/:file", objects.Delete)

	}

	return r, nil
}
