// 自动生成模板Equipment
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 设备信息 结构体  Equipment
type Equipment struct {
	global.GVA_MODEL
	EquipmentNumber     *string `json:"equipmentNumber" form:"equipmentNumber" gorm:"column:equipment_number;" binding:"required"`           //设备编号
	EquipmentName       *string `json:"equipmentName" form:"equipmentName" gorm:"column:equipment_name;" binding:"required"`                 //设备名称
	Location            *string `json:"location" form:"location" gorm:"column:location;" binding:"required"`                                 //所在位置
	OperationalStatus   *string `json:"operationalStatus" form:"operationalStatus" gorm:"default:1;comment:运营状态;column:operational_status;"` //运营状态
	NunningState        *string `json:"nunningState" form:"nunningState" gorm:"default:1;comment:运行状态;column:nunning_state;"`                //运行状态
	SctiveState         *string `json:"sctiveState" form:"sctiveState" gorm:"default:1;comment:激活状态;column:sctive_state;"`                   //激活状态
	SoftwareVersion     *string `json:"softwareVersion" form:"softwareVersion" gorm:"column:software_version;" binding:"required"`           //软件版本
	AfterSalesPersonnel *string `json:"afterSalesPersonnel" form:"afterSalesPersonnel" gorm:"comment:售后人员;column:after_sales_personnel;"`    //售后人员
	LastSeen            *int64  `json:"lastSeen" form:"lastSeen" gorm:"column:last_seen;comment:最后心跳时间（时间戳）"`                                // 最后心跳时间
	Online              *bool   `json:"online" form:"online" gorm:"column:online;comment:是否在线"`                                              // 在线状态
}

// TableName 设备信息 Equipment自定义表名 Equipment
func (Equipment) TableName() string {
	return "Equipment"
}
