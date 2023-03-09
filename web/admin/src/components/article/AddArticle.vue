<template>
  <div>
    <el-card>
      <el-form
        ref="article"
        :model="article"
        label-width="80px"
        :rules="articleRules"
      >
        <el-form-item label="文章标题" prop="title">
          <el-input
            v-model="article.title"
            placeholder=""
            style="width: 350px"
          ></el-input>
        </el-form-item>

        <el-form-item label="文章类型" prop="cid">
          <el-select
            placeholder="选择文章类型"
            v-model="article.cid"
            @change="changecategory"
          >
            <el-option
              v-for="item in categories"
              :key="item.ID"
              :label="item.name"
              :value="item.ID"
            >
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="文章简介">
          <el-input
            type="textarea"
            v-model="article.desc"
            placeholder=""
            style="width: 350px"
          ></el-input>
        </el-form-item>

        <el-form-item label="文章内容" prop="content">
          <el-input
            type="textarea"
            v-model="article.content"
            placeholder=""
            style="width: 350px"
          ></el-input>
        </el-form-item>
        <div class="but">
          <el-button type="primary" @click="addArticle('article')"
            >提交</el-button
          >
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  props: ["id"],
  data() {
    return {
      article: {
        Category: {},
        id: 0,
        title: "",
        cid: "",
        desc: "",
        content: "",
      },
      categories: [],
      articleRules: {
        title: [{ required: true, message: "标题不可为空", trigger: "blur" }],
        cid: [
          { required: true, message: "请选择文章分类", trigger: "change" }
        ],
        content: [
          { required: true, message: "文章内容不可为空", trigger: "blur" }
        ],
      },
    };
  },
  methods: {
    // 改变分类
    changecategory(value) {
      this.article.cid = value;
      this.article.Category.id = value;
    },
    // 获取文章分类列表
    async getCategoryList() {
      const { data: res } = await this.$http.get("category/get", {
        params: this.searchData,
      });
      if (res.state != 200) {
        return this.$message.error(res.data.$message);
      }
      this.categories = res.data;
    },
    // 获取需要修改的文章
    async getArticle(id) {
      const { data: res } = await this.$http.get("article/query", {
        params: { id: id },
      });
      if (res.state != 200) {
        return this.$message.error(res.data.$message);
      }
      this.article = res.data;
    },
    // 添加或者修改文章
    addArticle(formName) {
      this.$refs[formName].validate(async (valid) => {
        if (valid) {
          if (this.id) {
            // 修改文章
            const { data: res } = await this.$http.post(
              "article/edit",
              this.article
            );
            if (res.state != 200) {
              this.$message.error(res.message);
            } else {
              this.$message({
                type: "success",
                message: "提交成功",
              });
            }
          } else {
            // 添加文章
            const { data: res } = await this.$http.post(
              "article/add",
              this.article
            );
            if (res.state != 200) {
              this.$message.error(res.message);
            } else {
              this.$message({
                type: "success",
                message: "提交成功",
              });
            }
          }
          this.$router.push("/admin/artList");
        } else {
          this.$message({
            message: "格式错误",  
            type: "error",
            duration: 1750,
          });
        }
      });
    },
  },
  created() {
    this.getCategoryList();
    if (this.id) {
      this.getArticle(this.id);
    }
  },
};
</script>

<style>
.but {
  display: flex;
  justify-content: right;
  margin: 10px 50px;
}
</style>
