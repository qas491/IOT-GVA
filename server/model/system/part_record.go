// 自动生成模板PartRecord
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 易损件记录 结构体  PartRecord
type PartRecord struct {
	global.GVA_MODEL
	ExportBatch      *bool      `json:"exportBatch" form:"exportBatch" gorm:"column:export_batch;"`                 //批量导出
	PartModel        *string    `json:"partModel" form:"partModel" gorm:"column:part_model;" binding:"required"`    //型号
	PartNumber       *string    `json:"partNumber" form:"partNumber" gorm:"column:part_number;" binding:"required"` //编号
	UsageCount       *int       `json:"usageCount" form:"usageCount" gorm:"column:usage_count;"`                    //次数
	UsageDays        *int       `json:"usageDays" form:"usageDays" gorm:"column:usage_days;"`                       //天数
	FirstUsageDate   *time.Time `json:"firstUsageDate" form:"firstUsageDate" gorm:"column:first_usage_date;"`       //上机日
	SyncDate         *time.Time `json:"syncDate" form:"syncDate" gorm:"column:sync_date;"`                          //同步日
	UsageHistoryLink *string    `json:"usageHistoryLink" form:"usageHistoryLink" gorm:"column:usage_history_link;"` //记录
}

// TableName 易损件记录 PartRecord自定义表名 PartRecords
func (PartRecord) TableName() string {
	return "PartRecords"
}
