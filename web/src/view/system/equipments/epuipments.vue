<template>
  <div class="dashboard-container">
    <!-- 设备情况 + 待办事项 顶部区域 -->
    <div class="top-section">
      <!-- 设备情况 -->
      <div class="device-status">
        <div class="status-card use-stat">
          <h3>使用统计</h3>
          <div class="status-item" :class="{'online':deviceStats.total}">全部设备：{{ deviceStats.total || 0}}台</div>
          <div class="status-item" :class="{'online':deviceStats.inUse}">在用设备：{{ deviceStats.total }}台</div>
          <div class="status-item" :class="{'online':deviceStats.idle}">闲置设备：{{ deviceStats.idle }}台</div>
          <div class="status-item" :class="{'online':deviceStats.stopped}"> 停用设备：{{ deviceStats.stopped }}台</div>
        </div>
        <div class="status-card activation-stat">
          <h3>激活情况（在用设备）</h3>
          <div class="status-item">已激活设备：{{ activationStats.activated }}台</div>
          <div class="status-item">未激活设备：{{ activationStats.unactivated }}台</div>
        </div>
        <div class="status-card running-stat">
          <h3>运行状态（已激活设备）</h3>
          <div class="status-item">正常设备：{{ runningStats.normal }}台</div>
          <div class="status-item">故障设备：{{ runningStats.fault }}台</div>
        </div>
      </div>
      <!-- 待办事项 -->
      <div class="todo-section">
        <h3>待办事项</h3>
        <div class="todo-list">
          <div class="todo-item">待激活：{{ todoItems.activate }}</div>
          <div class="todo-item">待安装：{{ todoItems.install }}</div>
          <div class="todo-item">型号补充：{{ todoItems.model }}</div>
          <div class="todo-item">安装地点补充：{{ todoItems.location }}</div>
        </div>
      </div>
    </div>

    <!-- 中间区域：设备分布地图 + 在用设备排行 -->
    <div class="middle-section">
      <!-- 设备安装分布统计（ECharts 地图） -->
      <div class="map-container">
        <div ref="mapChart" class="map-chart"></div>
        <div class="controls">
          <el-button type="primary" @click="refreshMap">
            <i class="el-icon-refresh-right"></i> 刷新地图
          </el-button>
          <el-button type="success" @click="addRandomDevice">
            <i class="el-icon-plus"></i> 添加设备
          </el-button>
          <el-input
            v-model="searchKeyword"
            placeholder="搜索设备名称或类型"
            @keyup.enter="filterDevices"
          >
            <template #prefix>
              <i class="el-icon-search"></i>
            </template>
          </el-input>
        </div>
        <div class="legend">
          <div class="legend-item">
            <span class="legend-dot" style="background-color: #1ae364;"></span>
            <span class="legend-text">在线设备</span>
          </div>
          <div class="legend-item">
            <span class="legend-dot" style="background-color: #64748b;"></span>
            <span class="legend-text">离线设备</span>
          </div>
          <div class="legend-item">
            <span class="legend-dot" style="background-color: #ef4444;"></span>
            <span class="legend-text">故障设备</span>
          </div>
        </div>
        <div v-if="errorMessage" class="error-message">
          <i class="el-icon-warning"></i> {{ errorMessage }}
        </div>
      </div>

      <!-- 在用设备排行 -->
      <div class="ranking-section">
        <h3>在用设备排行（单位：台）</h3>
        <el-table :data="rankingData" border style="width: 100%" :show-header="false">
          <el-table-column prop="name" label="设备类型" align="left" />
          <el-table-column label="数量" align="right">
            <template #default="scope">
              <div class="bar-container">
                <div class="bar" :style="{ width: scope.row.count + '%', backgroundColor: '#d3dce6' }"></div>
                <span class="bar-text">{{ scope.row.count }}</span>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>
</template>

<script setup>
  import { ref, onMounted, onUnmounted } from 'vue'
  import * as echarts from 'echarts'
  import axios from 'axios'
 import { queryDeviceCountByStatus } from '@/api/system/equipment';



  const deviceStats = ref({
    total: 0,
    inUse: 0,
    idle: 0,
    stopped: 0,
    status: "1",
  })

  const fetchDeviceStats = async () => {
    try {
      const res = await queryDeviceCountByStatus(deviceStats.value);

      deviceStats.value=res.data;
    }catch (error) {
      console.error('获取设备统计数出错',error)
    }
  };
  onMounted(() => {
    fetchDeviceStats()
  })
  // 激活情况数据
  const activationStats = ref({
    activated: 0,
    unactivated: 0
  })
  // 运行状态数据
  const runningStats = ref({
    normal: 0,
    fault: 0
  })
  // 待办事项数据
  const todoItems = ref({
    activate: 0,
    install: 0,
    model: 0,
    location: 0
  })
  // 在用设备排行数据
  const rankingData = ref([
    { name: '光刻机', count: 94 },
    { name: '刻蚀机', count: 79 },
    { name: '薄膜沉积设备', count: 60 },
    { name: '离子注入机', count: 45 },
    { name: '机械抛光', count: 40 },
    { name: '清洗机', count: 21 },
    { name: '氧化炉', count: 18 },
    { name: '晶圆切割机', count: 15 },
    { name: '晶元键合机', count: 11 },
    { name: '测试机', count: 8 }
  ])

  // 地图相关
  const mapChart = ref(null)
  let myChart = null
  const errorMessage = ref('')
  const searchKeyword = ref('')
  const deviceList = ref([])

  // 模拟设备数据（实际应从后端获取）
  const generateRandomDevices = (count = 10) => {
    const devices = []
    const cities = [
      { name: '北京', lng: 116.407526, lat: 39.90403 },
      { name: '上海', lng: 121.473701, lat: 31.230416 },
      { name: '广州', lng: 113.264435, lat: 23.12911 },
      { name: '深圳', lng: 114.057868, lat: 22.543099 },
      { name: '成都', lng: 104.065735, lat: 30.572269 },
      { name: '杭州', lng: 120.15358, lat: 30.287458 },
      { name: '南京', lng: 118.767413, lat: 32.041544 },
      { name: '武汉', lng: 114.298572, lat: 30.584355 },
      { name: '重庆', lng: 106.504959, lat: 29.533155 },
      { name: '西安', lng: 108.948021, lat: 34.263161 }
    ]

    for (let i = 0; i < count; i++) {
      const city = cities[Math.floor(Math.random() * cities.length)]
      const lngOffset = (Math.random() - 0.5) * 2
      const latOffset = (Math.random() - 0.5) * 2

      devices.push({
        id: `device-${Date.now()}-${i}`,
        name: `设备${i + 1}`,
        type: ['服务器', '传感器', '摄像头', '路由器'][Math.floor(Math.random() * 4)],
        status: ['在线', '离线', '故障'][Math.floor(Math.random() * 3)],
        lng: city.lng + lngOffset,
        lat: city.lat + latOffset,
        temperature: Math.floor(Math.random() * 40) + 20,
        load: Math.floor(Math.random() * 100)
      })
    }
    return devices
  }

  onMounted(async () => {
    // 初始化地图
    try {
      console.log('尝试加载地图数据...')
      const response = await axios.get('/china.json')

      if (typeof response.data!== 'object' || response.data === null) {
        throw new Error('获取的地图数据不是有效的JSON格式')
      }

      echarts.registerMap('china', response.data)
      console.log('中国地图注册成功')

      myChart = echarts.init(mapChart.value)

      // 显示加载动画
      myChart.showLoading({
        text: '地图加载中...',
        color: '#36CFFB',
        textColor: '#FFFFFF',
        maskColor: 'rgba(15, 23, 42, 0.8)',
        zlevel: 0
      })

      // 模拟网络延迟，确保DOM已完全渲染
      setTimeout(() => {
        deviceList.value = generateRandomDevices(15)
        renderMap()
        myChart.hideLoading()
      }, 500)

      window.addEventListener('resize', () => {
        myChart.resize()
      })
    } catch (error) {
      console.error('地图初始化失败:', error)
      errorMessage.value = `地图加载失败: ${error.message}`
    }
  })

  onUnmounted(() => {
    if (myChart) {
      myChart.dispose()
    }
    window.removeEventListener('resize', () => {
      myChart.resize()
    })
  })

  const renderMap = () => {
    try {
      // 构建用于 lines 系列的数据，从地图坐标向上延伸，这里简单将终点纬度固定设大一些模拟向上延伸
      const lineData = deviceList.value.map(device => ({
        coords: [
          [device.lng, device.lat],
          [device.lng, device.lat + 3] // 让线向上延伸，可根据需求调整长度
        ],
        status: device.status,
        deviceInfo: device
      }))

      // 构建 scatter 系列的数据（用于显示线的顶端“光点”等，也可根据需求调整样式）
      const scatterData = deviceList.value.map(device => ({
        value: [device.lng, device.lat + 0],
        status: device.status,
        deviceInfo: device
      }))

      const option = {
        backgroundColor: 'rgba(15, 23, 42, 0.95)',
        title: {
          text: '全国设备地理位置分布图',
          left: 'center',
          textStyle: {
            color: '#e2e8f0',
            fontSize: 20,
            fontWeight: 500
          }
        },
        tooltip: {
          trigger: 'item',
          backgroundColor: 'rgba(15, 23, 42, 0.85)',
          borderColor: 'rgba(54, 207, 251, 0.5)',
          borderWidth: 1,
          textStyle: {
            color: '#E2E8F0'
          },
          formatter: (params) => {
            const device = params.data.deviceInfo
            return `
            <div style="font-weight:bold; color:#36CFFB">${device.name}</div>
            <div style="color:#94A3B8">类型: ${device.type}</div>
            <div style="color:#94A3B8">状态: <span style="color:${getDeviceColor(device.status)}">${getStatusText(device.status)}</span></div>
            <div style="color:#94A3B8">位置: ${device.lng.toFixed(2)}, ${device.lat.toFixed(2)}</div>
            <div style="color:#94A3B8">温度: ${device.temperature}°C</div>
            <div style="color:#94A3B8">负载: ${device.load}%</div>
          `
          }
        },
        geo: {
          map: 'china',
          roam: true,
          zoom: 1.2,
          label: {
            show: true,
            color: '#94a3b8'
          },
          itemStyle: {
            areaColor: 'rgba(30, 41, 59, 0.8)',
            borderColor: 'rgba(54, 207, 251, 0.3)',
            borderWidth: 0.8
          },
          emphasis: {
            itemStyle: {
              areaColor: 'rgba(54, 207, 251, 0.1)'
            }
          }
        },
        series: [
          {
            name: '设备光柱',
            type: 'lines',
            coordinateSystem: 'geo',
            data: lineData,
            effect: {
              show: true,
              period: 4, // 波纹动画周期
              trailLength: 0.4, // 波纹长度
              color: '#36CFFB', // 波纹颜色
              symbolSize: 6 // 波纹大小
            },
            lineStyle: {
              color: (params) => getDeviceColor(params.data.status),
              width: 2,
              opacity: 0.8,
              curveness: 0 // 线为直线
            }
          },
          {
            name: '光柱顶端',
            type: 'scatter',
            coordinateSystem: 'geo',
            data: scatterData,
            symbolSize: 6,
            itemStyle: {
              color: (params) => getDeviceColor(params.data.status)
            },
            label: {
              show: false
            }
          }
        ]
      }

      myChart.setOption(option)
    } catch (error) {
      console.error('渲染地图失败:', error)
      errorMessage.value = `地图渲染失败: ${error.message}`
    }
  }

  const getDeviceColor = (status) => {
    switch (status) {
      case '在线': return '#22c55e'
      case '离线': return '#64748b'
      case '故障': return '#ef4444'
      default: return '#3b82f6'
    }
  }

  const getStatusText = (status) => {
    const statusMap = {
      online: '在线',
      offline: '离线',
      fault: '故障'
    }
    return statusMap[status] || status
  }

  const refreshMap = () => {
    if (myChart) {
      myChart.showLoading()
      setTimeout(() => {
        deviceList.value = generateRandomDevices(15)
        renderMap()
        myChart.hideLoading()
      }, 500)
    }
  }

  const addRandomDevice = () => {
    const newDevice = generateRandomDevices(1)[0]
    deviceList.value = [...deviceList.value, newDevice]
    renderMap()
  }

  const filterDevices = () => {
    if (!searchKeyword.value.trim()) {
      renderMap()
      return
    }

    const keyword = searchKeyword.value.toLowerCase()
    const filtered = deviceList.value.filter(device =>
      device.name.toLowerCase().includes(keyword) ||
      device.type.toLowerCase().includes(keyword)
    )

    if (myChart) {
      // 这里简单处理，实际可根据过滤后的数据重新构建 lines 和 scatter 系列数据
      const lineData = filtered.map(device => ({
        coords: [
          [device.lng, device.lat],
          [device.lng, device.lat + 5]
        ],
        status: device.status,
        deviceInfo: device
      }))
      const scatterData = filtered.map(device => ({
        value: [device.lng, device.lat + 5],
        status: device.status,
        deviceInfo: device
      }))
      myChart.setOption({
        series: [
          {
            name: '设备光柱',
            data: lineData
          },
          {
            name: '光柱顶端',
            data: scatterData
          }
        ]
      })
    }
  }
</script>

<style scoped>
  /* 整体容器 */
  .dashboard-container {
    padding: 20px;
    background-color: #fff;
    font-family: Avenir, Helvetica, Arial, sans-serif;
  }

  /* 顶部区域布局 */
  .top-section {
    display: flex;
    flex-direction: column;
    gap: 20px;
    margin-bottom: 20px;
  }

  /* 设备状态卡片 */
  .device-status {
    display: flex;
    gap: 20px;
  }

  .status-card {
    flex: 1;
    background: #f8f9fa;
    padding: 15px;
    border-radius: 8px;
    box-shadow: 0 0 4px rgba(0, 0, 0, 0.1);
    transition: all 0.3s ease;
  }

  .status-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
  }

  .status-item {
    margin: 8px 0;
    padding: 6px 0;
    border-bottom: 1px solid #e9ecef;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .status-item:last-child {
    border-bottom: none;
  }

  /* 待办事项区域 */
  .todo-section {
    background: #f8f9fa;
    padding: 15px;
    border-radius: 8px;
    box-shadow: 0 0 4px rgba(0, 0, 0, 0.1);
    transition: all 0.3s ease;
  }

  .todo-section:hover {
    transform: translateY(-5px);
    box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
  }

  .todo-list {
    display: flex;
    gap: 20px;
    margin-top: 10px;
    flex-wrap: wrap;
  }

  .todo-item {
    padding: 10px 15px;
    background: #e9ecef;
    border-radius: 6px;
    flex: 1;
    min-width: 150px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    transition: all 0.3s ease;
  }

  .todo-item:hover {
    background: #dee2e6;
  }

  /* 中间区域布局 */
  .middle-section {
    display: flex;
    gap: 20px;
    height: 600px;
  }

  /* 地图容器 */
  .map-container {
    flex: 2;
    position: relative;
    background: linear-gradient(135deg, #0F172A 0%, #1E293B 100%);
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
  }

  .map-chart {
    width: 100%;
    height: 100%;
  }

  /* 地图控制栏 */
  .controls {
    position: absolute;
    bottom: 24px;
    left: 30px;
    z-index: 10;
    display: flex;
    gap: 16px;
    align-items: center;
  }

  .controls .el-button {
    padding: 10px 16px;
    border-radius: 8px;
    font-weight: 500;
    transition: all 0.3s ease;
    display: flex;
    align-items: center;
    gap: 8px;
    background: rgba(54, 207, 251, 0.1);
    border: 1px solid rgba(54, 207, 251, 0.3);
    color: #FFFFFF;
  }

  .controls .el-button:hover {
    background: rgba(54, 207, 251, 0.2);
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(54, 207, 251, 0.2);
  }

  .controls .el-button--primary {
    background: rgba(54, 207, 251, 0.8);
    border-color: rgba(54, 207, 251, 0.5);
  }

  .controls .el-button--primary:hover {
    background: rgba(54, 207, 251, 1);
  }

  .controls .el-button--success {
    background: rgba(54, 211, 153, 0.8);
    border-color: rgba(54, 211, 153, 0.5);
  }

  .controls .el-button--success:hover {
    background: rgba(54, 211, 153, 1);
  }

  .controls .el-input {
    width: 240px;
  }

  .controls .el-input__wrapper {
    background: rgba(30, 41, 59, 0.8);
    border: 1px solid rgba(54, 207, 251, 0.3);
    border-radius: 8px;
    color: #E2E8F0;
    box-shadow: none;
    transition: all 0.3s ease;
  }

  .controls .el-input__prefix {
    color: #94A3B8;
  }

  .controls .el-input__wrapper:focus,
  .controls .el-input__wrapper:hover {
    background: rgba(30, 41, 59, 1);
    border-color: rgba(54, 207, 251, 0.5);
  }

  /* 地图图例 */
  .legend {
    position: absolute;
    bottom: 24px;
    right: 24px;
    z-index: 10;
    display: flex;
    gap: 16px;
    background: rgba(15, 23, 42, 0.7);
    padding: 10px 16px;
    border-radius: 8px;
    backdrop-filter: blur(4px);
    border: 1px solid rgba(54, 207, 251, 0.2);
  }

  .legend-item {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .legend-dot {
    width: 12px;
    height: 12px;
    border-radius: 50%;
    display: inline-block;
  }

  .legend-text {
    color: #E2E8F0;
    font-size: 12px;
  }

  /* 错误提示 */
  .error-message {
    position: absolute;
    top: 100px;
    left: 50%;
    transform: translateX(-50%);
    background: rgba(248, 114, 114, 0.9);
    color: #FFFFFF;
    padding: 12px 24px;
    border-radius: 8px;
    z-index: 20;
    display: flex;
    align-items: center;
    gap: 8px;
    box-shadow: 0 10px 20px rgba(248, 114, 114, 0.2);
    opacity: 0;
    animation: fadeIn 0.5s ease forwards;
  }

  /* 排行区域 */
  .ranking-section {
    flex: 1;
    background: #f8f9fa;
    padding: 15px;
    border-radius: 8px;
    box-shadow: 0 0 4px rgba(0, 0, 0, 0.1);
    transition: all 0.3s ease;
    display: flex;
    flex-direction: column;
  }

  .ranking-section:hover {
    transform: translateY(-5px);
    box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
  }

  .ranking-section .el-table {
    flex: 1;
  }

  .ranking-section .el-table__row {
    border-bottom: 1px solid #ebeef5;
  }

  .ranking-section .el-table__row:last-child {
    border-bottom: none;
  }

  .bar-container {
    display: flex;
    align-items: center;
    height: 24px;
    background-color: #e9ecef;
    border-radius: 4px;
    overflow: hidden;
    width: 100%;
  }

  .bar {
    height: 100%;
    transition: width 0.8s ease;
  }

  .bar-text {
    margin-left: 8px;
    font-weight: 500;
    color: #303133;
  }

  /* 动画效果 */
  @keyframes fadeIn {
    from {
      opacity: 0;
      transform: translate(-50%, -20px);
    }
    to {
      opacity: 1;
      transform: translate(-50%, 0);
    }
  }

  /* 响应式布局 */
  @media (max-width: 1200px) {
    .middle-section {
      flex-direction: column;
      height: auto;
    }

    .map-container, .ranking-section {
      min-height: 400px;
    }

    .device-status {
      flex-direction: column;
    }

    .controls {
      flex-wrap: wrap;
      bottom: 15px;
      left: 15px;
    }

    .controls .el-input {
      width: 100%;
      margin-top: 10px;
    }

    .legend {
      bottom: 15px;
      right: 15px;
    }
  }
</style>