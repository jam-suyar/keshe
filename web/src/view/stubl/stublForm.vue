<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="学号:">
    <el-input v-model.number="formData.xuehao" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="姓名:">
    <el-input v-model="formData.name" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="系别:">
    <el-input v-model="formData.xibie" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="性别:">
    <el-input v-model="formData.xingbie" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="年龄:">
    <el-input v-model.number="formData.age" clearable placeholder="请输入"/>
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
  createStubl,
  updateStubl,
  findStubl
} from '@/api/stubl' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Stubl',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
            xuehao: 0,
            name: '',
            xibie: '',
            xingbie: '',
            age: 0,
            
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findStubl({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.restubl
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
          res = await createStubl(this.formData)
          break
        case 'update':
          res = await updateStubl(this.formData)
          break
        default:
          res = await createStubl(this.formData)
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