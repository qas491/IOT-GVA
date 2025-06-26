package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ItemInfoRouter struct{}

// InitItemInfoRouter 初始化 易损件 路由信息
func (s *ItemInfoRouter) InitItemInfoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	IIRouter := Router.Group("II").Use(middleware.OperationRecord())
	IIRouterWithoutRecord := Router.Group("II")
	IIRouterWithoutAuth := PublicRouter.Group("II")
	{
		IIRouter.POST("createItemInfo", IIApi.CreateItemInfo)             // 新建易损件
		IIRouter.DELETE("deleteItemInfo", IIApi.DeleteItemInfo)           // 删除易损件
		IIRouter.DELETE("deleteItemInfoByIds", IIApi.DeleteItemInfoByIds) // 批量删除易损件
		IIRouter.PUT("updateItemInfo", IIApi.UpdateItemInfo)              // 更新易损件
	}
	{
		IIRouterWithoutRecord.GET("findItemInfo", IIApi.FindItemInfo)       // 根据ID获取易损件
		IIRouterWithoutRecord.GET("getItemInfoList", IIApi.GetItemInfoList) // 获取易损件列表
	}
	{
		IIRouterWithoutAuth.GET("getItemInfoPublic", IIApi.GetItemInfoPublic) // 易损件开放接口
	}
}
