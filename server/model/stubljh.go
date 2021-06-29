// 自动生成模板Stubljh
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Stubljh struct {
	global.GVA_MODEL
	Zhenduan string `json:"zhenduan" form:"zhenduan" gorm:"column:zhenduan;comment:诊断;type:varchar(1000);"`
}

func (Stubljh) TableName() string {
	return "stubljh"
}
