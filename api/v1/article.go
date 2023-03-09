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

// 创建文章
func AddArticle(c *gin.Context) {
	var article model.Article
	err := c.ShouldBindJSON(&article)
	var code uint
	if err != nil {
		fmt.Printf("bind article fail%s", err)
		code = errmsg.TRANSPORT_ERR
	} else {
		code = dao.AddArticle(article)
	}
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询文章
func QueryArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Printf("str to int fail: %s", err)
		return
	}
	code, article := dao.QueryArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询文章列表
func QueryArtList(c *gin.Context) {
	title := c.Query("title")
	cid, _ := strconv.Atoi(c.Query("cid"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	articleList, total, code := dao.QueryArtList(title, cid, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    articleList,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArticle(c *gin.Context) {
	title := c.Query("title")
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	articleList, total, code := dao.GetArticle(title, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    articleList,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑文章
func EditArticle(c *gin.Context) {
	var article model.Article
	var code uint
	err := c.ShouldBindJSON(&article)
	if err != nil {
		fmt.Printf("bind article fail:%s", err)
		return
	}
	code = dao.EditArticle(article)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})

}

// 删除文章
func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Printf("articleId to int fail:%s", err)
		return
	}
	code := dao.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}
