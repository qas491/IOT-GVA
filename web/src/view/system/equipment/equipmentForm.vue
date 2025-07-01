<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="设备编号:" prop="equipmentNumber">
    <el-input v-model="formData.equipmentNumber" :clearable="false" placeholder="请输入设备编号" />
</el-form-item>
        <el-form-item label="设备名称:" prop="equipmentName">
    <el-input v-model="formData.equipmentName" :clearable="false" placeholder="请输入设备名称" />
</el-form-item>
        <el-form-item label="所在位置:" prop="location">
          <el-cascader
            v-model="areaValue"
            :options="areaOptions"
            :props="areaProps"
            clearable
            filterable
            placeholder="请选择省/市/区/镇/村"
            @change="handleAreaChange"
            style="width: 100%"
          />
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
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createEquipment,
  updateEquipment,
  findEquipment
} from '@/api/system/equipment'

defineOptions({
    name: 'EquipmentForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import areaData from 'china-area-data';

const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
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
// 验证规则
const rule = reactive({
               equipmentNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               equipmentName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               location : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               softwareVersion : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

const areaValue = ref([])
const areaOptions = ref([])
const areaProps = {
  checkStrictly: false,
  emitPath: true,
  value: 'value',
  label: 'label',
  children: 'children',
}

function formatAreaData(data, parentCode = '86', level = 1) {
  const result = []
  const current = data[parentCode]
  if (!current) return result
  for (const code in current) {
    const item = {
      value: code,
      label: current[code],
    }
    const children = formatAreaData(data, code, level + 1)
    if (children.length > 0) {
      item.children = children
    }
    result.push(item)
  }
  return result
}

areaOptions.value = formatAreaData(areaData)

function handleAreaChange(val) {
  // 取出label路径
  let labels = []
  let options = areaOptions.value
  for (let i = 0; i < val.length; i++) {
    const code = val[i]
    const node = options.find(item => item.value === code)
    if (node) {
      labels.push(node.label)
      options = node.children || []
    } else {
      break
    }
  }
  formData.value.location = labels.join('/')
}

// 回显处理
function setAreaValueByLocation(location) {
  if (!location) return
  const labels = location.split('/')
  let codes = []
  let options = areaOptions.value
  for (let i = 0; i < labels.length; i++) {
    const node = options.find(item => item.label === labels[i])
    if (node) {
      codes.push(node.value)
      options = node.children || []
    } else {
      break
    }
  }
  areaValue.value = codes
}

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEquipment({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        setAreaValueByLocation(res.data.location)
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    OperationalstatusOptions.value = await getDictFunc('Operationalstatus')
    runningstateOptions.value = await getDictFunc('runningstate ')
    activestateOptions.value = await getDictFunc('activestate')
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
