<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAtRange">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>

          <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="w-[380px]"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
        </el-form-item>

        <el-form-item label="设备编号" prop="equipmentNumber">
          <el-input v-model="searchInfo.equipmentNumber" placeholder="搜索条件" />
        </el-form-item>

        <el-form-item label="设备名称" prop="equipmentName">
          <el-input v-model="searchInfo.equipmentName" placeholder="搜索条件" />
        </el-form-item>

        <el-form-item label="运营状态" prop="operationalStatus">
          <el-select v-model="searchInfo.operationalStatus" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.operationalStatus=undefined}">
            <el-option v-for="(item,key) in OperationalstatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>

        <el-form-item label="运行状态" prop="nunningState">
          <el-select v-model="searchInfo.nunningState" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.nunningState=undefined}">
            <el-option v-for="(item,key) in runningstateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>

        <el-form-item label="激活状态" prop="sctiveState">
          <el-select v-model="searchInfo.sctiveState" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.sctiveState=undefined}">
            <el-option v-for="(item,key) in activestateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>

        <el-form-item label="软件版本" prop="softwareVersion">
          <el-input v-model="searchInfo.softwareVersion" placeholder="搜索条件" />
        </el-form-item>

        <el-form-item label="售后人员" prop="afterSalesPersonnel">
          <el-input v-model="searchInfo.afterSalesPersonnel" placeholder="搜索条件" />
        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
        <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="system_Equipment" />
        <ExportExcel v-auth="btnAuth.exportExcel" template-id="system_Equipment" filterDeleted/>
        <ImportExcel v-auth="btnAuth.importExcel" template-id="system_Equipment" @on-success="getTableData" />
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="设备编号" prop="equipmentNumber" width="120" />

        <el-table-column align="left" label="设备名称" prop="equipmentName" width="120" />

        <el-table-column align="left" label="所在位置" prop="location" width="120" />

        <el-table-column align="left" label="运营状态" prop="operationalStatus" width="120">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.operationalStatus, 'operational')">
              {{ filterDict(scope.row.operationalStatus,OperationalstatusOptions) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="运行状态" prop="nunningState" width="120">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.nunningState, 'running')">
              {{ filterDict(scope.row.nunningState,runningstateOptions) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="激活状态" prop="sctiveState" width="120">
          <template #default="scope">
            <span class="status-tag" :class="getStatusBgClass(scope.row.sctiveState)">
              <span class="status-dot" :class="getStatusClass(scope.row.sctiveState)"></span>
              {{ filterDict(scope.row.sctiveState, activestateOptions) }}
            </span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="软件版本" prop="softwareVersion" width="120" />

        <el-table-column align="left" label="售后人员" prop="afterSalesPersonnel" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateEquipmentFunc(scope.row)">编辑</el-button>
            <el-button  v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="设备编号:" prop="equipmentNumber">
          <el-input v-model="formData.equipmentNumber" :clearable="false" placeholder="请输入设备编号" />
        </el-form-item>
        <el-form-item label="设备名称:" prop="equipmentName">
          <el-input v-model="formData.equipmentName" :clearable="false" placeholder="请输入设备名称" />
        </el-form-item>
        <el-form-item label="所在位置:" prop="location">
          <el-input v-model="formData.location" :clearable="false" placeholder="请输入所在位置" />
        </el-form-item>
        <el-form-item label="运营状态:" prop="operationalStatus">
          <el-select v-model="formData.operationalStatus" placeholder="请选择运营状态" style="width:100%" filterable :clearable="false">
            <el-option v-for="(item,key) in OperationalstatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="运行状态:" prop="nunningState">
          <el-select v-model="formData.nunningState" placeholder="请选择运行状态" style="width:100%" filterable :clearable="true">
            <el-option v-for="(item,key) in runningstateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="激活状态:" prop="sctiveState">
          <el-select v-model="formData.sctiveState" placeholder="请选择激活状态" style="width:100%" filterable :clearable="true">
            <el-option v-for="(item,key) in activestateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="软件版本:" prop="softwareVersion">
          <el-input v-model="formData.softwareVersion" :clearable="true" placeholder="请输入软件版本" />
        </el-form-item>
        <el-form-item label="售后人员:" prop="afterSalesPersonnel">
          <el-input v-model="formData.afterSalesPersonnel" :clearable="true" placeholder="请输入售后人员" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="设备编号">
          {{ detailFrom.equipmentNumber }}
        </el-descriptions-item>
        <el-descriptions-item label="设备名称">
          {{ detailFrom.equipmentName }}
        </el-descriptions-item>
        <el-descriptions-item label="所在位置">
          {{ detailFrom.location }}
        </el-descriptions-item>
        <el-descriptions-item label="运营状态">
          <el-tag :type="getStatusType(detailFrom.operationalStatus, 'operational')">
            {{ detailFrom.operationalStatus }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="运行状态">
          <el-tag :type="getStatusType(detailFrom.nunningState, 'running')">
            {{ detailFrom.nunningState }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="激活状态">
          <span class="status-tag" :class="getStatusBgClass(detailFrom.sctiveState)">
            <span class="status-dot" :class="getStatusClass(detailFrom.sctiveState)"></span>
            {{ detailFrom.sctiveState }}
          </span>
        </el-descriptions-item>
        <el-descriptions-item label="软件版本">
          {{ detailFrom.softwareVersion }}
        </el-descriptions-item>
        <el-descriptions-item label="售后人员">
          {{ detailFrom.afterSalesPersonnel }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

  </div>
</template>

<script setup>
  import {
    createEquipment,
    deleteEquipment,
    deleteEquipmentByIds,
    updateEquipment,
    findEquipment,
    getEquipmentList
  } from '@/api/system/equipment'

  // 全量引入格式化工具 请按需保留
  import { getDictFunc, formatDate, filterDict } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive } from 'vue'
  // 引入按钮权限标识
  import { useBtnAuth } from '@/utils/btnAuth'
  import { useAppStore } from "@/pinia"

  // 导出组件
  import ExportExcel from '@/components/exportExcel/exportExcel.vue'
  // 导入组件
  import ImportExcel from '@/components/exportExcel/importExcel.vue'
  // 导出模板组件
  import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'

  defineOptions({
    name: 'Equipment'
  })

  // 按钮权限实例化
  const btnAuth = useBtnAuth()

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  // 自动化生成的字典（可能为空）以及字段
  const OperationalstatusOptions = ref([])
  const runningstateOptions = ref([])
  const activestateOptions = ref([])
  const formData = ref({
    equipmentNumber: '',
    equipmentName: '',
    location: '',
    operationalStatus: '',
    nunningState: '',
    sctiveState: '',
    softwareVersion: '',
    afterSalesPersonnel: '',
  })

  // 状态类映射配置（使用数值作为键）
  const statusClassMap = {
    '1': 'dot-success',    // 已激活 - 绿色
    '2': 'dot-warning',    // 未激活 - 橙色
    '3': 'dot-danger',     // 已过期 - 红色
    '4': 'dot-primary'     // 待激活 - 紫色
  }

  // 状态背景色映射（可选）
  const statusBgClassMap = {
    '1': 'bg-success-light',    // 已激活 - 浅绿色背景
    '2': 'bg-warning-light',    // 未激活 - 浅橙色背景
    '3': 'bg-danger-light',     // 已过期 - 浅红色背景
    '4': 'bg-primary-light'     // 待激活 - 浅蓝色背景
  }

  // 根据状态值获取对应的圆点CSS类
  const getStatusClass = (statusValue) => {
    return statusClassMap[statusValue] || 'dot-default'
  }

  // 根据状态值获取对应的背景CSS类（可选）
  const getStatusBgClass = (statusValue) => {
    return statusBgClassMap[statusValue] || ''
  }

  // 状态类型映射配置（保留原有的el-tag映射）
  const statusTypeMap = {
    operational: { // 运营状态
      '2': 'danger',        // 故障 - 红色
      '1': 'success',      // 已激活 - 绿色
      '3': 'grey',
    },
    running: { // 运行状态
      '2': 'danger',        // 故障 - 红色
      '1': 'success',     // 运行中 - 绿色
    },
    active: { // 激活状态
      '2': 'warning',    // 未激活 - 橙色
      '1': 'success',      // 已激活 - 绿色
      'pending': 'primary'      // 待激活 - 紫色
    }
  }

  // 根据状态值获取对应的Tag类型（保留原有方法）
  const getStatusType = (statusValue, statusType) => {
    if (!statusValue || !statusTypeMap[statusType]) return ''
    return statusTypeMap[statusType][statusValue] || ''
  }

  // 验证规则
  const rule = reactive({
    equipmentNumber : [{
      required: true,
      message: '',
      trigger: ['input','blur'],
    },
      {
        whitespace: true,
        message: '不能只输入空格',
        trigger: ['input', 'blur'],
      }
    ],
    equipmentName : [{
      required: true,
      message: '',
      trigger: ['input','blur'],
    },
      {
        whitespace: true,
        message: '不能只输入空格',
        trigger: ['input', 'blur'],
      }
    ],
    location : [{
      required: true,
      message: '',
      trigger: ['input','blur'],
    },
      {
        whitespace: true,
        message: '不能只输入空格',
        trigger: ['input', 'blur'],
      }
    ],
    softwareVersion : [{
      required: true,
      message: '',
      trigger: ['input','blur'],
    },
      {
        whitespace: true,
        message: '不能只输入空格',
        trigger: ['input', 'blur'],
      }
    ],
  })

  const elFormRef = ref()
  const elSearchFormRef = ref()

  // =========== 表格控制部分 ===========
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({})
  // 重置
  const onReset = () => {
    searchInfo.value = {}
    getTableData()
  }

  // 搜索
  const onSubmit = () => {
    elSearchFormRef.value?.validate(async(valid) => {
      if (!valid) return
      page.value = 1
      getTableData()
    })
  }

  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  // 修改页面容量
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  // 查询
  const getTableData = async() => {
    const table = await getEquipmentList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }

  getTableData()

  // ============== 表格控制部分结束 ===============

  // 获取需要的字典 可能为空 按需保留
  const setOptions = async () =>{
    OperationalstatusOptions.value = await getDictFunc('Operationalstatus')
    runningstateOptions.value = await getDictFunc('runningstate ')
    activestateOptions.value = await getDictFunc('activestate')
  }

  // 获取需要的字典 可能为空 按需保留
  setOptions()

  // 多选数据
  const multipleSelection = ref([])
  // 多选
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }

  // 删除行
  const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      deleteEquipmentFunc(row)
    })
  }

  // 多选删除
  const onDelete = async() => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
      multipleSelection.value.map(item => {
        IDs.push(item.ID)
      })
      const res = await deleteEquipmentByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
  }

  // 行为控制标记（弹窗内部需要增还是改）
  const type = ref('')

  // 更新行
  const updateEquipmentFunc = async(row) => {
    const res = await findEquipment({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteEquipmentFunc = async (row) => {
    const res = await deleteEquipment({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  }

  // 弹窗控制标记
  const dialogFormVisible = ref(false)

  // 打开弹窗
  const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
  }

  // 关闭弹窗
  const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
      equipmentNumber: '',
      equipmentName: '',
      location: '',
      operationalStatus: '',
      nunningState: '',
      sctiveState: '',
      softwareVersion: '',
      afterSalesPersonnel: '',
    }
  }

  // 弹窗确定
  const enterDialog = async () => {
    btnLoading.value = true
    elFormRef.value?.validate( async (valid) => {
      if (!valid) return btnLoading.value = false
      let res
      switch (type.value) {
        case 'create':
          res = await createEquipment(formData.value)
          break
        case 'update':
          res = await updateEquipment(formData.value)
          break
        default:
          res = await createEquipment(formData.value)
          break
      }
      btnLoading.value = false
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        closeDialog()
        getTableData()
      }
    })
  }

  const detailFrom = ref({})

  // 查看详情控制标记
  const detailShow = ref(false)

  // 打开详情弹窗
  const openDetailShow = () => {
    detailShow.value = true
  }

  // 打开详情
  const getDetails = async (row) => {
    // 打开弹窗
    const res = await findEquipment({ ID: row.ID })
    if (res.code === 0) {
      detailFrom.value = res.data
      openDetailShow()
    }
  }

  // 关闭详情弹窗
  const closeDetailShow = () => {
    detailShow.value = false
    detailFrom.value = {}
  }
</script>

<style scoped>
  /* 激活状态标签样式 */
  .status-tag {
    display: inline-flex;
    align-items: center;
    padding: 0 8px;
    height: 24px;
    line-height: 22px;
    border-radius: 12px;
    font-size: 12px;
    white-space: nowrap;
  }

  /* 状态圆点样式 */
  .status-dot {
    display: inline-block;
    width: 8px;
    height: 8px;
    border-radius: 50%;
    margin-right: 4px;
  }

  /* 状态颜色映射 */
  .dot-success { background-color: #52c41a; } /* 绿色 - 已激活 */
  .dot-warning { background-color: #faad14; } /* 橙色 - 未激活 */
  .dot-danger { background-color: #ff4d4f; }  /* 红色 - 已过期 */
  .dot-primary { background-color: #1890ff; } /* 蓝色 - 待激活 */
  .dot-default { background-color: #bfbfbf; } /* 默认 - 灰色 */

  /* 可选：状态背景色 */
  .bg-success-light { background-color: rgba(82, 196, 26, 0.1); }
  .bg-warning-light { background-color: rgba(250, 173, 20, 0.1); }
  .bg-danger-light { background-color: rgba(255, 77, 79, 0.1); }
  .bg-primary-light { background-color: rgba(24, 144, 255, 0.1); }

  /* 可选：如果需要自定义标签样式，可以在这里覆盖Element Plus默认样式 */
  .el-tag {
    font-size: 12px;
    padding: 0 6px;
    height: 22px;
    line-height: 20px;
  }
</style>