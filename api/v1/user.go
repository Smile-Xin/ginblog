package v1

import (
	"fmt"
	"ginbblog/dao"
	"ginbblog/model"
	"ginbblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 查重用户
func ExistUser(c *gin.Context) {
	//id, err := strconv.Atoi(c.Query("ID"))
	//if err != nil {
	//	fmt.Errorf("string to int fail:%s", err)
	//}
	//if dao.ExistUser(id) {
	//	c.JSON(200, gin.H{
	//		"err": 1001,
	//	})
	//}
}

// 查询用户
func QueryUser(c *gin.Context) {
	// id方式查询
	// 获取查询用户id
	//id, err := strconv.Atoi(c.Query("id"))
	//if err != nil {
	//	fmt.Printf("string to int fail:%s", err)
	//	return
	//}

	// name 方式查询
	name := c.Query("name")

	//fmt.Print(name)
	//fmt.Print(id)
	//user, code := dao.QueryUser(id)
	user, code := dao.QueryUser(name)
	c.JSON(200, gin.H{
		"state":   code,
		"data":    user,
		"message": errmsg.GetErrMsg(code),
	})
}

func QueryUserByModel(c *gin.Context) {

}

// 分页获取user列表
func GetUser(c *gin.Context) {
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	name := c.Query("name")
	if err != nil {
		fmt.Printf("string to int fail:%s", err)
	}

	userList, total, code := dao.GetUser(name, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    userList,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 添加用户
func AddUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Printf("bind user fail %s", err)
		return
	}
	code := dao.AddUser(&user)
	c.JSON(200, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 修改用户
func EditUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Printf("bind user fail %s", err)
		return
	}
	// 调用dao
	code := dao.EditUser(&user)
	// 返回信息
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})

}

// 删除用户
func DeleteUser(c *gin.Context) {
	name := c.Query("name")
	code := dao.DeleteUser(name)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}
