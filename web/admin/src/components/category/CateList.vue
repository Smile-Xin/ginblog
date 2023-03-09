<template>
    <div>
    <el-card>
      <el-row :gutter="15">
        <el-col :span="6">
          <el-input
            v-model="searchData.name"
            placeholder="分类名"
            @change="searchCategory()"
            clearable=""
          >
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </el-col>
        <el-col :span="6"
          ><el-button type="primary" @click="addCategoryFormVisible = true"
            >新增</el-button
          ></el-col
        >
      </el-row>

      <!-- // 添加分类的对话框 -->
      <el-dialog
        title="添加分类"
        :visible.sync="addCategoryFormVisible"
        width="500px"
        @close="addCategorylogClose"
      >
        <el-form :model="addCategoryData" :rules="addCategoryRules" ref="addCategoryRef">
          <el-form-item label="分类名" label-width="20%" prop="name">
            <el-input
              v-model="addCategoryData.name"
              autocomplete="off"
              v-on:keyup.enter.native="addCategory()"
            ></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="addCategoryFormVisible = false">取 消</el-button>
          <el-button type="primary" @click="addCategory()">确 定</el-button>
        </div>
      </el-dialog>
      <!-- 分类列表 -->
      <el-table :data="tableData" style="width: 100%" border="" stripe="">
        <el-table-column prop="ID" label="ID"></el-table-column>
        <el-table-column prop="name" label="分类名"></el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button type="text" size="small" @click="clickEdit(scope.row)"
              >编辑</el-button
            >
            <el-button @click="deleteCategory(scope.row)" type="text" size="small"
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
      <!--  编辑分类的对话框 -->
      <el-dialog
        title="修改分类"
        :visible.sync="editCategoryFormVisible"
        width="500px"
        @close="editCategorylogClose()"
      >
        <el-form :model="editCategoryData" :rules="editCategoryRules" ref="editCategoryRef">
          <el-form-item label="分类名" label-width="20%" prop="name">
            <el-input
              v-model="editCategoryData.name"
              autocomplete="off"
            ></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="editCategoryFormVisible = false">取 消</el-button>
          <el-button type="primary" @click="editCategory()">确 定</el-button>
        </div>
      </el-dialog>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // 获取分类列表时的参数
      searchData: {
        name: "",
        pageSize: 6,
        pageNum: 1,
      },
      // 获取的分类列表
      tableData: [],
      // 获取的分类总数
      total: 0,
      // 是否打开对话框
      addCategoryFormVisible: false,
      editCategoryFormVisible: false,
      // 添加分类的信息
      addCategoryData: {
        name: "",
      },
      // 修改分类的信息
      editCategoryData: {
        ID: 0,
        name: "",
        role: 1,
      },
      // 添加分类表单的规则
      addCategoryRules: {
        name: [
          { required: true, message: "请输入分类名", trigger: "blur" },
          {
            min: 1,
            max: 15,
            message: "长度在 1 到 15 个字符",
            trigger: "blur",
          },
        ],
      },
      // 修改分类表单的规则
      editCategoryRules: {
        name: [
          { required: true, message: "请输入分类名", trigger: "blur" },
          {
            min: 1,
            max: 15,
            message: "长度在 1 到 15 个字符",
            trigger: "blur",
          },
        ],
      },
    };
  },
  created() {
    this.getCategoryList();
  },
  methods: {
    async getCategoryList() {
      const { data: res } = await this.$http.get("category/get", {
        params: this.searchData,
      });
      this.tableData = res.data;
      console.log(res);
      this.total = res.total;
    },
    // 仅在输入框失去焦点或分类按下回车时触发
    searchCategory() {
      this.getCategoryList();
    },
    // 点击删除分类按钮函数
    deleteCategory(category) {
      this.$confirm("此操作将永久删除该文件, 是否继续？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          // 删除分类
          const res = await this.$http.delete("category/delete", {
            params: { name: category.name },
          });
          if (res.data.state != 200) {
            return this.$message.error(res.data.message);
          } else {
            this.$message({
              type: "success",
              message: "删除成功!",
            });
          }
          this.getCategoryList();
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    // 当对话框关闭时调用
    addCategorylogClose() {
      this.$refs.addCategoryRef.resetFields();
    },
    editCategorylogClose() {
      this.$refs.editCategoryRef.resetFields();
    },
    // 添加分类
    addCategory() {
      this.$refs.addCategoryRef.validate(async (valid) => {
        // 表单未通过验证
        if (!valid) {
          this.$message({
            message: "格式错误",
            type: "error",
            duration: 1750,
          });
        } else {
          // 通过验证，发送网络请求
          const res = await this.$http.post("category/add", this.addCategoryData);
          console.log(res);
          if (res.data.state != 200) {
            this.$message.error(res.data.message);
          } else {
            this.$message({
              message: "添加成功",
              type: "success",
              duration: 1750,
            });
            this.getCategoryList();
            this.addCategoryFormVisible = false;
          }
        }
      });
    },
    // 点击编辑按钮
    clickEdit(category) {
      console.log(category);
      this.editCategoryFormVisible = true;
      this.$nextTick(() => {
        this.editCategoryData.name = category.name;
        this.editCategoryData.ID = category.ID;
        this.editCategoryData.role = category.role;
      });
    },
    // 修改分类
    editCategory() {
      this.$refs.editCategoryRef.validate(async (valid) => {
        if (!valid) {
          // 表单验证未通过
          this.$message({
            message: "格式错误",
            type: "error",
            duration: 1750,
          });
        } else {
          // 验证通过，发送网络请求
          const { data: res } = await this.$http.post(
            "category/edit",
            this.editCategoryData
          );
          // 查看返回信息
          if (res.state != 200) {
            this.$message.error(res.message);
          } else {
            this.$message({
              message: "修改成功",
              type: "success",
              duration: 1750,
            });
            this.getCategoryList();
            this.editCategoryFormVisible = false;
          }
        }
      });
    },
    // table每页大小改变
    handleSizeChange(val) {
      this.searchData.pageSize = val;
      this.getCategoryList();
    },
    // table页数改变
    handleCurrentChange(val) {
      this.searchData.pageNum = val;
      this.getCategoryList();
    },
  },
};
</script>

<style>

</style>