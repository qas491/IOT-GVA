// 自动生成模板ItemInfo
package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// 易损件 结构体  ItemInfo
type ItemInfo struct {
	global.GVA_MODEL
	Code          *string `json:"code" form:"code" gorm:"comment:标识每条记录的唯一编号;column:code;" binding:"required"`                           //编码
	Name          *string `json:"name" form:"name" gorm:"comment:表示易损件的名称;column:name;" binding:"required"`                              //名称
	Specification *string `json:"specification" form:"specification" gorm:"comment:表示易损件的具体规格;column:specification;" binding:"required"` //规格
	Days          *int    `json:"days" form:"days" gorm:"comment:表示使用该易损件的天数;column:days;" binding:"required"`                           //周期
	UsageCount    *int    `json:"usageCount" form:"usageCount" gorm:"comment:表示该易损件的使用总次数;column:usage_count;" binding:"required"`       //使用次数
	//UsageCount    *int    `json:"usageCount" form:"usageCount" gorm:"comment:表示该易损件的使用总次数;column:usage_count;" binding:"required"`       //使用次数
}

// TableName 易损件 ItemInfo自定义表名 ItemInfo
func (ItemInfo) TableName() string {
	return "ItemInfo"

}
