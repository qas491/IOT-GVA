package system

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type EquipmentApi struct {}



// CreateEquipment 创建设备信息
// @Tags Equipment
// @Summary 创建设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.Equipment true "创建设备信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /EQ/createEquipment [post]
func (EQApi *EquipmentApi) CreateEquipment(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var EQ system.Equipment
	err := c.ShouldBindJSON(&EQ)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = EQService.CreateEquipment(ctx,&EQ)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteEquipment 删除设备信息
// @Tags Equipment
// @Summary 删除设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.Equipment true "删除设备信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /EQ/deleteEquipment [delete]
func (EQApi *EquipmentApi) DeleteEquipment(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := EQService.DeleteEquipment(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEquipmentByIds 批量删除设备信息
// @Tags Equipment
// @Summary 批量删除设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /EQ/deleteEquipmentByIds [delete]
func (EQApi *EquipmentApi) DeleteEquipmentByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := EQService.DeleteEquipmentByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEquipment 更新设备信息
// @Tags Equipment
// @Summary 更新设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.Equipment true "更新设备信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /EQ/updateEquipment [put]
func (EQApi *EquipmentApi) UpdateEquipment(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var EQ system.Equipment
	err := c.ShouldBindJSON(&EQ)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = EQService.UpdateEquipment(ctx,EQ)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEquipment 用id查询设备信息
// @Tags Equipment
// @Summary 用id查询设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询设备信息"
// @Success 200 {object} response.Response{data=system.Equipment,msg=string} "查询成功"
// @Router /EQ/findEquipment [get]
func (EQApi *EquipmentApi) FindEquipment(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reEQ, err := EQService.GetEquipment(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reEQ, c)
}
// GetEquipmentList 分页获取设备信息列表
// @Tags Equipment
// @Summary 分页获取设备信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.EquipmentSearch true "分页获取设备信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /EQ/getEquipmentList [get]
func (EQApi *EquipmentApi) GetEquipmentList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo systemReq.EquipmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := EQService.GetEquipmentInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetEquipmentPublic 不需要鉴权的设备信息接口
// @Tags Equipment
// @Summary 不需要鉴权的设备信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EQ/getEquipmentPublic [get]
func (EQApi *EquipmentApi) GetEquipmentPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    EQService.GetEquipmentPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的设备信息接口信息",
    }, "获取成功", c)
}
