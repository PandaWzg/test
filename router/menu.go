package router

import (
	"github.com/gin-gonic/gin"
	v1 "wm-infoflow-api-go/api/v1"
	"wm-infoflow-api-go/middleware"
)

func InitMenuRouter(Router *gin.RouterGroup) {
	MenuRouter := Router.Group("/menu").Use(middleware.CheckAuthToken())
	{
		//获取菜单
		MenuRouter.GET("list",v1.MenuList)
	}
}

