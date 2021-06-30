// 自动生成模板Stubljh
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Stubljh struct {
	global.GVA_MODEL
	Xuehao   int    `json:"xuehao" form:"xuehao" gorm:"column:xuehao;comment:学号;type:int;size:100;"`
	Name     int    `json:"name" form:"name" gorm:"column:name;comment:姓名;type:int;size:100;"`
	Zhenduan string `json:"zhenduan" form:"zhenduan" gorm:"column:zhenduan;comment:诊断;type:varchar(1000);"`
}

func (Stubljh) TableName() string {
	return "stubljh"
}
