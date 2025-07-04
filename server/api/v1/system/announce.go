package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type AnnounceApi struct{}

// 模拟公告数据存储
var announceList = []map[string]interface{}{
	{
		"id":        1,
		"type":      "version_update",
		"title":     "系统版本更新 v1.2.3",
		"content":   "新版本包含多项功能优化和bug修复，建议及时更新。",
		"version":   "v1.2.3",
		"url":       "https://example.com/release-notes",
		"createdAt": "2024-01-01 10:00:00",
	},
	{
		"id":        2,
		"type":      "system_maintenance",
		"title":     "系统维护通知",
		"content":   "系统将于今晚22:00-24:00进行维护，期间可能影响正常使用。",
		"version":   "",
		"url":       "",
		"createdAt": "2024-01-02 15:30:00",
	},
}

// GetAnnounceList 获取公告列表
// @Tags Announce
// @Summary 获取公告列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param type query string false "公告类型"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /announce/getAnnounceList [get]
func (a *AnnounceApi) GetAnnounceList(c *gin.Context) {
	/*page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")*/
	announceType := c.Query("type")

	// 这里应该从数据库查询，现在用模拟数据
	filteredList := announceList
	if announceType != "" {
		var filtered []map[string]interface{}
		for _, item := range announceList {
			if item["type"] == announceType {
				filtered = append(filtered, item)
			}
		}
		filteredList = filtered
	}

	response.OkWithDetailed(response.PageResult{
		List:     filteredList,
		Total:    int64(len(filteredList)),
		Page:     1,
		PageSize: 10,
	}, "获取成功", c)
}

// DeleteAnnounce 删除公告
// @Tags Announce
// @Summary 删除公告
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body map[string]interface{} true "删除参数"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /announce/deleteAnnounce [delete]
func (a *AnnounceApi) DeleteAnnounce(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	id, ok := req["id"].(float64)
	if !ok {
		response.FailWithMessage("缺少公告ID", c)
		return
	}

	// 这里应该从数据库删除，现在用模拟数据
	for i, item := range announceList {
		if item["id"] == int(id) {
			announceList = append(announceList[:i], announceList[i+1:]...)
			break
		}
	}

	response.OkWithMessage("删除成功", c)
}
