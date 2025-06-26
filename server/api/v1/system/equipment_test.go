package system

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

/*// 初始化测试数据库
func initTestDB(t *testing.T) {
	var err error
	global.GVA_DB, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	if err := global.GVA_DB.AutoMigrate(&system.Equipment{}); err != nil {
		t.Fatalf("Failed to migrate equipment table: %v", err)
	}
}

// 重置测试数据库
func resetTestDB(t *testing.T) {
	if err := global.GVA_DB.Migrator().DropTable(&system.Equipment{}); err != nil {
		t.Fatalf("Failed to drop equipment table: %v", err)
	}
	if err := global.GVA_DB.AutoMigrate(&system.Equipment{}); err != nil {
		t.Fatalf("Failed to migrate equipment table: %v", err)
	}
}

func TestCreateEquipment(t *testing.T) {
	initTestDB(t)
	defer resetTestDB(t)

	r := gin.Default()
	EQApi := &EquipmentApi{}
	r.POST("/EQ/createEquipment", EQApi.CreateEquipment)

	equipment := system.Equipment{
		EquipmentNumber:     new(string),
		EquipmentName:       new(string),
		Location:            new(string),
		OperationalStatus:   new(string),
		NunningState:        new(string),
		SctiveState:         new(string),
		SoftwareVersion:     new(string),
		AfterSalesPersonnel: new(string),
	}
	*equipment.EquipmentNumber = "12345"
	*equipment.EquipmentName = "Test Equipment"
	*equipment.Location = "Test Location"
	*equipment.OperationalStatus = "1"
	*equipment.NunningState = "1"
	*equipment.SctiveState = "1"
	*equipment.SoftwareVersion = "1.0"
	*equipment.AfterSalesPersonnel = "Test Person"

	jsonData, err := json.Marshal(equipment)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/EQ/createEquipment", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var resp response.Response
	if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if resp.Msg != "创建成功" {
		t.Errorf("handler returned unexpected message: got %v want %v",
			resp.Msg, "创建成功")
	}
}

func TestUpdateEquipmentStatus(t *testing.T) {
	initTestDB(t)
	defer resetTestDB(t)

	r := gin.Default()
	EQApi := &EquipmentApi{}
	r.POST("/EQ/updateEquipmentStatus/:deviceID", EQApi.UpdateEquipmentStatus)

	equipment := system.Equipment{
		EquipmentNumber:     new(string),
		EquipmentName:       new(string),
		Location:            new(string),
		OperationalStatus:   new(string),
		NunningState:        new(string),
		SctiveState:         new(string),
		SoftwareVersion:     new(string),
		AfterSalesPersonnel: new(string),
	}
	*equipment.EquipmentNumber = "12345"
	*equipment.EquipmentName = "Test Equipment"
	*equipment.Location = "Test Location"
	*equipment.OperationalStatus = "1"
	*equipment.NunningState = "1"
	*equipment.SctiveState = "1"
	*equipment.SoftwareVersion = "1.0"
	*equipment.AfterSalesPersonnel = "Test Person"

	if err := global.GVA_DB.Create(&equipment).Error; err != nil {
		t.Fatalf("Failed to create test equipment: %v", err)
	}

	deviceID := strconv.Itoa(int(equipment.ID))
	formData := url.Values{}
	formData.Add("status", "2")

	req, err := http.NewRequest("POST", "/EQ/updateEquipmentStatus/"+deviceID, strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var resp response.Response
	if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if resp.Msg != "设备状态更新请求已发送" {
		t.Errorf("handler returned unexpected message: got %v want %v",
			resp.Msg, "设备状态更新请求已发送")
	}
}
func TestDeleteEquipment(t *testing.T) {
	initTestDB(t)
	defer resetTestDB(t)

	// 创建一个新的 Gin 引擎
	r := gin.Default()
	EQApi := &EquipmentApi{}
	r.DELETE("/EQ/deleteEquipment", EQApi.DeleteEquipment)

	// 插入一个测试设备
	equipment := system.Equipment{

		OperationalStatus: new(string),
	}
	*equipment.OperationalStatus = "1"
	if err := global.GVA_DB.Create(&equipment).Error; err != nil {
		t.Fatalf("Failed to create test equipment: %v", err)
	}

	// 准备删除请求
	deviceID := strconv.Itoa(int(equipment.ID))
	req, err := http.NewRequest("DELETE", "/EQ/deleteEquipment?ID="+deviceID, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	// 检查响应状态码
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// 检查响应消息
	var resp response.Response
	if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if resp.Msg != "删除成功" {
		t.Errorf("handler returned unexpected message: got %v want %v",
			resp.Msg, "删除成功")
	}
}*/

// 测试前初始化数据库
func TestMain(m *testing.M) {
	// 从环境变量获取数据库配置，或使用默认值
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "root:wzydsb@tcp(14.103.195.100:3306)/gva?charset=utf8mb4&parseTime=True&loc=Local"
	}

	// 连接数据库
	var err error
	global.GVA_DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 确保数据库连接正常
	db, err := global.GVA_DB.DB()
	if err != nil {
		log.Fatalf("获取数据库连接失败: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("数据库连接测试失败: %v", err)
	}

	// 自动迁移设备表
	if err := global.GVA_DB.AutoMigrate(&system.Equipment{}); err != nil {
		log.Fatalf("迁移设备表失败: %v", err)
	}

	// 运行测试
	os.Exit(m.Run())
}

// 测试MQTT修改已有设备运营状态
func TestUpdateExistingEquipmentStatus(t *testing.T) {
	// 检查全局数据库连接是否已初始化
	if global.GVA_DB == nil {
		t.Fatalf("数据库连接未初始化")
	}

	// 用于同步消息处理的等待组
	var wg sync.WaitGroup
	wg.Add(1)

	// 创建MQTT客户端选项
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://14.103.149.204:1883")
	opts.SetClientID("test-client-" + fmt.Sprintf("%d", time.Now().UnixNano()))

	// 用于存储更新结果的通道
	updateResult := make(chan error, 1)

	// MQTT消息处理函数
	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		defer wg.Done()

		t.Logf("收到MQTT消息: 主题=%s, 载荷=%s", msg.Topic(), string(msg.Payload()))

		// 解析消息格式
		payload := string(msg.Payload())
		parts := strings.Split(payload, ":")
		if len(parts) != 2 {
			err := fmt.Errorf("无效的消息格式: %s", payload)
			updateResult <- err
			return
		}

		deviceID, err := strconv.Atoi(parts[0])
		if err != nil {
			err := fmt.Errorf("解析设备ID失败: %v", err)
			updateResult <- err
			return
		}

		// 更新设备状态
		err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			var equipment system.Equipment
			if err := tx.First(&equipment, deviceID).Error; err != nil {
				return fmt.Errorf("查找设备失败: ID=%d, 错误=%v", deviceID, err)
			}

			// 打印更新前的状态
			t.Logf("更新前设备状态: ID=%d, 状态=%s", deviceID, *equipment.OperationalStatus)

			// 更新设备状态字段
			result := tx.Model(&equipment).Update("operational_status", parts[1])
			if result.Error != nil {
				return fmt.Errorf("更新设备状态失败: %v", result.Error)
			}

			if result.RowsAffected == 0 {
				return fmt.Errorf("未更新任何记录: ID=%d", deviceID)
			}

			// 刷新设备数据，获取最新状态
			if err := tx.First(&equipment, deviceID).Error; err != nil {
				return fmt.Errorf("刷新设备数据失败: ID=%d, 错误=%v", deviceID, err)
			}

			t.Logf("更新后设备状态: ID=%d, 状态=%s", deviceID, *equipment.OperationalStatus)
			return nil
		})

		if err != nil {
			t.Logf("更新设备状态事务失败: %v", err)
			updateResult <- err
			return
		}

		updateResult <- nil // 更新成功
	}

	opts.SetDefaultPublishHandler(messageHandler)

	// 连接MQTT服务器
	global.GVA_MQTT_CLIENT = mqtt.NewClient(opts)
	if token := global.GVA_MQTT_CLIENT.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("连接MQTT服务器失败: %v", token.Error())
	}

	// 订阅设备状态主题
	if token := global.GVA_MQTT_CLIENT.Subscribe("equipment/status", 0, nil); token.Wait() && token.Error() != nil {
		t.Fatalf("订阅MQTT主题失败: %v", token.Error())
	}

	// 等待订阅完成
	time.Sleep(2 * time.Second)

	// 选择一个现有设备进行测试（假设设备ID为1存在）
	deviceID := 10
	t.Logf("准备测试修改现有设备，ID: %d", deviceID)

	// 验证设备存在并获取当前状态
	var existingEquipment system.Equipment
	if err := global.GVA_DB.First(&existingEquipment, deviceID).Error; err != nil {
		t.Fatalf("查找现有设备失败: ID=%d, 错误=%v", deviceID, err)
	}

	currentStatus := *existingEquipment.OperationalStatus
	t.Logf("现有设备状态: ID=%d, 当前状态=%s", deviceID, currentStatus)

	// 选择一个与当前状态不同的新状态
	newStatus := "3"
	if currentStatus == newStatus {
		newStatus = "2" // 如果当前状态已经是3，则使用2
	}

	// 发布MQTT消息更新设备状态
	topic := "equipment/status"
	message := fmt.Sprintf("%d:%s", deviceID, newStatus)

	t.Logf("发布MQTT消息: 主题=%s, 内容=%s", topic, message)
	token := global.GVA_MQTT_CLIENT.Publish(topic, 0, false, message)
	token.Wait()

	if token.Error() != nil {
		t.Fatalf("发布MQTT消息失败: %v", token.Error())
	}

	// 等待消息处理完成或超时
	ch := make(chan struct{})
	go func() {
		wg.Wait()
		close(ch)
	}()

	select {
	case <-ch:
		// 检查更新结果
		if err := <-updateResult; err != nil {
			t.Fatalf("MQTT消息处理失败: %v", err)
		}
		t.Log("MQTT消息处理完成")
	case <-time.After(5 * time.Second):
		t.Fatalf("等待MQTT消息处理超时")
	}

	// 验证数据库中的最终状态
	var updatedEquipment system.Equipment
	if err := global.GVA_DB.First(&updatedEquipment, deviceID).Error; err != nil {
		t.Fatalf("查询更新后的设备失败: ID=%d, 错误=%v", deviceID, err)
	}

	if *updatedEquipment.OperationalStatus != newStatus {
		t.Fatalf("设备状态更新失败，期望 %s，实际 %s",
			newStatus, *updatedEquipment.OperationalStatus)
	}

	t.Logf("设备状态已成功从 %s 更新为 %s", currentStatus, newStatus)

	// 注意：不再清理测试数据，因为我们使用的是现有设备
	defer global.GVA_MQTT_CLIENT.Disconnect(250)
}
