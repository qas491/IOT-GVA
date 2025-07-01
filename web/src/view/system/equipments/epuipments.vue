<template>
  <div class="dashboard-container">
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
  import AMapLoader from "@amap/amap-jsapi-loader";
  import axios from 'axios'
  import { queryDeviceCountByStatus } from '@/api/system/equipment';

  const deviceStats = ref({
    total: 10,
    inUse: 4,
    idle: 5,
    stopped: 1,
    status: "1",
  })

  const fetchDeviceStats = async () => {
    try {
      const res = await queryDeviceCountByStatus(deviceStats.value);
      deviceStats.value = res.data;
    } catch (error) {
      console.error('获取设备统计数出错', error)
    }
  };

  onMounted(() => {
    fetchDeviceStats()
  })

  // 激活情况数据
  const activationStats = ref({
    activated: 5,
    unactivated: 7
  })
  // 运行状态数据
  const runningStats = ref({
    normal: 5,
    fault: 3
  })
  // 待办事项数据
  const todoItems = ref({
    activate: 3,
    install: 1,
    model: 2,
    location: 5
  })
  // 在用设备排行数据
  const rankingData = ref([
    { name: '光刻机', count: 5 },
    { name: '刻蚀机', count: 0 },
    { name: '薄膜沉积设备', count: 3 },
    { name: '离子注入机', count: 1 },
    { name: '机械抛光', count: 5 },
    { name: '清洗机', count: 6 },
    { name: '氧化炉', count: 1 },
    { name: '晶圆切割机', count: 3 },
    { name: '晶元键合机', count: 7 },
    { name: '测试机', count: 8 }
  ])

  // 地图相关
  const mapContainer = ref(null)
  let map = null
  let markers = []
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
    try {
      const AMap = await AMapLoader.load({
        key: "bcef5eb6dccedfc3c36e02df5c167ab2", // 申请好的Web端开发者Key，首次调用 load 时必填
        version: "2.0", // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
        plugins: [], // 需要使用的的插件列表，如比例尺'AMap.Scale'等
      });

      map = new AMap.Map(mapContainer.value, {
        viewMode: "3D", // 是否为3D地图模式
        zoom: 11, // 初始化地图级别
        center: [116.397428, 39.90923], // 初始化地图中心点位置
      });

      deviceList.value = generateRandomDevices(15)
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
    // 清除之前的标记
    if (markers.length > 0) {
      map.remove(markers);
      markers = [];
    }

    deviceList.value.forEach(device => {
      const marker = new AMap.Marker({
        position: [device.lng, device.lat],
        icon: getDeviceIcon(device.status),
        offset: new AMap.Pixel(-10, -10),
      });

      const infoWindow = new AMap.InfoWindow({
        content: `
        <div style="font-weight:bold; color:#36CFFB">${device.name}</div>
        <div style="color:#94A3B8">类型: ${device.type}</div>
        <div style="color:#94A3B8">状态: <span style="color:${getDeviceColor(device.status)}">${getStatusText(device.status)}</span></div>
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

  const getStatusText = (status) => {
    const statusMap = {
      online: '在线',
      offline: '离线',
      fault: '故障'
    }
    return statusMap[status] || status
  }

  const refreshMap = () => {
    deviceList.value = generateRandomDevices(15)
    renderMap()
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

    deviceList.value = filtered
    renderMap()
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