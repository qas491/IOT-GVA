package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type EquipmentService struct{}

// CreateEquipment 创建设备信息记录
// Author [yourname](https://github.com/yourname)
func (EQService *EquipmentService) CreateEquipment(ctx context.Context, EQ *system.Equipment) (err error) {
	err = global.GVA_DB.Create(EQ).Error
	return err
}

// DeleteEquipment 删除设备信息记录
// Author [yourname](https://github.com/yourname)
func (EQService *EquipmentService) DeleteEquipment(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&system.Equipment{}, "id = ?", ID).Error
	return err
}

// DeleteEquipmentByIds 批量删除设备信息记录
// Author [yourname](https://github.com/yourname)
func (EQService *EquipmentService) DeleteEquipmentByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.Equipment{}, "id in ?", IDs).Error
	return err
}

// UpdateEquipment 更新设备信息记录
// Author [yourname](https://github.com/yourname)
func (EQService *EquipmentService) UpdateEquipment(ctx context.Context, EQ system.Equipment) (err error) {
	err = global.GVA_DB.Model(&system.Equipment{}).Where("id = ?", EQ.ID).Updates(&EQ).Error
	return err
}

// GetEquipment 根据ID获取设备信息记录
// Author [yourname](https://github.com/yourname)
func (EQService *EquipmentService) GetEquipment(ctx context.Context, ID string) (EQ system.Equipment, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&EQ).Error
	return
}

// GetEquipmentInfoList 分页获取设备信息记录
// Author [yourname](https://github.com/yourname)
func (EQService *EquipmentService) GetEquipmentInfoList(ctx context.Context, info systemReq.EquipmentSearch) (list []system.Equipment, total int64, err error) {
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
		db = db.Where("equipment_number LIKE ?", "%"+*info.EquipmentNumber+"%")
	}
	if info.EquipmentName != nil && *info.EquipmentName != "" {
		db = db.Where("equipment_name LIKE ?", "%"+*info.EquipmentName+"%")
	}
	if info.OperationalStatus != nil && *info.OperationalStatus != "" {
		db = db.Where("operational_status LIKE ?", "%"+*info.OperationalStatus+"%")
	}
	if info.NunningState != nil && *info.NunningState != "" {
		db = db.Where("nunning_state LIKE ?", "%"+*info.NunningState+"%")
	}
	if info.SctiveState != nil && *info.SctiveState != "" {
		db = db.Where("sctive_state LIKE ?", "%"+*info.SctiveState+"%")
	}
	if info.SoftwareVersion != nil && *info.SoftwareVersion != "" {
		db = db.Where("software_version LIKE ?", "%"+*info.SoftwareVersion+"%")
	}
	if info.AfterSalesPersonnel != nil && *info.AfterSalesPersonnel != "" {
		db = db.Where("after_sales_personnel LIKE ?", "%"+*info.AfterSalesPersonnel+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&EQs).Error
	return EQs, total, err
}
func (EQService *EquipmentService) GetEquipmentPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// QueryDeviceCount 查询后端设备数量
// Author [yourname](https://github.com/yourname)
func (EQService *EquipmentService) QueryDeviceCount(ctx context.Context) (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&system.Equipment{})
	return db.Error
}

func (EQService *EquipmentService) QueryTotalDeviceCount(ctx context.Context) (int64, error) {
	var total int64
	err := global.GVA_DB.Model(&system.Equipment{}).Count(&total).Error
	return total, err
}

// QueryDeviceCountByStatus 根据运营状态查询设备数量
func (EQService *EquipmentService) QueryDeviceCountByStatus(ctx context.Context, status string) (int64, error) {
	var total int64
	err := global.GVA_DB.Model(&system.Equipment{}).Where("operational_status = ?", status).Count(&total).Error
	return total, err
}

// GetDashboardStats 获取仪表盘统计数据
func (EQService *EquipmentService) GetDashboardStats(ctx context.Context) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// 1. 设备状态统计
	var total, inUse, idle, stopped int64

	// 总设备数
	err := global.GVA_DB.Model(&system.Equipment{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	// 在用设备数 (operational_status = "1")
	err = global.GVA_DB.Model(&system.Equipment{}).Where("operational_status = ?", "1").Count(&inUse).Error
	if err != nil {
		return nil, err
	}

	// 闲置设备数 (operational_status = "2")
	err = global.GVA_DB.Model(&system.Equipment{}).Where("operational_status = ?", "2").Count(&idle).Error
	if err != nil {
		return nil, err
	}

	// 停用设备数 (operational_status = "3")
	err = global.GVA_DB.Model(&system.Equipment{}).Where("operational_status = ?", "3").Count(&stopped).Error
	if err != nil {
		return nil, err
	}

	stats["deviceStats"] = map[string]int64{
		"total":   total,
		"inUse":   inUse,
		"idle":    idle,
		"stopped": stopped,
	}

	// 2. 激活情况统计（在用设备中）
	var activated, unactivated int64

	// 已激活设备数 (sctive_state = "1" 且 operational_status = "1")
	err = global.GVA_DB.Model(&system.Equipment{}).Where("sctive_state = ? AND operational_status = ?", "1", "1").Count(&activated).Error
	if err != nil {
		return nil, err
	}

	// 未激活设备数 (sctive_state = "2" 且 operational_status = "1")
	err = global.GVA_DB.Model(&system.Equipment{}).Where("sctive_state = ? AND operational_status = ?", "2", "1").Count(&unactivated).Error
	if err != nil {
		return nil, err
	}

	stats["activationStats"] = map[string]int64{
		"activated":   activated,
		"unactivated": unactivated,
	}

	// 3. 运行状态统计（已激活设备中）
	var normal, fault int64

	// 正常设备数 (nunning_state = "1" 且 sctive_state = "1")
	err = global.GVA_DB.Model(&system.Equipment{}).Where("nunning_state = ? AND sctive_state = ?", "1", "1").Count(&normal).Error
	if err != nil {
		return nil, err
	}

	// 故障设备数 (nunning_state = "2" 且 sctive_state = "1")
	err = global.GVA_DB.Model(&system.Equipment{}).Where("nunning_state = ? AND sctive_state = ?", "2", "1").Count(&fault).Error
	if err != nil {
		return nil, err
	}

	stats["runningStats"] = map[string]int64{
		"normal": normal,
		"fault":  fault,
	}

	// 4. 待办事项统计
	var todoActivate, todoInstall, todoModel, todoLocation int64

	// 待激活设备数 (sctive_state = "2")
	err = global.GVA_DB.Model(&system.Equipment{}).Where("sctive_state = ?", "2").Count(&todoActivate).Error
	if err != nil {
		return nil, err
	}

	// 待安装设备数 (location为空或null)
	err = global.GVA_DB.Model(&system.Equipment{}).Where("location IS NULL OR location = ''").Count(&todoInstall).Error
	if err != nil {
		return nil, err
	}

	// 型号补充 (equipment_name为空或null)
	err = global.GVA_DB.Model(&system.Equipment{}).Where("equipment_name IS NULL OR equipment_name = ''").Count(&todoModel).Error
	if err != nil {
		return nil, err
	}

	// 安装地点补充 (location为空或null)
	err = global.GVA_DB.Model(&system.Equipment{}).Where("location IS NULL OR location = ''").Count(&todoLocation).Error
	if err != nil {
		return nil, err
	}

	stats["todoItems"] = map[string]int64{
		"activate": todoActivate,
		"install":  todoInstall,
		"model":    todoModel,
		"location": todoLocation,
	}

	// 5. 设备类型排行统计
	var equipmentTypes []struct {
		EquipmentName string `json:"equipmentName"`
		Count         int64  `json:"count"`
	}

	err = global.GVA_DB.Model(&system.Equipment{}).
		Select("equipment_name, COUNT(*) as count").
		Where("operational_status = ?", "1").
		Group("equipment_name").
		Order("count DESC").
		Limit(10).
		Scan(&equipmentTypes).Error
	if err != nil {
		return nil, err
	}

	// 转换为前端需要的格式
	var rankingData []map[string]interface{}
	for _, et := range equipmentTypes {
		rankingData = append(rankingData, map[string]interface{}{
			"name":  et.EquipmentName,
			"count": et.Count,
		})
	}

	stats["rankingData"] = rankingData

	// 6. 地图设备数据
	var mapDevices []system.Equipment
	err = global.GVA_DB.Model(&system.Equipment{}).
		Select("id, equipment_name, equipment_number, location, operational_status, nunning_state, sctive_state, online").
		Where("location IS NOT NULL AND location != ''").
		Find(&mapDevices).Error
	if err != nil {
		return nil, err
	}

	// 转换为地图需要的格式
	var mapData []map[string]interface{}
	for _, device := range mapDevices {
		// 这里需要根据location解析经纬度，暂时使用模拟数据
		// 实际项目中可能需要添加经纬度字段或使用地理编码服务
		mapData = append(mapData, map[string]interface{}{
			"id":       device.ID,
			"name":     device.EquipmentName,
			"type":     device.EquipmentNumber,
			"status":   getStatusText(device.OperationalStatus, device.NunningState),
			"location": device.Location,
			"online":   device.Online,
		})
	}

	stats["mapData"] = mapData

	return stats, nil
}

// getStatusText 获取状态文本
func getStatusText(operationalStatus, runningState *string) string {
	if operationalStatus == nil {
		return "未知"
	}

	switch *operationalStatus {
	case "1":
		if runningState != nil && *runningState == "2" {
			return "故障"
		}
		return "在线"
	case "2":
		return "离线"
	case "3":
		return "停用"
	default:
		return "未知"
	}
}
