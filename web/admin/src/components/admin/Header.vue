<template>
  <el-row class="header">
    <el-col class="logo">
      <el-link type="primary" :underline="false"
        ><span class="title blog" @click="logo">Blog</span></el-link
      >
    </el-col>
    <el-col class="btn">
      <el-button type="danger" size="small" @click="logout">退出用户</el-button>
    </el-col>
  </el-row>
</template>

<script>
export default {
  methods: {
    logout() {
      window.sessionStorage.clear("token");
      this.$router.push("/login");
    },
    logo() {
      this.$router.push("/admin/index");
    },
  },
};
</script>

<style>
.header {
  height: 100%;
  width: 100%;
  display: flex;
}
.btn {
  display: flex;
  align-items: center;
  justify-content: right;
}
.logo {
  display: flex;
  align-items: center;
}

.title {
  font-size: 40px;
  font-weight: bold;
  color: #282c34;
}
</style>
