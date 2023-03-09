<template>
  <div>
    <el-card>
      <!-- 评论列表 -->
      <el-table :data="tableData" style="width: 100%" border="" stripe="">
        <el-table-column prop="ID" label="ID"></el-table-column>
        <el-table-column prop="CreatedAt" label="评论时间"></el-table-column>
        <el-table-column
          prop="article_title"
          label="评论文章"
        ></el-table-column>
        <el-table-column prop="username" label="评论者"></el-table-column>
        <el-table-column prop="content" label="评论内容"></el-table-column>
        <el-table-column label="审核状态">
          <template slot-scope="scope">
            <div>
              {{ scope.row.status == 1 ? "审核通过" : "未审核" }}
            </div>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              type="text"
              size="small"
              @click="commentCheck(scope.row.ID)"
              >过审</el-button
            >
            <el-button
              @click="commentUncheck(scope.row.ID)"
              type="text"
              size="small"
              >退回</el-button
            >
            <el-button
              @click="deleteComment(scope.row.ID)"
              type="text"
              size="small"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <!-- 页数 -->
      <div class="pag">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :page-sizes="[1, 2, 3, 6]"
          :page-size="searchData.pageSize"
          :total="total"
        >
        </el-pagination>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // 获取评论列表时的参数
      searchData: {
        pageSize: 6,
        pageNum: 1,
      },
      // 获取的评论列表
      tableData: [],
      // 获取的评论总数
      total: 0,
    };
  },
  created() {
    this.getCommentList();
  },
  methods: {
    async getCommentList() {
      const { data: res } = await this.$http.get("comment/get", {
        params: this.searchData,
      });
      this.tableData = res.data;
      console.log(res);
      this.total = res.total;
    },
    // 通过审核
    commentCheck(id) {
      console.log(id);
      this.$confirm("此操作将通过该评论, 是否继续？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          // 过审评论
          const { data: res_status } = await this.$http.get(
            `comment/info/${id}`
          );
          if (res_status.data.status === 1)
            return this.$message.error("该评论已处于显示状态，无需审核");
          const { data: res } = await this.$http.put(`checkcomment/${id}`, {
            status: 1,
          });
          if (res.status != 200) {
            return this.$message.error(res.message);
          } else {
            this.$message({
              type: "success",
              message: "操作成功!",
            });
          }
          this.getCommentList();
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消操作",
          });
        });
    },

    // 撤下评论
    commentUncheck(id) {
      this.$confirm("此操作将撤下该评论, 是否继续？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          // 撤下评论
          const { data: res_status } = await this.$http.get(
            `comment/info/${id}`
          );
          if (res_status.data.status === 2)
            return this.$message.error("该评论已处于未审核状态，无需撤下");
          const { data: res } = await this.$http.put(`uncheckcomment/${id}`, {
            status: 2,
          });
          if (res.status != 200) return this.$message.error(res.message);
          this.$message.success("评论已撤下");
          this.getCommentList();
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消操作",
          });
        });
    },

    // 删除评论
    deleteComment(id) {
      this.$confirm("此操作将永久删除该评论, 是否继续？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          // 删除评论
          const { data: res } = await this.$http.delete(`delcomment/${id}`);
          if (res.status != 200) return this.$message.error(res.message);
          this.$message.success("删除成功");
          this.getCommentList();
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    // table每页大小改变
    handleSizeChange(val) {
      this.searchData.pageSize = val;
      this.getArticleList();
    },
    // table页数改变
    handleCurrentChange(val) {
      this.searchData.pageNum = val;
      this.getArticleList();
    },
  },
};
</script>

<style></style>
