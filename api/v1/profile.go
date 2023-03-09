package v1

import (
	"ginbblog/dao"
	"ginbblog/model"
	"ginbblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func EditProfile(c *gin.Context) {
	var profile model.Profile

	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&profile)

	code := dao.EditProfile(id, &profile)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	profile, code := dao.GetProfile(id)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
		"data":    profile,
	})
}
