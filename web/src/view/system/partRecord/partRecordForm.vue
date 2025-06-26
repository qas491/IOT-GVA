
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="批量导出:" prop="exportBatch">
    <el-switch v-model="formData.exportBatch" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="型号:" prop="partModel">
    <el-input v-model="formData.partModel" :clearable="false" placeholder="请输入型号" />
</el-form-item>
        <el-form-item label="编号:" prop="partNumber">
    <el-input v-model="formData.partNumber" :clearable="false" placeholder="请输入编号" />
</el-form-item>
        <el-form-item label="次数:" prop="usageCount">
    <el-input v-model.number="formData.usageCount" :clearable="false" placeholder="请输入次数" />
</el-form-item>
        <el-form-item label="天数:" prop="usageDays">
    <el-input v-model.number="formData.usageDays" :clearable="false" placeholder="请输入天数" />
</el-form-item>
        <el-form-item label="上机日:" prop="firstUsageDate">
    <el-date-picker v-model="formData.firstUsageDate" type="date" style="width:100%" placeholder="选择日期" :clearable="false" />
</el-form-item>
        <el-form-item label="同步日:" prop="syncDate">
    <el-date-picker v-model="formData.syncDate" type="date" style="width:100%" placeholder="选择日期" :clearable="false" />
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
  createPartRecord,
  updatePartRecord,
  findPartRecord
} from '@/api/system/partRecord'

defineOptions({
    name: 'PartRecordForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            exportBatch: false,
            partModel: '',
            partNumber: '',
            usageCount: 0,
            usageDays: 0,
            firstUsageDate: new Date(),
            syncDate: new Date(),
        })
// 验证规则
const rule = reactive({
               partModel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               partNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPartRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
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
               res = await createPartRecord(formData.value)
               break
             case 'update':
               res = await updatePartRecord(formData.value)
               break
             default:
               res = await createPartRecord(formData.value)
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
