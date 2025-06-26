package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type PartRecordService struct{}

// CreatePartRecord 创建易损件记录记录
// Author [yourname](https://github.com/yourname)
func (PRService *PartRecordService) CreatePartRecord(ctx context.Context, PR *system.PartRecord) (err error) {
	err = global.GVA_DB.Create(PR).Error
	return err
}

// DeletePartRecord 删除易损件记录记录
// Author [yourname](https://github.com/yourname)
func (PRService *PartRecordService) DeletePartRecord(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&system.PartRecord{}, "id = ?", ID).Error
	return err
}

// DeletePartRecordByIds 批量删除易损件记录记录
// Author [yourname](https://github.com/yourname)
func (PRService *PartRecordService) DeletePartRecordByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.PartRecord{}, "id in ?", IDs).Error
	return err
}

// UpdatePartRecord 更新易损件记录记录
// Author [yourname](https://github.com/yourname)
func (PRService *PartRecordService) UpdatePartRecord(ctx context.Context, PR system.PartRecord) (err error) {
	err = global.GVA_DB.Model(&system.PartRecord{}).Where("id = ?", PR.ID).Updates(&PR).Error
	return err
}

// GetPartRecord 根据ID获取易损件记录记录
// Author [yourname](https://github.com/yourname)
func (PRService *PartRecordService) GetPartRecord(ctx context.Context, ID string) (PR system.PartRecord, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&PR).Error
	return
}

// GetPartRecordInfoList 分页获取易损件记录记录
// Author [yourname](https://github.com/yourname)
func (PRService *PartRecordService) GetPartRecordInfoList(ctx context.Context, info systemReq.PartRecordSearch) (list []system.PartRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.PartRecord{})
	var PRs []system.PartRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.PartModel != nil && *info.PartModel != "" {
		db = db.Where("part_model LIKE ?", "%"+*info.PartModel+"%")
	}
	if info.PartNumber != nil && *info.PartNumber != "" {
		db = db.Where("part_number LIKE ?", "%"+*info.PartNumber+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&PRs).Error
	return PRs, total, err
}
func (PRService *PartRecordService) GetPartRecordPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
