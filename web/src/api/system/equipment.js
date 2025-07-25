import service from '@/utils/request'
// @Tags Equipment
// @Summary 创建全部设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Equipment true "创建全部设备"
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
// @Summary 删除全部设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Equipment true "删除全部设备"
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
// @Summary 批量删除全部设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除全部设备"
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
// @Summary 更新全部设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Equipment true "更新全部设备"
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
// @Summary 用id查询全部设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Equipment true "用id查询全部设备"
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
// @Summary 分页获取全部设备列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取全部设备列表"
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
// @Summary 不需要鉴权的全部设备接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.EquipmentSearch true "分页获取全部设备列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EQ/getEquipmentPublic [get]

// Dashboard 仪表盘
// @Tags Equipment
// @Summary 仪表盘
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /EQ/Dashboard [GET]
export const Dashboard = (params) => {
  return service({
    url: '/EQ/Dashboard',
    method: 'GET',
    params
  })
}

// QueryDeviceCount 查询后端设备数量
// @Tags Equipment
// @Summary 查询后端设备数量
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /EQ/devicecount [GET]
export const devicecount = (params) => {
  return service({
    url: '/EQ/devicecount',
    method: 'GET',
    params
  })
}
// QueryDeviceCountByStatus  根据运营状态查询设备数量
// @Tags Equipment
// @Summary 根据运营状态查询设备数量
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /EQ/queryDeviceCountByStatus  [GET]
export const queryDeviceCountByStatus  = (params) => {
  return service({
    url: '/EQ/queryDeviceCountByStatus',
    method: 'GET',
    params
  })
}

// Dashboard 获取仪表盘统计数据
// @Tags Equipment
// @Summary 获取仪表盘统计数据
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EQ/Dashboard [GET]
export const getDashboardStats = () => {
  return service({
    url: '/EQ/Dashboard',
    method: 'GET'
  })
}

// 批量激活设备
export const activateEquipment = (data) => {
  return service({
    url: '/EQ/activate',
    method: 'post',
    data
  })
}
