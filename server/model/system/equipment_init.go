package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

func InitEquipment() {
	equipments := []Equipment{
		{
			EquipmentNumber:     stringPtr("EQ001"),
			EquipmentName:       stringPtr("光刻机"),
			Location:            stringPtr("北京市朝阳区"),
			OperationalStatus:   stringPtr("1"),
			NunningState:        stringPtr("1"),
			SctiveState:         stringPtr("1"),
			SoftwareVersion:     stringPtr("v2.1.0"),
			AfterSalesPersonnel: stringPtr("张三"),
			Online:              boolPtr(true),
		},
		{
			EquipmentNumber:     stringPtr("EQ002"),
			EquipmentName:       stringPtr("刻蚀机"),
			Location:            stringPtr("上海市浦东新区"),
			OperationalStatus:   stringPtr("1"),
			NunningState:        stringPtr("1"),
			SctiveState:         stringPtr("1"),
			SoftwareVersion:     stringPtr("v1.8.5"),
			AfterSalesPersonnel: stringPtr("李四"),
			Online:              boolPtr(true),
		},
		{
			EquipmentNumber:     stringPtr("EQ003"),
			EquipmentName:       stringPtr("薄膜沉积设备"),
			Location:            stringPtr("深圳市南山区"),
			OperationalStatus:   stringPtr("1"),
			NunningState:        stringPtr("2"),
			SctiveState:         stringPtr("1"),
			SoftwareVersion:     stringPtr("v2.0.1"),
			AfterSalesPersonnel: stringPtr("王五"),
			Online:              boolPtr(false),
		},
		{
			EquipmentNumber:     stringPtr("EQ004"),
			EquipmentName:       stringPtr("离子注入机"),
			Location:            stringPtr("广州市天河区"),
			OperationalStatus:   stringPtr("2"),
			NunningState:        stringPtr("1"),
			SctiveState:         stringPtr("2"),
			SoftwareVersion:     stringPtr("v1.9.2"),
			AfterSalesPersonnel: stringPtr("赵六"),
			Online:              boolPtr(false),
		},
		{
			EquipmentNumber:     stringPtr("EQ005"),
			EquipmentName:       stringPtr("机械抛光"),
			Location:            stringPtr("成都市高新区"),
			OperationalStatus:   stringPtr("1"),
			NunningState:        stringPtr("1"),
			SctiveState:         stringPtr("1"),
			SoftwareVersion:     stringPtr("v1.7.8"),
			AfterSalesPersonnel: stringPtr("钱七"),
			Online:              boolPtr(true),
		},
		{
			EquipmentNumber:     stringPtr("EQ006"),
			EquipmentName:       stringPtr("清洗机"),
			Location:            stringPtr("杭州市西湖区"),
			OperationalStatus:   stringPtr("1"),
			NunningState:        stringPtr("1"),
			SctiveState:         stringPtr("1"),
			SoftwareVersion:     stringPtr("v2.2.0"),
			AfterSalesPersonnel: stringPtr("孙八"),
			Online:              boolPtr(true),
		},
		{
			EquipmentNumber:     stringPtr("EQ007"),
			EquipmentName:       stringPtr("氧化炉"),
			Location:            stringPtr("南京市江宁区"),
			OperationalStatus:   stringPtr("3"),
			NunningState:        stringPtr("1"),
			SctiveState:         stringPtr("1"),
			SoftwareVersion:     stringPtr("v1.6.4"),
			AfterSalesPersonnel: stringPtr("周九"),
			Online:              boolPtr(false),
		},
		{
			EquipmentNumber:     stringPtr("EQ008"),
			EquipmentName:       stringPtr("晶圆切割机"),
			Location:            stringPtr("武汉市东湖高新区"),
			OperationalStatus:   stringPtr("1"),
			NunningState:        stringPtr("2"),
			SctiveState:         stringPtr("1"),
			SoftwareVersion:     stringPtr("v2.0.5"),
			AfterSalesPersonnel: stringPtr("吴十"),
			Online:              boolPtr(false),
		},
		{
			EquipmentNumber:     stringPtr("EQ009"),
			EquipmentName:       stringPtr("晶元键合机"),
			Location:            stringPtr("重庆市渝北区"),
			OperationalStatus:   stringPtr("1"),
			NunningState:        stringPtr("1"),
			SctiveState:         stringPtr("2"),
			SoftwareVersion:     stringPtr("v1.8.9"),
			AfterSalesPersonnel: stringPtr("郑十一"),
			Online:              boolPtr(false),
		},
		{
			EquipmentNumber:     stringPtr("EQ010"),
			EquipmentName:       stringPtr("测试机"),
			Location:            stringPtr("西安市高新区"),
			OperationalStatus:   stringPtr("1"),
			NunningState:        stringPtr("1"),
			SctiveState:         stringPtr("1"),
			SoftwareVersion:     stringPtr("v2.1.3"),
			AfterSalesPersonnel: stringPtr("王十二"),
			Online:              boolPtr(true),
		},
		{
			EquipmentNumber:     stringPtr("EQ011"),
			EquipmentName:       stringPtr("光刻机"),
			Location:            stringPtr("北京市海淀区"),
			OperationalStatus:   stringPtr("1"),
			NunningState:        stringPtr("1"),
			SctiveState:         stringPtr("1"),
			SoftwareVersion:     stringPtr("v2.1.0"),
			AfterSalesPersonnel: stringPtr("张三"),
			Online:              boolPtr(true),
		},
		{
			EquipmentNumber:     stringPtr("EQ012"),
			EquipmentName:       stringPtr("清洗机"),
			Location:            stringPtr("上海市黄浦区"),
			OperationalStatus:   stringPtr("1"),
			NunningState:        stringPtr("1"),
			SctiveState:         stringPtr("1"),
			SoftwareVersion:     stringPtr("v2.2.0"),
			AfterSalesPersonnel: stringPtr("孙八"),
			Online:              boolPtr(true),
		},
		{
			EquipmentNumber:     stringPtr("EQ013"),
			EquipmentName:       stringPtr("机械抛光"),
			Location:            stringPtr("深圳市福田区"),
			OperationalStatus:   stringPtr("1"),
			NunningState:        stringPtr("1"),
			SctiveState:         stringPtr("1"),
			SoftwareVersion:     stringPtr("v1.7.8"),
			AfterSalesPersonnel: stringPtr("钱七"),
			Online:              boolPtr(true),
		},
		{
			EquipmentNumber:     stringPtr("EQ014"),
			EquipmentName:       stringPtr("晶元键合机"),
			Location:            stringPtr("广州市越秀区"),
			OperationalStatus:   stringPtr("1"),
			NunningState:        stringPtr("1"),
			SctiveState:         stringPtr("1"),
			SoftwareVersion:     stringPtr("v1.8.9"),
			AfterSalesPersonnel: stringPtr("郑十一"),
			Online:              boolPtr(true),
		},
		{
			EquipmentNumber:     stringPtr("EQ015"),
			EquipmentName:       stringPtr("测试机"),
			Location:            stringPtr("成都市锦江区"),
			OperationalStatus:   stringPtr("1"),
			NunningState:        stringPtr("1"),
			SctiveState:         stringPtr("1"),
			SoftwareVersion:     stringPtr("v2.1.3"),
			AfterSalesPersonnel: stringPtr("王十二"),
			Online:              boolPtr(true),
		},
	}

	for _, equipment := range equipments {
		var existingEquipment Equipment
		err := global.GVA_DB.Where("equipment_number = ?", equipment.EquipmentNumber).First(&existingEquipment).Error
		if err != nil {
			// 设备不存在，创建新设备
			err = global.GVA_DB.Create(&equipment).Error
			if err != nil {
				global.GVA_LOG.Error("创建设备失败", zap.Error(err))
			}
		}
	}
}

func stringPtr(s string) *string { return &s }
func boolPtr(b bool) *bool       { return &b }

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
