<template>
  <div>
    <el-card>
      <el-row :gutter="15">
        <el-col :span="6">
          <el-input
            v-model="searchData.title"
            placeholder="文章名"
            @change="searchArticle()"
            clearable=""
          >
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </el-col>
        <el-col :span="6"
          ><el-button type="primary" @click="addArticle()"
            >新增</el-button
          ></el-col
        >
      </el-row>


      <!-- 文章列表 -->
      <el-table :data="tableData" style="width: 100%" border="" stripe="">
        <el-table-column prop="ID" label="ID"></el-table-column>
        <el-table-column prop="title" label="文章标题"></el-table-column>
        <el-table-column prop="desc" label="文章描述"></el-table-column>
        <el-table-column prop="Category.name" label="文章分类"></el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button type="text" size="small" @click="$router.push(`addArticle/${scope.row.ID}`)"
              >编辑</el-button
            >
            <el-button @click="deleteArticle(scope.row)" type="text" size="small"
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
      <!--  编辑文章的对话框 -->
      <el-dialog
        title="修改文章"
        :visible.sync="editArticleFormVisible"
        width="500px"
        @close="editArticlelogClose()"
      >
        <el-form :model="editArticleData" :rules="editArticleRules" ref="editArticleRef">
          <el-form-item label="文章名" label-width="20%" prop="title">
            <el-input
              v-model="editArticleData.title"
              autocomplete="off"
            ></el-input>
          </el-form-item>
          <el-form-item label="文章权限" label-width="20%" prop="role">
            <el-select v-model="editArticleData.role" placeholder="选择文章权限">
              <el-option
                v-for="(item, index) in bannerRole"
                :key="index"
                :label="item.name"
                :value="item.id * 1"
              ></el-option>
            </el-select>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="editArticleFormVisible = false">取 消</el-button>
          <el-button type="primary" @click="editArticle()">确 定</el-button>
        </div>
      </el-dialog>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // 获取文章列表时的参数
      searchData: {
        title: "",
        pageSize: 6,
        pageNum: 1,
      },
      // 获取的文章列表
      tableData: [],
      // 获取的文章总数
      total: 0,
      // 是否打开对话框
      addArticleFormVisible: false,
      editArticleFormVisible: false,
      // 添加文章的信息
      addArticleData: {
        title: "",
        password: "",
        role: 1,
      },
      // 修改文章的信息
      editArticleData: {
        ID: 0,
        title: "",
        role: 1,
      },
      // 文章权限选择的选项
      bannerRole: [
        {
          id: 0,
          name: "管理员",
        },
        {
          id: 1,
          name: "读者",
        },
      ],
      // 添加文章表单的规则
      addArticleRules: {
        title: [
          { required: true, message: "请输入文章名", trigger: "blur" },
          {
            min: 3,
            max: 10,
            message: "长度在 3 到 10 个字符",
            trigger: "blur",
          },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          {
            min: 3,
            max: 15,
            message: "长度在 3 到 15 个字符",
            trigger: "blur",
          },
        ],
      },
      // 修改文章表单的规则
      editArticleRules: {
        title: [
          { required: true, message: "请输入文章名", trigger: "blur" },
          {
            min: 3,
            max: 10,
            message: "长度在 3 到 10 个字符",
            trigger: "blur",
          },
        ],
      },
    };
  },
  created() {
    this.getArticleList();
  },
  methods: {
    async getArticleList() {
      const { data: res } = await this.$http.get("article/get", {
        params: this.searchData,
      });
      this.tableData = res.data;
      console.log(res);
      this.total = res.total;
    },
    // 仅在输入框失去焦点或文章按下回车时触发
    searchArticle() {
      this.getArticleList();
    },
    deleteArticle(article) {
    // 点击删除文章按钮函数
      console.log(article)
      this.$confirm("此操作将永久删除该文件, 是否继续？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          // 删除文章
          const res = await this.$http.delete("article/delete", {
            params: { id: article.ID },
          });
          if (res.data.state != 200) {
            return this.$message.error(res.data.message);
          } else {
            this.$message({
              type: "success",
              message: "删除成功!",
            });
          }
          this.getArticleList();
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },

    // 添加文章
    addArticle() {
     this.$router.push("addArticle")
    },
    // 点击编辑按钮
    clickEdit(article) {
      console.log(article);
      this.editArticleFormVisible = true;
      this.$nextTick(() => {
        this.editArticleData.title = article.title;
        this.editArticleData.ID = article.ID;
        this.editArticleData.role = article.role;
      });
    },
    // 修改文章
    editArticle(article) {
      // this.$router.push({
      //   name: "addArticle",
        
      // })
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
