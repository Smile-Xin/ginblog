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

// 获取分类
func GetCategory(c *gin.Context) {
	name := c.Query("name")
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	// 调用dao方法查询
	categoryList, total, code := dao.GetCategory(pageSize, pageNum, name)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    categoryList,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 增加分类
func AddCategory(c *gin.Context) {
	var category model.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		fmt.Printf("bind category fail:%s", err)
		return
	}
	code := dao.AddCategory(category)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    category,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类
func QueryCategory(c *gin.Context) {
	name := c.Query("name")
	// 调用dao方法
	category, code := dao.QueryCategory(name)
	// 返回消息
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    category,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类
func EditCategory(c *gin.Context) {
	var category model.Category
	var code uint
	err := c.ShouldBindJSON(&category)
	if err != nil {
		fmt.Printf("bind category fail:%s", err)
		code = errmsg.TRANSPORT_ERR
	} else {
		code = dao.EditCategory(category)
	}

	// 返回消息
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    category,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCategory(c *gin.Context) {
	name := c.Query("name")
	code := dao.DeleteCategory(name)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}
