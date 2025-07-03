package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EquipmentRouter struct{}

func (s *EquipmentRouter) InitEquipmentRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	EQRouter := Router.Group("EQ").Use(middleware.OperationRecord())
	EQRouterWithoutRecord := Router.Group("EQ")
	EQRouterWithoutAuth := PublicRouter.Group("EQ")
	{
		EQRouter.POST(":deviceID/updateStatus", EQApi.UpdateEquipmentStatus)
		EQRouter.POST("createEquipment", EQApi.CreateEquipment)
		EQRouter.DELETE("deleteEquipment", EQApi.DeleteEquipment)
		EQRouter.DELETE("deleteEquipmentByIds", EQApi.DeleteEquipmentByIds)
		EQRouter.PUT("updateEquipment", EQApi.UpdateEquipment)
		EQRouter.POST("/announce/pushVersionUpdate", EQApi.PushVersionUpdateAPI)
	}
	{

		EQRouterWithoutRecord.GET("findEquipment", EQApi.FindEquipment)
		EQRouterWithoutRecord.GET("getEquipmentList", EQApi.GetEquipmentList)
		EQRouterWithoutAuth.GET("getEquipmentPublic", EQApi.GetEquipmentPublic)
	}
	{
		//		EQRouterWithoutAuth.GET("getEquipmentPublic", EQApi.GetEquipmentPublic)
		EQRouterWithoutAuth.GET("devicecount", EQApi.QueryDeviceCount)
		EQRouterWithoutAuth.GET("queryDeviceCountByStatus", EQApi.QueryDeviceCountByStatus)
		EQRouterWithoutAuth.GET("Dashboard", EQApi.Dashboard)
		EQRouterWithoutAuth.GET("/ws/announce", EQApi.WebSocketAnnounce)
	}

}
