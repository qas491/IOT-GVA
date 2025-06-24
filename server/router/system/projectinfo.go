package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProjectInfoRouter struct{}

// InitProjectInfoRouter 初始化 项目信息 路由信息
func (s *ProjectInfoRouter) InitProjectInfoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	PIRouter := Router.Group("PI").Use(middleware.OperationRecord())
	PIRouterWithoutRecord := Router.Group("PI")
	PIRouterWithoutAuth := PublicRouter.Group("PI")
	{
		PIRouter.POST("createProjectInfo", PIApi.CreateProjectInfo)             // 新建项目信息
		PIRouter.DELETE("deleteProjectInfo", PIApi.DeleteProjectInfo)           // 删除项目信息
		PIRouter.DELETE("deleteProjectInfoByIds", PIApi.DeleteProjectInfoByIds) // 批量删除项目信息
		PIRouter.PUT("updateProjectInfo", PIApi.UpdateProjectInfo)              // 更新项目信息
	}
	{
		PIRouterWithoutRecord.GET("findProjectInfo", PIApi.FindProjectInfo)       // 根据ID获取项目信息
		PIRouterWithoutRecord.GET("getProjectInfoList", PIApi.GetProjectInfoList) // 获取项目信息列表
	}
	{
		PIRouterWithoutAuth.GET("getProjectInfoPublic", PIApi.GetProjectInfoPublic) // 项目信息开放接口
	}
}
