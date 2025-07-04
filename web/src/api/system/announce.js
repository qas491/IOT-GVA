import service from '@/utils/request'

// @Tags Announce
// @Summary 推送版本更新公告
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body object true "公告内容"
// @Success 200 {object} response.Response{msg=string} "推送成功"
// @Router /announce/pushVersionUpdate [post]
export const pushVersionUpdate = (data) => {
  return service({
    url: '/announce/pushVersionUpdate',
    method: 'post',
    data
  })
}

// @Tags Announce
// @Summary 获取公告列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query object true "查询参数"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /announce/getAnnounceList [get]
export const getAnnounceList = (params) => {
  return service({
    url: '/announce/getAnnounceList',
    method: 'get',
    params
  })
}

// @Tags Announce
// @Summary 删除公告
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body object true "删除参数"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /announce/deleteAnnounce [delete]
export const deleteAnnounce = (data) => {
  return service({
    url: '/announce/deleteAnnounce',
    method: 'delete',
    data
  })
} 