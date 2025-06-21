import service from '@/utils/request'
// @Tags Equipment
// @Summary 创建设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Equipment true "创建设备信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /EQ/createEquipment [post]
export const createEquipment = (data) => {
  return service({
    url: '/EQ/createEquipment',
    method: 'post',
    data
  })
}

// @Tags Equipment
// @Summary 删除设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Equipment true "删除设备信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EQ/deleteEquipment [delete]
export const deleteEquipment = (params) => {
  return service({
    url: '/EQ/deleteEquipment',
    method: 'delete',
    params
  })
}

// @Tags Equipment
// @Summary 批量删除设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除设备信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EQ/deleteEquipment [delete]
export const deleteEquipmentByIds = (params) => {
  return service({
    url: '/EQ/deleteEquipmentByIds',
    method: 'delete',
    params
  })
}

// @Tags Equipment
// @Summary 更新设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Equipment true "更新设备信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /EQ/updateEquipment [put]
export const updateEquipment = (data) => {
  return service({
    url: '/EQ/updateEquipment',
    method: 'put',
    data
  })
}

// @Tags Equipment
// @Summary 用id查询设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Equipment true "用id查询设备信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /EQ/findEquipment [get]
export const findEquipment = (params) => {
  return service({
    url: '/EQ/findEquipment',
    method: 'get',
    params
  })
}

// @Tags Equipment
// @Summary 分页获取设备信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取设备信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /EQ/getEquipmentList [get]
export const getEquipmentList = (params) => {
  return service({
    url: '/EQ/getEquipmentList',
    method: 'get',
    params
  })
}

// @Tags Equipment
// @Summary 不需要鉴权的设备信息接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.EquipmentSearch true "分页获取设备信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EQ/getEquipmentPublic [get]
export const getEquipmentPublic = () => {
  return service({
    url: '/EQ/getEquipmentPublic',
    method: 'get',
  })
}
