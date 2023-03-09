package model

import (
	"ginbblog/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null" json:"userName"`
	Password string `gorm:"type:varchar(64);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password, err = utils.ScryptPW(u.Password)
	return
}
