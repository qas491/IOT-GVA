<template>
  <div class="device-dashboard">
    <!-- 顶部统计区 -->
    <div class="top-container">
      <!-- 设备数量环形图 -->
      <div class="card">
        <div class="title">当前设备数量</div>
        <div class="chart-wrap">
          <div id="ringChart" style="width: 160px; height: 160px;"></div>
        </div>
        <div class="stats">
          <p class="total">设备总量：{{ deviceStats.total }}</p>
          <p><span class="dot active"></span>已激活：{{ deviceStats.active }}</p>
          <p><span class="dot inactive"></span>未激活：{{ deviceStats.inactive }}</p>
          <p><span class="dot online"></span>在线设备：{{ deviceStats.online }}</p>
          <p><span class="dot offline"></span>离线设备：{{ deviceStats.offline }}</p>
        </div>
      </div>

      <!-- 当日消息折线图 -->
      <div class="card">
        <div class="title">当日设备消息数量</div>
        <div class="chart-wrap">
          <div id="lineChart" style="width: 300px; height: 160px;"></div>
        </div>
        <div class="sub-title">当月设备消息总量：{{ msgStats.monthTotal }}</div>
      </div>

      <!-- CPU 仪表盘 -->
      <div class="card">
        <div class="title">CPU 使用率</div>
        <div class="chart-wrap">
          <div id="cpuGauge" style="width: 160px; height: 160px;"></div>
        </div>
        <div class="value">{{ cpuUsage }}%</div>
      </div>

      <!-- 内存仪表盘 -->
      <div class="card">
        <div class="title">内存使用率</div>
        <div class="chart-wrap">
          <div id="memoryGauge" style="width: 160px; height: 160px;"></div>
        </div>
        <div class="value">{{ memoryUsage }}%</div>
      </div>
    </div>

    <!-- 设备状态柱状图 -->
    <div class="middle-container">
      <div class="header">
        <div class="title">设备状态</div>
        <div class="update-time">数据更新时间：{{ updateTime }}</div>
      </div>
      <div id="barChart" style="width: 100%; height: 300px;"></div>
    </div>

    <!-- 设备消息折线图（带时间筛选） -->
    <div class="bottom-container">
      <div class="header">
        <div class="title">设备消息</div>
        <div class="filter">
          <button
            v-for="(type, index) in filterTypes"
            :key="index"
            @click="handleFilter(type)"
            :class="{ active: currentFilter === type }"
          >{{ type }}</button>
          <input
            v-if="currentFilter === '自定义'"
            type="date"
            class="custom-date"
            @change="handleCustomDate"
          >
        </div>
      </div>
      <div id="msgLineChart" style="width: 100%; height: 250px;"></div>
    </div>
  </div>
</template>

<script setup>
  import { onMounted, ref } from 'vue'
  import * as echarts from 'echarts'

  // 模拟数据（实际需对接接口）
  const deviceStats = ref({
    total: 15360,
    active: 15325,
    inactive: 35,
    online: 15210,
    offline: 115
  })
  const msgStats = ref({
    monthTotal: 286360,
    dayData: [20, 30, 10, 50, 40, 60, 30, 70, 100] // 当日消息数据
  })
  const cpuUsage = ref(10.00)
  const memoryUsage = ref(15.00)
  const updateTime = ref('2021/06/05 23:30:05')
  const barData = {
    categories: ['智能门禁','车辆道闸','智能梯控','电气火灾','智能烟感','视频网关','智能水表','智能电表','智能家居'],
    online: [95, 90, 98, 85, 96, 99, 97, 96, 95], // 在线率（百分比）
    offline: [5, 10, 2, 15, 4, 1, 3, 4, 5] // 离线率（百分比）
  }
  const filterTypes = ref(['日', '月', '年', '自定义'])
  const currentFilter = ref('日')
  const msgLineData = ref([20, 25, 18, 30, 22, 28, 19, 35, 24, 20]) // 示例数据

  onMounted(() => {
    renderRingChart()
    renderLineChart()
    renderCpuGauge()
    renderMemoryGauge()
    renderBarChart()
    renderMsgLineChart()
  })

  // 1. 渲染环形图
  function renderRingChart() {
    const chartDom = document.getElementById('ringChart')
    const myChart = echarts.init(chartDom)
    const option = {
      series: [
        {
          type: 'pie',
          radius: ['40%', '70%'],
          center: ['50%', '50%'],
          data: [
            { value: deviceStats.value.active, name: '已激活' },
            { value: deviceStats.value.inactive, name: '未激活' },
            { value: deviceStats.value.online, name: '在线设备' },
            { value: deviceStats.value.offline, name: '离线设备' }
          ],
          label: { show: false },
          labelLine: { show: false },
          itemStyle: {
            borderRadius: 10,
            borderColor: '#fff',
            borderWidth: 2
          }
        }
      ]
    }
    myChart.setOption(option)
  }

  // 2. 渲染当日消息折线图
  function renderLineChart() {
    const chartDom = document.getElementById('lineChart')
    const myChart = echarts.init(chartDom)
    const option = {
      xAxis: {
        type: 'category',
        show: false,
        data: Array.from({ length: 24 }, (_, i) => `${i}:00`)
      },
      yAxis: { type: 'value', show: false },
      series: [
        {
          type: 'line',
          data: msgStats.value.dayData,
          smooth: true,
          areaStyle: {
            color: {
              type: 'linear',
              x: 0,
              y: 0,
              x2: 0,
              y2: 1,
              colorStops: [{
                offset: 0, color: 'rgba(135,206,250,0.3)'
              }, {
                offset: 1, color: 'rgba(135,206,250,0)'
              }]
            }
          },
          lineStyle: { color: '#87cefa', width: 2 },
          symbol: 'circle',
          symbolSize: 6
        }
      ],
      grid: { left: 0, right: 0, top: 0, bottom: 0 }
    }
    myChart.setOption(option)
  }

  // 3. 渲染 CPU 仪表盘
  function renderCpuGauge() {
    const chartDom = document.getElementById('cpuGauge')
    const myChart = echarts.init(chartDom)
    const option = {
      series: [
        {
          type: 'gauge',
          radius: '80%',
          center: ['50%', '50%'],
          startAngle: 180,
          endAngle: 0,
          pointer: { show: false },
          axisLine: {
            lineStyle: {
              width: 12,
              color: [[1, '#d1e6ff']]
            }
          },
          axisTick: { show: false },
          splitLine: { show: false },
          axisLabel: { show: false },
          detail: { show: false },
          data: [
            {
              value: cpuUsage.value,
              itemStyle: { color: '#87cefa' },
              name: '',
              label: {
                show: true,
                offsetCenter: [0, '40%'],
                color: '#666',
                fontSize: 14
              }
            }
          ]
        }
      ]
    }
    myChart.setOption(option)
  }

  // 4. 渲染内存仪表盘
  function renderMemoryGauge() {
    const chartDom = document.getElementById('memoryGauge')
    const myChart = echarts.init(chartDom)
    const option = {
      series: [
        {
          type: 'gauge',
          radius: '80%',
          center: ['50%', '50%'],
          startAngle: 180,
          endAngle: 0,
          pointer: { show: false },
          axisLine: {
            lineStyle: {
              width: 12,
              color: [[1, '#d1e6ff']]
            }
          },
          axisTick: { show: false },
          splitLine: { show: false },
          axisLabel: { show: false },
          detail: { show: false },
          data: [
            {
              value: memoryUsage.value,
              itemStyle: { color: '#90ee90' },
              name: '',
              label: {
                show: true,
                offsetCenter: [0, '40%'],
                color: '#666',
                fontSize: 14
              }
            }
          ]
        }
      ]
    }
    myChart.setOption(option)
  }

  // 5. 渲染设备状态柱状图
  function renderBarChart() {
    const chartDom = document.getElementById('barChart')
    const myChart = echarts.init(chartDom)
    const option = {
      xAxis: {
        type: 'category',
        data: barData.categories,
        axisLabel: { color: '#666', fontSize: 12 },
        axisLine: { lineStyle: { color: '#eee' } }
      },
      yAxis: {
        type: 'value',
        axisLabel: { color: '#666', formatter: '{value}%' },
        axisLine: { lineStyle: { color: '#eee' } },
        splitLine: { lineStyle: { color: '#f5f5f5' } }
      },
      series: [
        {
          name: '在线',
          type: 'bar',
          data: barData.online,
          itemStyle: { color: '#28a745' },
          barWidth: '35%',
          label: {
            show: true,
            position: 'top',
            color: '#28a745',
            formatter: '{c}%'
          }
        },
        {
          name: '离线',
          type: 'bar',
          data: barData.offline,
          itemStyle: { color: '#dee2e6' },
          barWidth: '35%',
          label: {
            show: true,
            position: 'top',
            color: '#666',
            formatter: '{c}%'
          }
        }
      ],
      grid: { left: '10px', right: '10px', bottom: '20px', containLabel: true }
    }
    myChart.setOption(option)
  }

  // 6. 渲染设备消息折线图
  function renderMsgLineChart() {
    const chartDom = document.getElementById('msgLineChart')
    const myChart = echarts.init(chartDom)
    const option = {
      xAxis: {
        type: 'category',
        data: Array.from({ length: 10 }, (_, i) => `${i + 1}日`),
        axisLabel: { color: '#666', fontSize: 12 },
        axisLine: { lineStyle: { color: '#eee' } }
      },
      yAxis: {
        type: 'value',
        axisLabel: { color: '#666' },
        axisLine: { lineStyle: { color: '#eee' } },
        splitLine: { lineStyle: { color: '#f5f5f5' } }
      },
      series: [
        {
          type: 'line',
          data: msgLineData.value,
          smooth: true,
          areaStyle: {
            color: {
              type: 'linear',
              x: 0,
              y: 0,
              x2: 0,
              y2: 1,
              colorStops: [{
                offset: 0, color: 'rgba(255, 215, 0, 0.2)'
              }, {
                offset: 1, color: 'rgba(255, 215, 0, 0)'
              }]
            }
          },
          lineStyle: { color: '#ffd700', width: 2 },
          symbol: 'circle',
          symbolSize: 6,
          label: {
            show: true,
            position: 'top',
            color: '#ffd700',
            fontSize: 12
          }
        }
      ],
      grid: { left: '10px', right: '10px', bottom: '20px', containLabel: true }
    }
    myChart.setOption(option)
  }

  // 筛选逻辑（示例，实际需对接接口）
  function handleFilter(type) {
    currentFilter.value = type
    if (type === '自定义') return
    // 模拟切换数据
    msgLineData.value = Array.from({ length: 10 }, () => Math.floor(Math.random() * 30) + 20)
    renderMsgLineChart()
  }

  function handleCustomDate(e) {
    // 实际需根据日期请求接口，这里仅演示
    console.log('自定义日期：', e.target.value)
    msgLineData.value = Array.from({ length: 10 }, () => Math.floor(Math.random() * 30) + 20)
    renderMsgLineChart()
  }
</script>

<style scoped>
  .device-dashboard {
    padding: 20px;
    background: #fff;
  }
  .top-container {
    display: flex;
    flex-wrap: wrap;
    gap: 20px;
    margin-bottom: 20px;
  }
  .card {
    flex: 1;
    min-width: 220px;
    background: #f8f9fa;
    border-radius: 8px;
    padding: 20px;
    box-shadow: 0 2px 8px rgba(0,0,0,0.05);
  }
  .title {
    font-size: 16px;
    font-weight: 500;
    color: #333;
    margin-bottom: 15px;
  }
  .chart-wrap {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-bottom: 15px;
  }
  .stats {
    font-size: 14px;
    color: #666;
    line-height: 1.6;
  }
  .dot {
    display: inline-block;
    width: 10px;
    height: 10px;
    border-radius: 50%;
    margin-right: 5px;
  }
  .active { background: #409eff; }
  .inactive { background: #6c757d; }
  .online { background: #28a745; }
  .offline { background: #dc3545; }
  .total { font-weight: 500; color: #333; }
  .sub-title {
    font-size: 14px;
    color: #666;
    margin-top: 10px;
  }
  .value {
    text-align: center;
    font-size: 16px;
    color: #333;
    margin-top: 10px;
  }
</style>