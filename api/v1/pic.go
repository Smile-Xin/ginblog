package v1

import (
	"fmt"
	"ginbblog/dao"
	"ginbblog/model"
	"ginbblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryPic(c *gin.Context) {
	name := c.Query("name")

	code, url := dao.QueryPic(name)

	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}

func AddPic(c *gin.Context) {
	var pic model.Pic
	err := c.ShouldBindJSON(&pic)
	if err != nil {
		fmt.Printf("bind pic fail:%s", err)
		return
	}
	code := dao.AddPic(&pic)

	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
		"data":    pic,
	})
}
