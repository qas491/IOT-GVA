package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProjectInfoApi struct{}

// CreateProjectInfo 创建项目信息
// @Tags ProjectInfo
// @Summary 创建项目信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.ProjectInfo true "创建项目信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /PI/createProjectInfo [post]
func (PIApi *ProjectInfoApi) CreateProjectInfo(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var PI system.ProjectInfo
	err := c.ShouldBindJSON(&PI)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = PIService.CreateProjectInfo(ctx, &PI)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteProjectInfo 删除项目信息
// @Tags ProjectInfo
// @Summary 删除项目信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.ProjectInfo true "删除项目信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /PI/deleteProjectInfo [delete]
func (PIApi *ProjectInfoApi) DeleteProjectInfo(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := PIService.DeleteProjectInfo(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteProjectInfoByIds 批量删除项目信息
// @Tags ProjectInfo
// @Summary 批量删除项目信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /PI/deleteProjectInfoByIds [delete]
func (PIApi *ProjectInfoApi) DeleteProjectInfoByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := PIService.DeleteProjectInfoByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateProjectInfo 更新项目信息
// @Tags ProjectInfo
// @Summary 更新项目信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.ProjectInfo true "更新项目信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /PI/updateProjectInfo [put]
func (PIApi *ProjectInfoApi) UpdateProjectInfo(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var PI system.ProjectInfo
	err := c.ShouldBindJSON(&PI)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = PIService.UpdateProjectInfo(ctx, PI)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindProjectInfo 用id查询项目信息
// @Tags ProjectInfo
// @Summary 用id查询项目信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询项目信息"
// @Success 200 {object} response.Response{data=system.ProjectInfo,msg=string} "查询成功"
// @Router /PI/findProjectInfo [get]
func (PIApi *ProjectInfoApi) FindProjectInfo(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rePI, err := PIService.GetProjectInfo(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rePI, c)
}

// GetProjectInfoList 分页获取项目信息列表
// @Tags ProjectInfo
// @Summary 分页获取项目信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.ProjectInfoSearch true "分页获取项目信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /PI/getProjectInfoList [get]
func (PIApi *ProjectInfoApi) GetProjectInfoList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo systemReq.ProjectInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := PIService.GetProjectInfoInfoList(ctx, pageInfo)
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

// GetProjectInfoPublic 不需要鉴权的项目信息接口
// @Tags ProjectInfo
// @Summary 不需要鉴权的项目信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PI/getProjectInfoPublic [get]
func (PIApi *ProjectInfoApi) GetProjectInfoPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	PIService.GetProjectInfoPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的项目信息接口信息",
	}, "获取成功", c)
}
