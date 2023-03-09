<template>
  <el-aside width="200px">
    <el-row class="tac">
      <el-menu
        default-active=""
        router=""
        mode="vertical"
        class="el-menu-vertical-demo"
        background-color="#d3dce6"
        text-color="#282c34"
        active-text-color="#004d40"
      >
        <!-- 第一个标签  仪表盘 -->
        <el-menu-item index="1" route="/admin/index">
          <template slot="title">
            <i class="el-icon-pie-chart"></i>
            <span>仪表盘</span>
          </template>
        </el-menu-item>

          <!-- 第一个标签分类  文章管理 -->
        <el-submenu index="2">
          <template slot="title">
            <i class="el-icon-folder-opened"></i>
            <span>文章管理</span>
          </template>

          <!-- 第二个标签 添加文章-->
          <el-menu-item index="2-1" route="/admin/addArticle">
            <i class="el-icon-edit"></i>
            <span>写文章</span>
          </el-menu-item>
          <!-- 第三个标签  文章列表 -->
          <el-menu-item index="2-2" route="/admin/artList">
            <i class="el-icon-document-copy"></i>
            <span>文章列表</span>
          </el-menu-item>
        </el-submenu>

        <!-- 第二个大标签（第四个标签） 文章分类 -->
        <el-menu-item index="3" route="/admin/cateList">
          <i class="el-icon-menu"></i>
          <span slot="title">文章分类</span>
        </el-menu-item>
        <!-- 第三个大标签（第五个标签） 用户列表  -->
        <el-menu-item index="4" route="/admin/userList">
          <i class="el-icon-document"></i>
          <span slot="title">用户列表</span>
        </el-menu-item>
        <!-- 评论管理  -->
        <el-menu-item index="5" route="/admin/comment">
          <i class="el-icon-document"></i>
          <span slot="title">评论管理</span>
        </el-menu-item>
         <!-- 评论管理  -->
        <el-menu-item index="6" route="/admin/profile">
          <i class="el-icon-document"></i>
          <span slot="title">个人简介</span>
        </el-menu-item>
      </el-menu>
    </el-row>
  </el-aside>
</template>

<script>
import "element-ui/lib/theme-chalk/display.css";
export default {
  data() {
    return {
      def:"2-1",
    }
  },
};
</script>

<style></style>
