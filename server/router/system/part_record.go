package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PartRecordRouter struct{}

// InitPartRecordRouter 初始化 易损件记录 路由信息
func (s *PartRecordRouter) InitPartRecordRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	PRRouter := Router.Group("PR").Use(middleware.OperationRecord())
	PRRouterWithoutRecord := Router.Group("PR")
	PRRouterWithoutAuth := PublicRouter.Group("PR")
	{
		PRRouter.POST("createPartRecord", PRApi.CreatePartRecord)             // 新建易损件记录
		PRRouter.DELETE("deletePartRecord", PRApi.DeletePartRecord)           // 删除易损件记录
		PRRouter.DELETE("deletePartRecordByIds", PRApi.DeletePartRecordByIds) // 批量删除易损件记录
		PRRouter.PUT("updatePartRecord", PRApi.UpdatePartRecord)              // 更新易损件记录
	}
	{
		PRRouterWithoutRecord.GET("findPartRecord", PRApi.FindPartRecord)       // 根据ID获取易损件记录
		PRRouterWithoutRecord.GET("getPartRecordList", PRApi.GetPartRecordList) // 获取易损件记录列表
	}
	{
		PRRouterWithoutAuth.GET("getPartRecordPublic", PRApi.GetPartRecordPublic) // 易损件记录开放接口
	}
}
