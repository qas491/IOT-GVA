package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(system.Equipment{}, system.ProjectInfo{}, system.PartRecord{}, system.ItemInfo{})
	if err != nil {
		return err
	}
	return nil
}
