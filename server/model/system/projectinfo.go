// 自动生成模板ProjectInfo
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 项目信息 结构体  ProjectInfo
type ProjectInfo struct {
	global.GVA_MODEL
	InstallationLocation *string    `json:"installationLocation" form:"installationLocation" gorm:"column:installation_location;" binding:"required"` //安装地点
	DetailedLocation     *string    `json:"detailedLocation" form:"detailedLocation" gorm:"column:detailed_location;" binding:"required"`             //详细位置
	City                 *string    `json:"city" form:"city" gorm:"column:city;" binding:"required"`                                                  //所在城市
	UpdateDate           *time.Time `json:"updateDate" form:"updateDate" gorm:"column:update_date;" binding:"required"`                               //更新日期
	//UpdateDate  *time.Time `json:"updateDate" form:"updateDate" gorm:"column:update_date;" binding:"required"`  //更新日期
}

// TableName 项目信息 ProjectInfo自定义表名 Projects
func (ProjectInfo) TableName() string {
	//return "Projects"
	return "Projects"
}
