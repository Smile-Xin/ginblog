package v1

import (
	"fmt"
	"ginbblog/dao"
	"ginbblog/middleware"
	"ginbblog/model"
	"ginbblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 后台登录
func Login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Printf("bind user fail:%s", err)
		return
	}
	fmt.Println(user)

	// 调用loginDAO
	ok, code := dao.LoginUser(&user)
	// 登陆失败
	if !ok {
		fmt.Println("login fail")
		c.JSON(http.StatusOK, gin.H{
			"state":   code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	// 登陆成功,获取token
	token, code := middleware.JwtGenerateToken(user.UserName, time.Hour*7*24)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
	return
}

// 前台登录
func FrontLogin(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Printf("bind user fail:%s", err)
		return
	}
	fmt.Println(user)

	// 调用loginDAO
	ok, code := dao.FrontLogin(&user)
	// 登陆失败
	if !ok {
		fmt.Println("login fail")
		c.JSON(http.StatusOK, gin.H{
			"state":   code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    user.UserName,
		"id":      user.ID,
		"message": errmsg.GetErrMsg(code),
	})
	return

}
