import service from '@/utils/request'
// @Tags ItemInfo
// @Summary 创建易损件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ItemInfo true "创建易损件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /II/createItemInfo [post]
export const createItemInfo = (data) => {
  return service({
    url: '/II/createItemInfo',
    method: 'post',
    data
  })
}

// @Tags ItemInfo
// @Summary 删除易损件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ItemInfo true "删除易损件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /II/deleteItemInfo [delete]
export const deleteItemInfo = (params) => {
  return service({
    url: '/II/deleteItemInfo',
    method: 'delete',
    params
  })
}

// @Tags ItemInfo
// @Summary 批量删除易损件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除易损件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /II/deleteItemInfo [delete]
export const deleteItemInfoByIds = (params) => {
  return service({
    url: '/II/deleteItemInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags ItemInfo
// @Summary 更新易损件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ItemInfo true "更新易损件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /II/updateItemInfo [put]
export const updateItemInfo = (data) => {
  return service({
    url: '/II/updateItemInfo',
    method: 'put',
    data
  })
}

// @Tags ItemInfo
// @Summary 用id查询易损件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ItemInfo true "用id查询易损件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /II/findItemInfo [get]
export const findItemInfo = (params) => {
  return service({
    url: '/II/findItemInfo',
    method: 'get',
    params
  })
}

// @Tags ItemInfo
// @Summary 分页获取易损件列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取易损件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /II/getItemInfoList [get]
export const getItemInfoList = (params) => {
  return service({
    url: '/II/getItemInfoList',
    method: 'get',
    params
  })
}

// @Tags ItemInfo
// @Summary 不需要鉴权的易损件接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.ItemInfoSearch true "分页获取易损件列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /II/getItemInfoPublic [get]
export const getItemInfoPublic = () => {
  return service({
    url: '/II/getItemInfoPublic',
    method: 'get',
  })
}
