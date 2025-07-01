<template>
  <div class="dashboard-container">
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-overlay">
      <el-loading-spinner></el-loading-spinner>
      <p>正在加载数据...</p>
    </div>

    <!-- 错误提示 -->
    <div v-if="error" class="error-banner">
      <i class="el-icon-warning"></i>
      {{ error }}
      <el-button type="text" @click="fetchDashboardData">重试</el-button>
    </div>

    <!-- 设备情况 + 待办事项 顶部区域 -->
    <div class="top-section">
      <!-- 设备情况 -->
      <div class="device-status">
        <div class="status-card use-stat">
          <h3>使用统计</h3>
          <div class="status-item" :class="{'online':deviceStats.total}">全部设备：{{ deviceStats.total || 0}}台</div>
          <div class="status-item" :class="{'online':deviceStats.inUse}">在用设备：{{ deviceStats.inUse }}台</div>
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
      <!-- 设备安装分布统计（高德地图） -->
      <div class="map-container">
        <div ref="mapContainer" class="map-chart"></div>
        <div class="controls">
          <el-button type="primary" @click="refreshMap" :loading="loading">
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
        <div v-if="rankingData.length === 0" class="empty-data">
          <i class="el-icon-data-analysis"></i>
          <p>暂无排行数据</p>
        </div>
        <el-table v-else :data="rankingData" border style="width: 100%" :show-header="false">
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
  import AMapLoader from "@amap/amap-jsapi-loader";
  import axios from 'axios'
  import { queryDeviceCountByStatus, getDashboardStats } from '@/api/system/equipment';

  // 数据加载状态
  const loading = ref(false)
  const error = ref('')

  // 设备统计数据
  const deviceStats = ref({
    total: 0,
    inUse: 0,
    idle: 0,
    stopped: 0,
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
  const rankingData = ref([])

  // 地图设备数据
  const deviceList = ref([])

  // 获取仪表盘数据
  const fetchDashboardData = async () => {
    loading.value = true
    error.value = ''
    
    try {
      const res = await getDashboardStats()
      if (res.code === 0 && res.data) {
        const data = res.data
        
        // 更新设备统计数据
        if (data.deviceStats) {
          deviceStats.value = data.deviceStats
        }
        
        // 更新激活情况数据
        if (data.activationStats) {
          activationStats.value = data.activationStats
        }
        
        // 更新运行状态数据
        if (data.runningStats) {
          runningStats.value = data.runningStats
        }
        
        // 更新待办事项数据
        if (data.todoItems) {
          todoItems.value = data.todoItems
        }
        
        // 更新排行数据
        if (data.rankingData) {
          rankingData.value = data.rankingData
        }
        
        // 更新地图数据
        if (data.mapData) {
          deviceList.value = data.mapData.map(device => ({
            id: device.id,
            name: device.name || '未知设备',
            type: device.type || '未知类型',
            status: device.status || '未知',
            lng: device.longitude,
            lat: device.latitude,
            temperature: Math.floor(Math.random() * 40) + 20,
            load: Math.floor(Math.random() * 100)
          }))
        }
      }
    } catch (err) {
      console.error('获取仪表盘数据失败', err)
      error.value = '获取数据失败，请稍后重试'
    } finally {
      loading.value = false
    }
  }

  // 根据位置信息获取经度（这里需要根据实际的位置格式来解析）
  const getLocationLng = (location) => {
    if (!location) return 116.397428 // 默认北京经度
    
    // 这里可以根据实际的位置格式来解析经纬度
    // 例如：如果location是"北京市朝阳区"，可以通过地理编码服务获取经纬度
    // 暂时使用模拟数据
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
    
    // 简单的城市匹配逻辑
    for (const city of cities) {
      if (location.includes(city.name)) {
        return city.lng + (Math.random() - 0.5) * 2 // 添加一些随机偏移
      }
    }
    
    // 如果没有匹配到城市，返回随机位置
    const randomCity = cities[Math.floor(Math.random() * cities.length)]
    return randomCity.lng + (Math.random() - 0.5) * 2
  }

  // 根据位置信息获取纬度
  const getLocationLat = (location) => {
    if (!location) return 39.90923 // 默认北京纬度
    
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
    
    for (const city of cities) {
      if (location.includes(city.name)) {
        return city.lat + (Math.random() - 0.5) * 2
      }
    }
    
    const randomCity = cities[Math.floor(Math.random() * cities.length)]
    return randomCity.lat + (Math.random() - 0.5) * 2
  }

  // 地图相关
  const mapContainer = ref(null)
  let map = null
  let markers = []
  const errorMessage = ref('')
  const searchKeyword = ref('')

  onMounted(async () => {
    // 首先获取仪表盘数据
    await fetchDashboardData()
    
    // 然后初始化地图
    try {
      const AMap = await AMapLoader.load({
        key: "bcef5eb6dccedfc3c36e02df5c167ab2",
        version: "2.0",
        plugins: [],
      });

      map = new AMap.Map(mapContainer.value, {
        viewMode: "3D",
        zoom: 11,
        center: [116.397428, 39.90923],
      });

      renderMap()
    } catch (e) {
      console.log(e);
      errorMessage.value = `地图加载失败: ${e.message}`
    }
  })

  onUnmounted(() => {
    map?.destroy();
  })

  const renderMap = () => {
    if (!map) return
    
    // 清除之前的标记
    if (markers.length > 0) {
      map.remove(markers);
      markers = [];
    }

    deviceList.value.forEach(device => {
      if (device.lng && device.lat) {
        const marker = new AMap.Marker({
          position: [device.lng, device.lat],
          icon: getDeviceIcon(device.status),
          offset: new AMap.Pixel(-10, -10),
        });

        const infoWindow = new AMap.InfoWindow({
          content: `
          <div style="font-weight:bold; color:#36CFFB">${device.name}</div>
          <div style="color:#94A3B8">类型: ${device.type}</div>
          <div style="color:#94A3B8">状态: <span style="color:${getDeviceColor(device.status)}">${device.status}</span></div>
          <div style="color:#94A3B8">位置: ${device.lng.toFixed(2)}, ${device.lat.toFixed(2)}</div>
          <div style="color:#94A3B8">温度: ${device.temperature}°C</div>
          <div style="color:#94A3B8">负载: ${device.load}%</div>
        `,
          offset: new AMap.Pixel(0, -30)
        });

        marker.on('click', () => {
          infoWindow.open(map, marker.getPosition());
        });

        marker.setMap(map);
        markers.push(marker);
      }
    });
  }

  const getDeviceColor = (status) => {
    switch (status) {
      case '在线': return '#22c55e'
      case '离线': return '#64748b'
      case '故障': return '#ef4444'
      default: return '#3b82f6'
    }
  }

  const getDeviceIcon = (status) => {
    const color = getDeviceColor(status);
    return new AMap.Icon({
      image: `data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20"><circle cx="10" cy="10" r="8" fill="${color}" /></svg>`,
      size: new AMap.Size(20, 20),
      imageSize: new AMap.Size(20, 20)
    });
  }

  const refreshMap = async () => {
    await fetchDashboardData()
    renderMap()
  }

  const addRandomDevice = () => {
    // 这里可以调用创建设备的API
    console.log('添加设备功能需要实现')
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

    // 临时过滤显示
    const originalList = [...deviceList.value]
    deviceList.value = filtered
    renderMap()
    
    // 恢复原始列表
    setTimeout(() => {
      deviceList.value = originalList
    }, 3000)
  }

  async function getLngLatByAddress(address) {
    return new Promise((resolve, reject) => {
      AMapLoader.load({
        key: 'bcef5eb6dccedfc3c36e02df5c167ab2',
        version: '2.0',
        plugins: ['AMap.Geocoder']
      }).then(AMap => {
        const geocoder = new AMap.Geocoder()
        geocoder.getLocation(address, (status, result) => {
          if (status === 'complete' && result.geocodes.length) {
            const { location } = result.geocodes[0]
            resolve({ lng: location.lng, lat: location.lat })
          } else {
            reject('地址解析失败')
          }
        })
      })
    })
  }
</script>

<style scoped>
  /* 整体容器 */
  .dashboard-container {
    padding: 20px;
    background-color: #fff;
    font-family: Avenir, Helvetica, Arial, sans-serif;
    position: relative;
  }

  /* 加载状态覆盖层 */
  .loading-overlay {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(255, 255, 255, 0.9);
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    backdrop-filter: blur(4px);
  }

  .loading-overlay p {
    margin-top: 16px;
    color: #606266;
    font-size: 14px;
  }

  /* 错误提示横幅 */
  .error-banner {
    background: linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%);
    color: white;
    padding: 12px 20px;
    border-radius: 8px;
    margin-bottom: 20px;
    display: flex;
    align-items: center;
    gap: 12px;
    box-shadow: 0 4px 12px rgba(255, 107, 107, 0.3);
    animation: slideInDown 0.3s ease;
  }

  .error-banner i {
    font-size: 18px;
  }

  .error-banner .el-button {
    color: white;
    margin-left: auto;
  }

  .error-banner .el-button:hover {
    color: #f0f0f0;
  }

  /* 空数据状态 */
  .empty-data {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 40px 20px;
    color: #909399;
    text-align: center;
  }

  .empty-data i {
    font-size: 48px;
    margin-bottom: 16px;
    opacity: 0.5;
  }

  .empty-data p {
    margin: 0;
    font-size: 14px;
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

  @keyframes slideInDown {
    from {
      opacity: 0;
      transform: translateY(-20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
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