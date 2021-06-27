// 自动生成模板Stutjts
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Stutjts struct {
	global.GVA_MODEL
	Height int `json:"height" form:"height" gorm:"column:height;comment:身高;type:int;size:200;"`
	Weight int `json:"weight" form:"weight" gorm:"column:weight;comment:体重;type:int;size:300;"`
	Xunwei int `json:"xunwei" form:"xunwei" gorm:"column:xunwei;comment:胸围;type:int;size:200;"`
}

func (Stutjts) TableName() string {
	return "stutjts"
}
