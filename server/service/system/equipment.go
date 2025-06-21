
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type EquipmentService struct {}
// CreateEquipment 创建设备信息记录
// Author [yourname](https://github.com/yourname)
func (EQService *EquipmentService) CreateEquipment(ctx context.Context, EQ *system.Equipment) (err error) {
	err = global.GVA_DB.Create(EQ).Error
	return err
}

// DeleteEquipment 删除设备信息记录
// Author [yourname](https://github.com/yourname)
func (EQService *EquipmentService)DeleteEquipment(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&system.Equipment{},"id = ?",ID).Error
	return err
}

// DeleteEquipmentByIds 批量删除设备信息记录
// Author [yourname](https://github.com/yourname)
func (EQService *EquipmentService)DeleteEquipmentByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.Equipment{},"id in ?",IDs).Error
	return err
}

// UpdateEquipment 更新设备信息记录
// Author [yourname](https://github.com/yourname)
func (EQService *EquipmentService)UpdateEquipment(ctx context.Context, EQ system.Equipment) (err error) {
	err = global.GVA_DB.Model(&system.Equipment{}).Where("id = ?",EQ.ID).Updates(&EQ).Error
	return err
}

// GetEquipment 根据ID获取设备信息记录
// Author [yourname](https://github.com/yourname)
func (EQService *EquipmentService)GetEquipment(ctx context.Context, ID string) (EQ system.Equipment, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&EQ).Error
	return
}
// GetEquipmentInfoList 分页获取设备信息记录
// Author [yourname](https://github.com/yourname)
func (EQService *EquipmentService)GetEquipmentInfoList(ctx context.Context, info systemReq.EquipmentSearch) (list []system.Equipment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&system.Equipment{})
    var EQs []system.Equipment
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.EquipmentNumber != nil && *info.EquipmentNumber != "" {
        db = db.Where("equipment_number LIKE ?", "%"+ *info.EquipmentNumber+"%")
    }
    if info.EquipmentName != nil && *info.EquipmentName != "" {
        db = db.Where("equipment_name LIKE ?", "%"+ *info.EquipmentName+"%")
    }
    if info.OperationalStatus != nil && *info.OperationalStatus != "" {
        db = db.Where("operational_status LIKE ?", "%"+ *info.OperationalStatus+"%")
    }
    if info.NunningState != nil && *info.NunningState != "" {
        db = db.Where("nunning_state LIKE ?", "%"+ *info.NunningState+"%")
    }
    if info.SctiveState != nil && *info.SctiveState != "" {
        db = db.Where("sctive_state LIKE ?", "%"+ *info.SctiveState+"%")
    }
    if info.SoftwareVersion != nil && *info.SoftwareVersion != "" {
        db = db.Where("software_version LIKE ?", "%"+ *info.SoftwareVersion+"%")
    }
    if info.AfterSalesPersonnel != nil && *info.AfterSalesPersonnel != "" {
        db = db.Where("after_sales_personnel LIKE ?", "%"+ *info.AfterSalesPersonnel+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&EQs).Error
	return  EQs, total, err
}
func (EQService *EquipmentService)GetEquipmentPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
