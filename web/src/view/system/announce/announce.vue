<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="公告类型" prop="type">
          <el-select v-model="searchInfo.type" clearable filterable placeholder="请选择公告类型">
            <el-option label="版本更新" value="version_update" />
            <el-option label="系统维护" value="system_maintenance" />
            <el-option label="重要通知" value="important_notice" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog()">发布公告</el-button>
        <el-button type="success" icon="video-play" @click="testWebSocket">测试WebSocket连接</el-button>
      </div>

      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
      >
        <el-table-column align="left" label="公告类型" prop="type" width="120">
          <template #default="scope">
            <el-tag :type="getTypeTagType(scope.row.type)">
              {{ getTypeLabel(scope.row.type) }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column align="left" label="标题" prop="title" width="200" />

        <el-table-column align="left" label="内容" prop="content" width="300" show-overflow-tooltip />

        <el-table-column align="left" label="版本号" prop="version" width="120" />

        <el-table-column align="left" label="发布时间" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="操作" fixed="right" width="200">
          <template #default="scope">
            <el-button type="primary" link icon="view" @click="viewAnnounce(scope.row)">查看</el-button>
            <el-button type="success" link icon="video-play" @click="resendAnnounce(scope.row)">重新发送</el-button>
            <el-button type="danger" link icon="delete" @click="deleteAnnounce(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!-- 发布公告弹窗 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">发布版本更新公告</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="publishAnnounce">发布公告</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="公告类型:" prop="type">
          <el-select v-model="formData.type" placeholder="请选择公告类型" style="width:100%">
            <el-option label="版本更新" value="version_update" />
            <el-option label="系统维护" value="system_maintenance" />
            <el-option label="重要通知" value="important_notice" />
          </el-select>
        </el-form-item>

        <el-form-item label="公告标题:" prop="title">
          <el-input v-model="formData.title" placeholder="请输入公告标题" />
        </el-form-item>

        <el-form-item label="公告内容:" prop="content">
          <el-input 
            v-model="formData.content" 
            type="textarea" 
            :rows="4"
            placeholder="请输入公告内容" 
          />
        </el-form-item>

        <el-form-item label="版本号:" prop="version">
          <el-input v-model="formData.version" placeholder="请输入版本号，如：v1.2.3" />
        </el-form-item>

        <el-form-item label="详情链接:" prop="url">
          <el-input v-model="formData.url" placeholder="请输入详情链接（可选）" />
        </el-form-item>

        <el-form-item label="预览效果:">
          <div class="preview-box" style="border: 1px solid #dcdfe6; border-radius: 4px; padding: 16px; background-color: #f5f7fa;">
            <h4 style="margin: 0 0 8px 0; color: #303133;">{{ formData.title || '公告标题' }}</h4>
            <p style="margin: 0 0 8px 0; color: #606266; line-height: 1.5;">{{ formData.content || '公告内容' }}</p>
            <div style="font-size: 12px; color: #909399;">
              <span>版本: {{ formData.version || 'v1.0.0' }}</span>
              <span v-if="formData.url" style="margin-left: 16px;">
                <a :href="formData.url" target="_blank" style="color: #409eff;">查看详情</a>
              </span>
            </div>
          </div>
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- 查看公告详情弹窗 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="公告详情">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="公告类型">
          <el-tag :type="getTypeTagType(detailFrom.type)">
            {{ getTypeLabel(detailFrom.type) }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="公告标题">
          {{ detailFrom.title }}
        </el-descriptions-item>
        <el-descriptions-item label="公告内容">
          {{ detailFrom.content }}
        </el-descriptions-item>
        <el-descriptions-item label="版本号">
          {{ detailFrom.version }}
        </el-descriptions-item>
        <el-descriptions-item label="详情链接" v-if="detailFrom.url">
          <a :href="detailFrom.url" target="_blank">{{ detailFrom.url }}</a>
        </el-descriptions-item>
        <el-descriptions-item label="发布时间">
          {{ formatDate(detailFrom.createdAt) }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia"
import { formatDate } from '@/utils/format'
import { pushVersionUpdate } from '@/api/system/announce'

defineOptions({
  name: 'Announce'
})

const appStore = useAppStore()

// 提交按钮loading
const btnLoading = ref(false)

// 表格数据
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 弹窗控制
const dialogFormVisible = ref(false)
const detailShow = ref(false)
const detailFrom = ref({})

// 表单数据
const formData = ref({
  type: 'version_update',
  title: '',
  content: '',
  version: '',
  url: ''
})

// 表单验证规则
const rule = reactive({
  type: [{ required: true, message: '请选择公告类型', trigger: 'change' }],
  title: [{ required: true, message: '请输入公告标题', trigger: 'blur' }],
  content: [{ required: true, message: '请输入公告内容', trigger: 'blur' }],
  version: [{ required: true, message: '请输入版本号', trigger: 'blur' }]
})

const elFormRef = ref()
const elSearchFormRef = ref()

// 获取公告类型标签
const getTypeLabel = (type) => {
  const typeMap = {
    'version_update': '版本更新',
    'system_maintenance': '系统维护',
    'important_notice': '重要通知'
  }
  return typeMap[type] || type
}

// 获取公告类型标签颜色
const getTypeTagType = (type) => {
  const typeMap = {
    'version_update': 'success',
    'system_maintenance': 'warning',
    'important_notice': 'danger'
  }
  return typeMap[type] || 'info'
}

// 搜索
const onSubmit = () => {
  page.value = 1
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 获取表格数据（模拟数据）
const getTableData = async() => {
  // 这里应该调用真实的API，现在用模拟数据
  const mockData = [
    {
      id: 1,
      type: 'version_update',
      title: '系统版本更新 v1.2.3',
      content: '新版本包含多项功能优化和bug修复，建议及时更新。',
      version: 'v1.2.3',
      url: 'https://example.com/release-notes',
      createdAt: new Date()
    },
    {
      id: 2,
      type: 'system_maintenance',
      title: '系统维护通知',
      content: '系统将于今晚22:00-24:00进行维护，期间可能影响正常使用。',
      version: '',
      url: '',
      createdAt: new Date(Date.now() - 86400000)
    }
  ]
  
  tableData.value = mockData
  total.value = mockData.length
}

// 打开发布公告弹窗
const openDialog = () => {
  dialogFormVisible.value = true
  formData.value = {
    type: 'version_update',
    title: '',
    content: '',
    version: '',
    url: ''
  }
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    type: 'version_update',
    title: '',
    content: '',
    version: '',
    url: ''
  }
}

// 发布公告
const publishAnnounce = async() => {
  btnLoading.value = true
  elFormRef.value?.validate(async (valid) => {
    if (!valid) {
      btnLoading.value = false
      return
    }
    
    try {
      const res = await pushVersionUpdate(formData.value)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '公告发布成功！'
        })
        closeDialog()
        getTableData()
      } else {
        ElMessage({
          type: 'error',
          message: res.msg || '发布失败'
        })
      }
    } catch (error) {
      ElMessage({
        type: 'error',
        message: '发布失败：' + error.message
      })
    } finally {
      btnLoading.value = false
    }
  })
}

// 查看公告详情
const viewAnnounce = (row) => {
  detailFrom.value = row
  detailShow.value = true
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}

// 重新发送公告
const resendAnnounce = async(row) => {
  try {
    const res = await pushVersionUpdate({
      type: row.type,
      title: row.title,
      content: row.content,
      version: row.version,
      url: row.url
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '公告重新发送成功！'
      })
    } else {
      ElMessage({
        type: 'error',
        message: res.msg || '发送失败'
      })
    }
  } catch (error) {
    ElMessage({
      type: 'error',
      message: '发送失败：' + error.message
    })
  }
}

// 删除公告
const deleteAnnounce = (row) => {
  ElMessageBox.confirm('确定要删除这条公告吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    // 这里应该调用删除API
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    getTableData()
  })
}

// 测试WebSocket连接
const testWebSocket = () => {
  ElMessage({
    type: 'info',
    message: 'WebSocket连接测试功能已实现，请查看控制台输出'
  })
  console.log('WebSocket连接状态：', window.ws ? '已连接' : '未连接')
}

// 初始化
onMounted(() => {
  getTableData()
})
</script>

<style scoped>
.preview-box {
  min-height: 100px;
}
</style> 