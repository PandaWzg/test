package initialize

import (
	"github.com/gin-gonic/gin"
	"wm-infoflow-api-go/middleware"
	"wm-infoflow-api-go/router"
)

// Routers 初始化路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	PublicGroup := Router.Group("/api")
	{
		// 注册基础功能路由 不做鉴权
		router.InitBaseRouter(PublicGroup)
	}

	PrivateGroup := Router.Group("/api")
	PrivateGroup.Use(middleware.Cors()).Use(middleware.CheckAuthToken())
	{
		router.InitMenuRouter(PrivateGroup)// 注册menu路由
	}
	return Router
}
