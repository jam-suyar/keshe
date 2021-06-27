<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="身高:">
    <el-input v-model.number="formData.height" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="体重:">
    <el-input v-model.number="formData.weight" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="胸围:">
    <el-input v-model.number="formData.xunwei" clearable placeholder="请输入"/>
    </el-form-item>
    <el-form-item>
        <el-button type="primary" @click="save">保存</el-button>
        <el-button type="primary" @click="back">返回</el-button>
      </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
  createStutjts,
  updateStutjts,
  findStutjts
} from '@/api/stutjts' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Stutjts',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
            height: 0,
            weight: 0,
            xunwei: 0,
            
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findStutjts({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.restutjts
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createStutjts(this.formData)
          break
        case 'update':
          res = await updateStutjts(this.formData)
          break
        default:
          res = await createStutjts(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>