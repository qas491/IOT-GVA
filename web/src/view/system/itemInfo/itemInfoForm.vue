
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="编码:" prop="code">
    <el-input v-model="formData.code" :clearable="false" placeholder="请输入编码" />
</el-form-item>
        <el-form-item label="名称:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入名称" />
</el-form-item>
        <el-form-item label="规格:" prop="specification">
    <el-input v-model="formData.specification" :clearable="false" placeholder="请输入规格" />
</el-form-item>
        <el-form-item label="周期:" prop="days">
    <el-input v-model.number="formData.days" :clearable="false" placeholder="请输入周期" />
</el-form-item>
        <el-form-item label="使用次数:" prop="usageCount">
    <el-input v-model.number="formData.usageCount" :clearable="false" placeholder="请输入使用次数" />
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
  createItemInfo,
  updateItemInfo,
  findItemInfo
} from '@/api/system/itemInfo'

defineOptions({
    name: 'ItemInfoForm'
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
            code: '',
            name: '',
            specification: '',
            days: 0,
            usageCount: 0,
        })
// 验证规则
const rule = reactive({
               code : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               specification : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               days : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               usageCount : [{
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
      const res = await findItemInfo({ ID: route.query.id })
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
               res = await createItemInfo(formData.value)
               break
             case 'update':
               res = await updateItemInfo(formData.value)
               break
             default:
               res = await createItemInfo(formData.value)
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
