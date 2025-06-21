package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EquipmentRouter struct {}

// InitEquipmentRouter 初始化 设备信息 路由信息
func (s *EquipmentRouter) InitEquipmentRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	EQRouter := Router.Group("EQ").Use(middleware.OperationRecord())
	EQRouterWithoutRecord := Router.Group("EQ")
	EQRouterWithoutAuth := PublicRouter.Group("EQ")
	{
		EQRouter.POST("createEquipment", EQApi.CreateEquipment)   // 新建设备信息
		EQRouter.DELETE("deleteEquipment", EQApi.DeleteEquipment) // 删除设备信息
		EQRouter.DELETE("deleteEquipmentByIds", EQApi.DeleteEquipmentByIds) // 批量删除设备信息
		EQRouter.PUT("updateEquipment", EQApi.UpdateEquipment)    // 更新设备信息
	}
	{
		EQRouterWithoutRecord.GET("findEquipment", EQApi.FindEquipment)        // 根据ID获取设备信息
		EQRouterWithoutRecord.GET("getEquipmentList", EQApi.GetEquipmentList)  // 获取设备信息列表
	}
	{
	    EQRouterWithoutAuth.GET("getEquipmentPublic", EQApi.GetEquipmentPublic)  // 设备信息开放接口
	}
}
