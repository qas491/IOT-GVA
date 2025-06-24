package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"strconv"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 初始化数据库
func initDB() {
	var err error
	global.GVA_DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	// 自动迁移设备表
	global.GVA_DB.AutoMigrate(&system.Equipment{})
}

// 测试 MQTT 实时修改设备状态的效果
func TestMQTTRealTimeUpdate(t *testing.T) {
	// 初始化数据库
	initDB()

	// 初始化 MQTT 客户端
	InitMQTT()

	// 等待一段时间确保订阅成功
	time.Sleep(2 * time.Second)

	// 创建设备
	deviceID := 4
	status := 1
	statusStr := strconv.Itoa(status)
	EQ := system.Equipment{
		GVA_MODEL: global.GVA_MODEL{
			ID: uint(deviceID),
		},
		OperationalStatus: &statusStr,
	}
	err := global.GVA_DB.Create(&EQ).Error
	if err != nil {
		t.Fatalf("创建设备失败: %v", err)
	}

	// 验证设备信息是否成功插入
	var insertedEQ system.Equipment
	err = global.GVA_DB.Where("id = ?", deviceID).First(&insertedEQ).Error
	if err != nil {
		if err.Error() == "record not found" {
			t.Errorf("未找到设备记录，可能是数据库操作失败，请检查数据库连接和数据插入操作")
		} else {
			t.Errorf("验证设备信息插入失败: %v", err)
		}
		return
	}

	defer func() {
		// 清理测试数据
		global.GVA_DB.Delete(&system.Equipment{}, "id = ?", deviceID)
	}()

	// 模拟发布 MQTT 消息更新设备状态
	newStatus := 2
	newStatusStr := strconv.Itoa(newStatus)
	topic := fmt.Sprintf("equipment/status")
	message := fmt.Sprintf("%d:%s", deviceID, newStatusStr)
	token := global.GVA_MQTT_CLIENT.Publish(topic, 0, false, message)
	token.Wait()

	if token.Error() != nil {
		t.Errorf("发布 MQTT 消息失败: %v", token.Error())
		return
	}

	// 等待一段时间确保消息被处理
	time.Sleep(2 * time.Second)

	// 验证设备状态是否更新
	var updatedEQ system.Equipment
	err = global.GVA_DB.Where("id = ?", deviceID).First(&updatedEQ).Error
	if err != nil {
		if err.Error() == "record not found" {
			t.Errorf("未找到设备记录，可能是数据库操作失败，请检查数据库连接和数据插入操作")
		} else {
			t.Errorf("查询设备信息失败: %v", err)
		}
		return
	}

	if *updatedEQ.OperationalStatus != newStatusStr {
		t.Errorf("设备运营状态更新失败，期望 %s，实际 %s", newStatusStr, *updatedEQ.OperationalStatus)
	}
}
