package system

import (
	"bytes"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 初始化测试数据库
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
}
