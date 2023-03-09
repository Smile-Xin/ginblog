package model

import "gorm.io/gorm"

type Pic struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20)" json:"userName"`
	Url      string `gorm:"type:varchar(200)" json:"url"`
}
