<template>
  <div class="container">
    <div class="loginBox">

      <el-form :model="loginForm" ref="loginForm" :rules="loginRules" label-width="80px" class="loginForm">

        <el-form-item label="用户名" size="normal" prop="userName">
          <el-input v-on:keyup.enter.native="login('loginForm')" v-model="loginForm.userName" placeholder="userName" size="normal" prefix-icon="el-icon-user"></el-input>
        </el-form-item>
        
        <el-form-item label="密码" size="normal" prop="password">
          <el-input v-on:keyup.enter.native="login('loginForm')" v-model="loginForm.password" type="password" placeholder="password" size="normal" prefix-icon="el-icon-lock"></el-input>
        </el-form-item>

        <el-form-item label="" size="normal">
          <el-button type="primary" @click="login('loginForm')">登录</el-button>
          <el-button type="" @click="resetForm('loginForm')">取消</el-button>
        </el-form-item>
        
      </el-form>

    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loginForm:{
        userName:'',
        password:''
      },
      loginRules:{
        userName:[
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
        ],
        password:[
          {required: true, message:'请输入密码', trigger: 'blur'},
          { min: 3, max: 15, message: '长度在 3 到 15 个字符', trigger: 'blur' }
        ]
      }
    };
  },
  methods: {  
      resetForm(formName){
      //  this.$refs.loginForm.resetFields();
      this.$refs[formName].resetFields();
      },
      login(formName){
        this.$refs[formName].validate(async valid => {
          if (valid) {
            // console.log(this.loginForm);
            const {data:res} = await this.$http.post("login", this.loginForm)
            if (res.state != 200){return this.$message.error(res.message)}
            window.sessionStorage.setItem("token", res.token)
            this.$message({
              message:"登陆成功",
              type:'success',
              duration: 1750
            })
            this.$router.push("admin/index")
          } else {
            this.$message({
              message:"格式错误",
              type:'error',
              duration: 1750
          })
          }
        })
      }
    },
 
  
}
</script>

<style scope> 
.container{
  height: 100%;
  background-color: #c2c2c2;

  /* flex布局兼容性较差，这里仅作参考 */
  /* display: flex;
  justify-content: center;
  align-items: center; */

}
.loginBox{
  width: 450px;
  height: 300px;
  background-color: #fff;

  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 5px;
}

.loginForm{
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  box-sizing: border-box;

}

</style>