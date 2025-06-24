package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type ProjectInfoService struct{}

// CreateProjectInfo 创建项目信息记录
// Author [yourname](https://github.com/yourname)
func (PIService *ProjectInfoService) CreateProjectInfo(ctx context.Context, PI *system.ProjectInfo) (err error) {
	err = global.GVA_DB.Create(PI).Error
	return err
}

// DeleteProjectInfo 删除项目信息记录
// Author [yourname](https://github.com/yourname)
func (PIService *ProjectInfoService) DeleteProjectInfo(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&system.ProjectInfo{}, "id = ?", ID).Error
	return err
}

// DeleteProjectInfoByIds 批量删除项目信息记录
// Author [yourname](https://github.com/yourname)
func (PIService *ProjectInfoService) DeleteProjectInfoByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.ProjectInfo{}, "id in ?", IDs).Error
	return err
}

// UpdateProjectInfo 更新项目信息记录
// Author [yourname](https://github.com/yourname)
func (PIService *ProjectInfoService) UpdateProjectInfo(ctx context.Context, PI system.ProjectInfo) (err error) {
	err = global.GVA_DB.Model(&system.ProjectInfo{}).Where("id = ?", PI.ID).Updates(&PI).Error
	return err
}

// GetProjectInfo 根据ID获取项目信息记录
// Author [yourname](https://github.com/yourname)
func (PIService *ProjectInfoService) GetProjectInfo(ctx context.Context, ID string) (PI system.ProjectInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&PI).Error
	return
}

// GetProjectInfoInfoList 分页获取项目信息记录
// Author [yourname](https://github.com/yourname)
func (PIService *ProjectInfoService) GetProjectInfoInfoList(ctx context.Context, info systemReq.ProjectInfoSearch) (list []system.ProjectInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.ProjectInfo{})
	var PIs []system.ProjectInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["ID"] = true
	orderMap["CreatedAt"] = true
	orderMap["installation_location"] = true
	orderMap["city"] = true
	orderMap["update_date"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&PIs).Error
	return PIs, total, err
}
func (PIService *ProjectInfoService) GetProjectInfoPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
