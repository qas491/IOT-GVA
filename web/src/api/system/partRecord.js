import service from '@/utils/request'
// @Tags PartRecord
// @Summary 创建易损件记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PartRecord true "创建易损件记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PR/createPartRecord [post]
export const createPartRecord = (data) => {
  return service({
    url: '/PR/createPartRecord',
    method: 'post',
    data
  })
}

// @Tags PartRecord
// @Summary 删除易损件记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PartRecord true "删除易损件记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PR/deletePartRecord [delete]
export const deletePartRecord = (params) => {
  return service({
    url: '/PR/deletePartRecord',
    method: 'delete',
    params
  })
}

// @Tags PartRecord
// @Summary 批量删除易损件记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除易损件记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PR/deletePartRecord [delete]
export const deletePartRecordByIds = (params) => {
  return service({
    url: '/PR/deletePartRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags PartRecord
// @Summary 更新易损件记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PartRecord true "更新易损件记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PR/updatePartRecord [put]
export const updatePartRecord = (data) => {
  return service({
    url: '/PR/updatePartRecord',
    method: 'put',
    data
  })
}

// @Tags PartRecord
// @Summary 用id查询易损件记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PartRecord true "用id查询易损件记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PR/findPartRecord [get]
export const findPartRecord = (params) => {
  return service({
    url: '/PR/findPartRecord',
    method: 'get',
    params
  })
}

// @Tags PartRecord
// @Summary 分页获取易损件记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取易损件记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PR/getPartRecordList [get]
export const getPartRecordList = (params) => {
  return service({
    url: '/PR/getPartRecordList',
    method: 'get',
    params
  })
}

// @Tags PartRecord
// @Summary 不需要鉴权的易损件记录接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.PartRecordSearch true "分页获取易损件记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PR/getPartRecordPublic [get]
export const getPartRecordPublic = () => {
  return service({
    url: '/PR/getPartRecordPublic',
    method: 'get',
  })
}
