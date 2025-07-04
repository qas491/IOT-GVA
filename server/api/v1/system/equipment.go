package system

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EquipmentApi struct{}

// 定义 MQTT 消息处理函数
var mqttHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	message := string(msg.Payload())
	global.GVA_LOG.Info("收到MQTT消息", zap.String("payload", message))
	parts := strings.Split(message, ":")
	if len(parts) < 2 {
		global.GVA_LOG.Error("MQTT 消息格式错误", zap.String("message", message))
		return
	}

	// 设备ID为整型
	deviceIDStr := parts[0]
	statusStr := parts[1]
	isLWT := len(parts) > 2 && parts[2] == "LWT"

	// 设备ID转为整型
	deviceID, err := strconv.ParseUint(deviceIDStr, 10, 64)
	if err != nil {
		global.GVA_LOG.Error("设备ID转换错误", zap.String("deviceID", deviceIDStr), zap.Error(err))
		return
	}

	// 校验状态码
	if statusStr != "1" && statusStr != "2" && statusStr != "3" {
		global.GVA_LOG.Error("运营状态非法", zap.String("statusStr", statusStr))
		return
	}

	// 查询设备
	var EQ system.Equipment
	err = global.GVA_DB.Where("id = ?", deviceID).First(&EQ).Error
	if err != nil {
		global.GVA_LOG.Error("查询设备信息失败", zap.Uint("deviceID", uint(deviceID)), zap.Error(err))
		return
	}

	// 更新运营状态、心跳时间、在线状态
	currentTime := time.Now().Unix()
	online := true

	err = global.GVA_DB.Model(&system.Equipment{}).Where("id = ?", deviceID).Updates(map[string]interface{}{
		"operational_status": statusStr,
		"last_seen":          currentTime,
		"online":             online,
	}).Error
	if err != nil {
		global.GVA_LOG.Error("更新设备运营状态失败", zap.Uint("deviceID", uint(deviceID)), zap.Error(err))
		return
	}

	global.GVA_LOG.Info("设备运营状态和心跳时间更新成功", zap.Uint("deviceID", uint(deviceID)), zap.String("status", statusStr))

	// LWT处理（如需特殊处理断开）
	if isLWT {
		offline := false
		lwtStatus := "2" // 断开时可设为"2"（闲置）或自定义
		err := global.GVA_DB.Model(&system.Equipment{}).Where("id = ?", deviceID).Updates(map[string]interface{}{
			"online":             offline,
			"operational_status": lwtStatus,
		}).Error
		if err != nil {
			global.GVA_LOG.Error("LWT自动断开设备失败", zap.Uint("deviceID", uint(deviceID)), zap.Error(err))
		} else {
			global.GVA_LOG.Info("LWT设备超时自动断开", zap.Uint("deviceID", uint(deviceID)))
		}
	}
}

// 新增：定时任务，定期检查设备心跳，超时未收到心跳则自动断开
func StartEquipmentHeartbeatChecker(ctx context.Context, timeoutSeconds int64, checkInterval time.Duration) {
	go func() {
		ticker := time.NewTicker(checkInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				var equipments []system.Equipment
				err := global.GVA_DB.Find(&equipments).Error
				if err != nil {
					global.GVA_LOG.Error("定时检查设备心跳失败", zap.Error(err))
					continue
				}
				now := time.Now().Unix()
				for _, eq := range equipments {
					if eq.LastSeen == nil || eq.Online == nil || *eq.Online == false {
						continue
					}
					if now-(*eq.LastSeen) > timeoutSeconds {
						offline := false
						operationalStatus := "2" // 2=断开/故障
						err := global.GVA_DB.Model(&system.Equipment{}).Where("id = ?", eq.ID).Updates(map[string]interface{}{
							"online":             offline,
							"operational_status": operationalStatus,
						}).Error
						if err != nil {
							global.GVA_LOG.Error("自动断开设备失败", zap.Uint("deviceID", eq.ID), zap.Error(err))
						} else {
							global.GVA_LOG.Info("设备超时自动断开", zap.Uint("deviceID", eq.ID))
						}
					}
				}
			}
		}
	}()
}

// 初始化 MQTT 客户端
func InitMQTT() {
	// 配置 MQTT 客户端
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://14.103.149.204:1883") // 替换为实际的 MQTT 代理地址
	opts.SetClientID("equipment-api-client")
	opts.SetDefaultPublishHandler(mqttHandler)

	// 创建 MQTT 客户端
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		global.GVA_LOG.Fatal("连接 MQTT 代理失败", zap.Error(token.Error()))
	}

	// 订阅 MQTT 主题
	if token := client.Subscribe("equipment/status", 0, nil); token.Wait() && token.Error() != nil {
		global.GVA_LOG.Fatal("订阅 MQTT 主题失败", zap.Error(token.Error()))
	}

	global.GVA_MQTT_CLIENT = client // 假设在 global 包中定义了 GVA_MQTT_CLIENT 变量
	// 启动心跳检查定时任务
	ctx := context.Background()
	StartEquipmentHeartbeatChecker(ctx, 300, 60*time.Second) // 5分钟无心跳判定为断开，每60秒检查一次
}

// CreateEquipment 创建设备信息
// @Tags Equipment
// @Summary 创建设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.Equipment true "创建设备信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /EQ/createEquipment [post]
func (EQApi *EquipmentApi) CreateEquipment(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var EQ system.Equipment
	err := c.ShouldBindJSON(&EQ)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = EQService.CreateEquipment(ctx, &EQ)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (EQApi *EquipmentApi) UpdateEquipmentStatus(c *gin.Context) {
	deviceID := c.Param("deviceID")
	if deviceID == "" {
		response.FailWithMessage("缺少设备ID", c)
		return
	}

	statusStr := c.PostForm("status")
	if statusStr == "" {
		response.FailWithMessage("缺少状态参数", c)
		return
	}

	// 验证状态值是否合法
	status, err := strconv.Atoi(statusStr)
	if err != nil || (status != 1 && status != 2 && status != 3) {
		response.FailWithMessage("无效的状态值(必须是1, 2, 3)", c)
		return
	}

	// 构建MQTT消息
	topic := fmt.Sprintf("equipment/status")
	message := fmt.Sprintf("%s:%s", deviceID, statusStr)

	// 确保 MQTT 客户端已初始化
	if global.GVA_MQTT_CLIENT == nil {
		InitMQTT()
		time.Sleep(2 * time.Second) // 等待连接和订阅完成
	}

	// 发布MQTT消息
	token := global.GVA_MQTT_CLIENT.Publish(topic, 0, false, message)
	token.Wait()

	if token.Error() != nil {
		global.GVA_LOG.Error("发布MQTT消息失败", zap.Error(token.Error()))
		response.FailWithMessage("更新请求发送失败", c)
		return
	}

	response.OkWithMessage("设备状态更新请求已发送", c)
}

// DeleteEquipment 删除设备信息
// @Tags Equipment
// @Summary 删除设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.Equipment true "删除设备信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /EQ/deleteEquipment [delete]
func (EQApi *EquipmentApi) DeleteEquipment(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := EQService.DeleteEquipment(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEquipmentByIds 批量删除设备信息
// @Tags Equipment
// @Summary 批量删除设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /EQ/deleteEquipmentByIds [delete]
func (EQApi *EquipmentApi) DeleteEquipmentByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := EQService.DeleteEquipmentByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEquipment 更新设备信息
// @Tags Equipment
// @Summary 更新设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.Equipment true "更新设备信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /EQ/updateEquipment [put]
func (EQApi *EquipmentApi) UpdateEquipment(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var EQ system.Equipment
	err := c.ShouldBindJSON(&EQ)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = EQService.UpdateEquipment(ctx, EQ)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEquipment 用id查询设备信息
// @Tags Equipment
// @Summary 用id查询设备信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询设备信息"
// @Success 200 {object} response.Response{data=system.Equipment,msg=string} "查询成功"
// @Router /EQ/findEquipment [get]
func (EQApi *EquipmentApi) FindEquipment(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reEQ, err := EQService.GetEquipment(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reEQ, c)
}

// GetEquipmentList 分页获取设备信息列表
// @Tags Equipment
// @Summary 分页获取设备信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.EquipmentSearch true "分页获取设备信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /EQ/getEquipmentList [get]
func (EQApi *EquipmentApi) GetEquipmentList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo systemReq.EquipmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := EQService.GetEquipmentInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetEquipmentPublic 不需要鉴权的设备信息接口
// @Tags Equipment
// @Summary 不需要鉴权的设备信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EQ/getEquipmentPublic [get]
func (EQApi *EquipmentApi) GetEquipmentPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	EQService.GetEquipmentPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的设备信息接口信息",
	}, "获取成功", c)
}

// QueryDeviceCount 查询后端设备数量
// @Tags Equipment
// @Summary 查询后端设备数量
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.EquipmentSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /EQ/devicecount [GET]
func (EQApi *EquipmentApi) QueryDeviceCount(c *gin.Context) {
	ctx := c.Request.Context()
	total, err := EQService.QueryTotalDeviceCount(ctx)
	if err != nil {
		global.GVA_LOG.Error("查询总设备数量失败!", zap.Error(err))
		response.FailWithMessage("查询总设备数量失败", c)
		return
	}
	response.OkWithData(total, c)
}

// QueryDeviceCountByStatus  根据运营状态查询设备数量
// @Tags Equipment
// @Summary 根据运营状态查询设备数量
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.EquipmentSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /EQ/queryDeviceCountByStatus  [GET]
func (EQApi *EquipmentApi) QueryDeviceCountByStatus(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()
	status := c.Query("status")
	if status == "" {
		response.FailWithMessage("缺少运营状态参数", c)
		return
	}
	total, err := EQService.QueryDeviceCountByStatus(ctx, status)
	if err != nil {
		global.GVA_LOG.Error("根据运营状态查询设备数量失败!", zap.Error(err))
		response.FailWithMessage("根据运营状态查询设备数量失败", c)
		return
	}
	response.OkWithData(total, c)
}

// Dashboard 获取仪表盘统计数据
// @Tags Equipment
// @Summary 获取仪表盘统计数据
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EQ/Dashboard [GET]
func (EQApi *EquipmentApi) Dashboard(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	stats, err := EQService.GetDashboardStats(ctx)
	if err != nil {
		global.GVA_LOG.Error("获取仪表盘统计数据失败!", zap.Error(err))
		response.FailWithMessage("获取仪表盘统计数据失败", c)
		return
	}

	response.OkWithData(stats, c)
}

// ActivateEquipment 批量激活设备
// @Tags Equipment
// @Summary 批量激活设备
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body struct{IDs []uint `json:"IDs"`} true "设备ID列表"
// @Success 200 {object} response.Response{msg=string} "激活成功"
// @Router /equipment/activate [post]
func (EQApi *EquipmentApi) ActivateEquipment(c *gin.Context) {
	var req struct {
		IDs []uint `json:"IDs"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if len(req.IDs) == 0 {
		response.FailWithMessage("请选择要激活的设备", c)
		return
	}

	// 假设"已激活"状态值为'1'
	alreadyActivated := make([]uint, 0)
	toActivate := make([]uint, 0)

	for _, id := range req.IDs {
		var eq system.Equipment
		err := global.GVA_DB.Where("id = ?", id).First(&eq).Error
		if err != nil {
			global.GVA_LOG.Error("查询设备信息失败", zap.Uint("deviceID", id), zap.Error(err))
			continue
		}
		if eq.SctiveState != nil && *eq.SctiveState == "1" {
			alreadyActivated = append(alreadyActivated, id)
		} else {
			toActivate = append(toActivate, id)
		}
	}

	if len(alreadyActivated) == len(req.IDs) {
		response.FailWithMessage("所选设备均已激活，无需重复激活", c)
		return
	}

	if len(toActivate) > 0 {
		err := global.GVA_DB.Model(&system.Equipment{}).
			Where("id IN ?", toActivate).
			Updates(map[string]interface{}{"sctive_state": "1"}).Error
		if err != nil {
			global.GVA_LOG.Error("激活失败", zap.Error(err))
			response.FailWithMessage("激活失败", c)
			return
		}
		// 这里统一返回详细提示
		if len(alreadyActivated) > 0 {
			response.OkWithMessage(fmt.Sprintf("部分设备已激活（ID: %v），其余设备已激活成功", alreadyActivated), c)
		} else {
			response.OkWithMessage("激活成功", c)
		}
		return
	}

	// 如果没有需要激活的设备
	response.OkWithMessage("所有设备已激活", c)
}
