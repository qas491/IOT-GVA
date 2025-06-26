package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ItemInfoApi struct{}

// CreateItemInfo 创建易损件
// @Tags ItemInfo
// @Summary 创建易损件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.ItemInfo true "创建易损件"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /II/createItemInfo [post]
func (IIApi *ItemInfoApi) CreateItemInfo(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var II system.ItemInfo
	err := c.ShouldBindJSON(&II)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = IIService.CreateItemInfo(ctx, &II)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteItemInfo 删除易损件
// @Tags ItemInfo
// @Summary 删除易损件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.ItemInfo true "删除易损件"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /II/deleteItemInfo [delete]
func (IIApi *ItemInfoApi) DeleteItemInfo(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := IIService.DeleteItemInfo(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteItemInfoByIds 批量删除易损件
// @Tags ItemInfo
// @Summary 批量删除易损件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /II/deleteItemInfoByIds [delete]
func (IIApi *ItemInfoApi) DeleteItemInfoByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := IIService.DeleteItemInfoByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateItemInfo 更新易损件
// @Tags ItemInfo
// @Summary 更新易损件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.ItemInfo true "更新易损件"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /II/updateItemInfo [put]
func (IIApi *ItemInfoApi) UpdateItemInfo(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var II system.ItemInfo
	err := c.ShouldBindJSON(&II)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = IIService.UpdateItemInfo(ctx, II)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindItemInfo 用id查询易损件
// @Tags ItemInfo
// @Summary 用id查询易损件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询易损件"
// @Success 200 {object} response.Response{data=system.ItemInfo,msg=string} "查询成功"
// @Router /II/findItemInfo [get]
func (IIApi *ItemInfoApi) FindItemInfo(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reII, err := IIService.GetItemInfo(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reII, c)
}

// GetItemInfoList 分页获取易损件列表
// @Tags ItemInfo
// @Summary 分页获取易损件列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.ItemInfoSearch true "分页获取易损件列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /II/getItemInfoList [get]
func (IIApi *ItemInfoApi) GetItemInfoList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo systemReq.ItemInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := IIService.GetItemInfoInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetItemInfoPublic 不需要鉴权的易损件接口
// @Tags ItemInfo
// @Summary 不需要鉴权的易损件接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /II/getItemInfoPublic [get]
func (IIApi *ItemInfoApi) GetItemInfoPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	IIService.GetItemInfoPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的易损件接口信息",
	}, "获取成功", c)
}
