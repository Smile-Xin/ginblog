<template>
  <div>
    <el-card>
      <el-row :gutter="15">
        <el-col :span="6">
          <el-input
            v-model="searchData.name"
            placeholder="用户名"
            @change="searchUser()"
            clearable=""
          >
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </el-col>
        <el-col :span="6"
          ><el-button type="primary" @click="addUserFormVisible = true"
            >新增</el-button
          ></el-col
        >
      </el-row>

      <!-- // 添加用户的对话框 -->
      <el-dialog
        title="添加用户"
        :visible.sync="addUserFormVisible"
        width="500px"
        @close="addUserlogClose"
      >
        <el-form :model="addUserData" :rules="addUserRules" ref="addUserRef">
          <el-form-item label="用户名" label-width="20%" prop="userName">
            <el-input
              v-model="addUserData.userName"
              autocomplete="off"
              v-on:keyup.enter.native="addUser()"
            ></el-input>
          </el-form-item>
          <el-form-item label="密码" label-width="20%" prop="password">
            <el-input
              v-on:keyup.enter.native="addUser()"
              v-model="addUserData.password"
              type="password"
              autocomplete="off"
            ></el-input>
          </el-form-item>
          <el-form-item label="用户权限" label-width="20%" prop="role">
            <el-select v-model="addUserData.role" placeholder="选择用户权限">
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
          <el-button @click="addUserFormVisible = false">取 消</el-button>
          <el-button type="primary" @click="addUser()">确 定</el-button>
        </div>
      </el-dialog>
      <!-- 用户列表 -->
      <el-table :data="tableData" style="width: 100%" border="" stripe="">
        <el-table-column prop="ID" label="ID"></el-table-column>
        <el-table-column prop="userName" label="用户名"></el-table-column>
        <el-table-column label="角色">
          <template slot-scope="scope">
            <div>
              {{ scope.row.role >= 1 ? "读者" : "管理员" }}
            </div>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button type="text" size="small" @click="clickEdit(scope.row)"
              >编辑</el-button
            >
            <el-button @click="deleteUser(scope.row)" type="text" size="small"
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
      <!--  编辑用户的对话框 -->
      <el-dialog
        title="修改用户"
        :visible.sync="editUserFormVisible"
        width="500px"
        @close="editUserlogClose()"
      >
        <el-form :model="editUserData" :rules="editUserRules" ref="editUserRef">
          <el-form-item label="用户名" label-width="20%" prop="userName">
            <el-input
              v-model="editUserData.userName"
              autocomplete="off"
            ></el-input>
          </el-form-item>
          <el-form-item label="用户权限" label-width="20%" prop="role">
            <el-select v-model="editUserData.role" placeholder="选择用户权限">
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
          <el-button @click="editUserFormVisible = false">取 消</el-button>
          <el-button type="primary" @click="editUser()">确 定</el-button>
        </div>
      </el-dialog>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // 获取用户列表时的参数
      searchData: {
        name: "",
        pageSize: 6,
        pageNum: 1,
      },
      // 获取的用户列表
      tableData: [],
      // 获取的用户总数
      total: 0,
      // 是否打开对话框
      addUserFormVisible: false,
      editUserFormVisible: false,
      // 添加用户的信息
      addUserData: {
        userName: "",
        password: "",
        role: 1,
      },
      // 修改用户的信息
      editUserData: {
        ID: 0,
        userName: "",
        role: 1,
      },
      // 用户权限选择的选项
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
      // 添加用户表单的规则
      addUserRules: {
        userName: [
          { required: true, message: "请输入用户名", trigger: "blur" },
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
      // 修改用户表单的规则
      editUserRules: {
        userName: [
          { required: true, message: "请输入用户名", trigger: "blur" },
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
    this.getUserList();
  },
  methods: {
    async getUserList() {
      const { data: res } = await this.$http.get("user/get", {
        params: this.searchData,
      });
      this.tableData = res.data;
      console.log(res);
      this.total = res.total;
    },
    // 仅在输入框失去焦点或用户按下回车时触发
    searchUser() {
      this.getUserList();
    },
    // 点击删除用户按钮函数
    deleteUser(user) {
      this.$confirm("此操作将永久删除该文件, 是否继续？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          // 删除用户
          const res = await this.$http.delete("user/delete", {
            params: { name: user.userName },
          });
          if (res.data.state != 200) {
            return this.$message.error(res.data.message);
          } else {
            this.$message({
              type: "success",
              message: "删除成功!",
            });
          }
          this.getUserList();
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    // 当对话框关闭时调用
    addUserlogClose() {
      this.$refs.addUserRef.resetFields();
    },
    editUserlogClose() {
      this.$refs.editUserRef.resetFields();
    },
    // 添加用户
    addUser() {
      this.$refs.addUserRef.validate(async (valid) => {
        // 表单未通过验证
        if (!valid) {
          this.$message({
            message: "格式错误",
            type: "error",
            duration: 1750,
          });
        } else {
          // 通过验证，发送网络请求
          const res = await this.$http.post("user/add", this.addUserData);
          console.log(res);
          if (res.data.state != 200) {
            this.$message.error(res.data.message);
          } else {
            this.$message({
              message: "添加成功",
              type: "success",
              duration: 1750,
            });
            this.getUserList();
            this.addUserFormVisible = false;
          }
        }
      });
    },
    // 点击编辑按钮
    clickEdit(user) {
      console.log(user);
      this.editUserFormVisible = true;
      this.$nextTick(() => {
        this.editUserData.userName = user.userName;
        this.editUserData.ID = user.ID;
        this.editUserData.role = user.role;
      });
    },
    // 修改用户
    editUser() {
      this.$refs.editUserRef.validate(async (valid) => {
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
            "user/edit",
            this.editUserData
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
            this.getUserList();
            this.editUserFormVisible = false;
          }
        }
      });
    },
    // table每页大小改变
    handleSizeChange(val) {
      this.searchData.pageSize = val;
      this.getUserList();
    },
    // table页数改变
    handleCurrentChange(val) {
      this.searchData.pageNum = val;
      this.getUserList();
    },
  },
};
</script>

<style></style>
