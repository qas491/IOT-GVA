package system

import (
	"github.com/gin-gonic/gin"
)

type AnnounceRouter struct{}

// InitAnnounceRouter 初始化公告路由
func (s *AnnounceRouter) InitAnnounceRouter(Router *gin.RouterGroup) {
	/*announceRouter := Router.Group("announce").Use(middleware.OperationRecord())
	announceRouterWithoutRecord := Router.Group("announce")*/
	/*announceApi := v1.ApiGroupApp.SystemApiGroup.AnnounceApi
	{
		announceRouter.POST("pushVersionUpdate", announceApi.PushVersionUpdateAPI) // 推送版本更新公告
		announceRouter.GET("getAnnounceList", announceApi.GetAnnounceList)         // 获取公告列表
		announceRouter.DELETE("deleteAnnounce", announceApi.DeleteAnnounce)        // 删除公告
	}
	{
		announceRouterWithoutRecord.GET("getAnnounceList", announceApi.GetAnnounceList) // 获取公告列表
	}*/
}
