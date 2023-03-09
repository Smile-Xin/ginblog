package dao

import (
	"ginbblog/model"
	"ginbblog/utils/errmsg"
)

// GetProfile 获取个人信息设置
func GetProfile(id int) (model.Profile, uint) {
	var profile model.Profile
	err = db.Where("ID = ?", id).First(&profile).Error
	if err != nil {
		return profile, errmsg.DATABASE_WRITE_FAIL
	}
	return profile, errmsg.SUCCESS
}

// EditProfile 更新个人信息设置
func EditProfile(id int, data *model.Profile) uint {
	var profile model.Profile
	err = db.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.DATABASE_WRITE_FAIL
	}
	return errmsg.SUCCESS
}
