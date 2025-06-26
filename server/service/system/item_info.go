package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type ItemInfoService struct{}

// CreateItemInfo 创建易损件记录
// Author [yourname](https://github.com/yourname)
func (IIService *ItemInfoService) CreateItemInfo(ctx context.Context, II *system.ItemInfo) (err error) {
	err = global.GVA_DB.Create(II).Error
	return err
}

// DeleteItemInfo 删除易损件记录
// Author [yourname](https://github.com/yourname)
func (IIService *ItemInfoService) DeleteItemInfo(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&system.ItemInfo{}, "id = ?", ID).Error
	return err
}

// DeleteItemInfoByIds 批量删除易损件记录
// Author [yourname](https://github.com/yourname)
func (IIService *ItemInfoService) DeleteItemInfoByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.ItemInfo{}, "id in ?", IDs).Error
	return err
}

// UpdateItemInfo 更新易损件记录
// Author [yourname](https://github.com/yourname)
func (IIService *ItemInfoService) UpdateItemInfo(ctx context.Context, II system.ItemInfo) (err error) {
	err = global.GVA_DB.Model(&system.ItemInfo{}).Where("id = ?", II.ID).Updates(&II).Error
	return err
}

// GetItemInfo 根据ID获取易损件记录
// Author [yourname](https://github.com/yourname)
func (IIService *ItemInfoService) GetItemInfo(ctx context.Context, ID string) (II system.ItemInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&II).Error
	return
}

// GetItemInfoInfoList 分页获取易损件记录
// Author [yourname](https://github.com/yourname)
func (IIService *ItemInfoService) GetItemInfoInfoList(ctx context.Context, info systemReq.ItemInfoSearch) (list []system.ItemInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.ItemInfo{})
	var IIs []system.ItemInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&IIs).Error
	return IIs, total, err
}
func (IIService *ItemInfoService) GetItemInfoPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
