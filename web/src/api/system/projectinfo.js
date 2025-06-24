import service from '@/utils/request'
// @Tags ProjectInfo
// @Summary 创建项目信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProjectInfo true "创建项目信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PI/createProjectInfo [post]
export const createProjectInfo = (data) => {
  return service({
    url: '/PI/createProjectInfo',
    method: 'post',
    data
  })
}

// @Tags ProjectInfo
// @Summary 删除项目信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProjectInfo true "删除项目信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PI/deleteProjectInfo [delete]
export const deleteProjectInfo = (params) => {
  return service({
    url: '/PI/deleteProjectInfo',
    method: 'delete',
    params
  })
}

// @Tags ProjectInfo
// @Summary 批量删除项目信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除项目信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PI/deleteProjectInfo [delete]
export const deleteProjectInfoByIds = (params) => {
  return service({
    url: '/PI/deleteProjectInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags ProjectInfo
// @Summary 更新项目信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProjectInfo true "更新项目信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PI/updateProjectInfo [put]
export const updateProjectInfo = (data) => {
  return service({
    url: '/PI/updateProjectInfo',
    method: 'put',
    data
  })
}

// @Tags ProjectInfo
// @Summary 用id查询项目信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ProjectInfo true "用id查询项目信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PI/findProjectInfo [get]
export const findProjectInfo = (params) => {
  return service({
    url: '/PI/findProjectInfo',
    method: 'get',
    params
  })
}

// @Tags ProjectInfo
// @Summary 分页获取项目信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取项目信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PI/getProjectInfoList [get]
export const getProjectInfoList = (params) => {
  return service({
    url: '/PI/getProjectInfoList',
    method: 'get',
    params
  })
}

// @Tags ProjectInfo
// @Summary 不需要鉴权的项目信息接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.ProjectInfoSearch true "分页获取项目信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PI/getProjectInfoPublic [get]
export const getProjectInfoPublic = () => {
  return service({
    url: '/PI/getProjectInfoPublic',
    method: 'get',
  })
}
