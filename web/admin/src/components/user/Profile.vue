<template>
  <el-card>
    <el-form labelAlign="left" :label-col="{ span: 2 }" :wrapper-col="{ span: 12 }">
      <el-form-item label="作者名称">
        <el-input style="width: 300px" v-model="profileInfo.name"></el-input>
      </el-form-item>

      <el-form-item label="个人简介">
        <el-input type="textarea" v-model="profileInfo.desc"></el-input>
      </el-form-item>

      <el-form-item label="网站备案号">
        <el-input style="width: 300px" v-model="profileInfo.icp_record"></el-input>
      </el-form-item>

      <el-form-item label="QQ号码">
        <el-input style="width: 300px" v-model="profileInfo.qq_chat"></el-input>
      </el-form-item>

      <el-form-item label="微信">
        <el-input style="width: 300px" v-model="profileInfo.wechat"></el-input>
      </el-form-item>

      <el-form-item label="微博">
        <el-input style="width: 300px" v-model="profileInfo.weibo"></el-input>
      </el-form-item>

      <el-form-item label="B站地址">
        <el-input style="width: 300px" v-model="profileInfo.bili"></el-input>
      </el-form-item>

      <el-form-item label="Email">
        <el-input style="width: 300px" v-model="profileInfo.email"></el-input>
      </el-form-item>
<!-- 
      <el-form-item label="头像">
        <a-upload
          listType="picture"
          name="file"
          :action="upUrl"
          :headers="headers"
          @change="avatarChange"
        >
          <a-button style="margin-right:10px">
            <a-icon type="upload" />点击上传
          </a-button>

          <template v-if="profileInfo.avatar">
            <img :src="profileInfo.avatar" style="width: 120px; height: 100px" />
          </template>
        </a-upload>
      </el-form-item>

      <el-form-item label="头像背景图">
        <a-upload
          listType="picture"
          name="file"
          :action="upUrl"
          :headers="headers"
          @change="imgChange"
        >
          <a-button style="margin-right:10px">
            <a-icon type="upload" />点击上传
          </a-button>

          <template v-if="profileInfo.img">
            <img :src="profileInfo.img" style="width: 120px; height: 100px" />
          </template>
        </a-upload>
      </el-form-item> -->

      <el-form-item>
        <el-button type="danger" style="margin-right: 15px" @click="updateProfile">更新</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script>
// import { Url } from '../../plugin/http'

export default {
  data() {
    return {
      profileInfo: {
        id: 1,
        name: '',
        desc: '',
        qq_chat: '',
        wechat: '',
        weibo: '',
        bili: '',
        email: '',
        img: '',
        avatar: '',
        icp_record: '',
      },
      // upUrl: Url + 'upload',
      // headers: {},
    }
  },
  created() {
    this.getProfileInfo()
    // this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
  },
  methods: {
    // 获取个人设置
    async getProfileInfo() {
      const { data: res } = await this.$http.get(`admin/profile/${this.profileInfo.id}`)
      console.log(res)
      if (res.state !== 200) {
        // if (res.status === 1004 || 1005 || 1006 || 1007) {
        //   window.sessionStorage.clear()
        //   this.$router.push('/login')
        // }
        this.$message.error(res.message)
      }
      this.profileInfo = res.data
    },

    // // 上传头像
    // avatarChange(info) {
    //   if (info.file.status !== 'uploading') {
    //   }
    //   if (info.file.status === 'done') {
    //     this.$message.success(`图片上传成功`)
    //     const imgUrl = info.file.response.url
    //     this.profileInfo.avatar = imgUrl
    //   } else if (info.file.status === 'error') {
    //     this.$message.error(`图片上传失败`)
    //   }
    // },

    // // 上传头像背景图
    // imgChange(info) {
    //   if (info.file.status !== 'uploading') {
    //   }
    //   if (info.file.status === 'done') {
    //     this.$message.success(`图片上传成功`)
    //     const imgUrl = info.file.response.url
    //     this.profileInfo.img = imgUrl
    //   } else if (info.file.status === 'error') {
    //     this.$message.error(`图片上传失败`)
    //   }
    // },

    // 更新
    async updateProfile() {
      const { data: res } = await this.$http.post(`profile/edit/${this.profileInfo.id}`, this.profileInfo)
      if (res.state !== 200) return this.$message.error(res.message)
      this.$message.success(`个人信息更新成功`)
      this.$router.push('/admin/index')
    },
  },
}
</script>

<style scoped>
.upBtn {
  margin-right: 10px;
}
</style>