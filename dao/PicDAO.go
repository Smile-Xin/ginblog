package dao

import (
	"fmt"
	. "ginbblog/model"
	"ginbblog/utils/errmsg"
)

func QueryPic(name string) (code uint, url []string) {
	code = errmsg.SUCCESS
	var pic []Pic
	err := db.Where("`user_name` = ?", name).Find(&pic).Error
	if err != nil {
		fmt.Printf("query pic fail:%s", err)
		code = errmsg.DATABASE_WRITE_FAIL
	}
	for i := 0; i < len(pic); i++ {
		url = append(url, pic[i].Url)
	}

	return
}

func AddPic(pic *Pic) (code uint) {
	code = errmsg.SUCCESS
	err := db.Create(&pic).Error
	if err != nil {
		fmt.Printf("creat pic fail:%s", err)
		code = errmsg.DATABASE_WRITE_FAIL
	}

	return code
}
