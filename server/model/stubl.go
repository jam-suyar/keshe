package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Stubl struct {
	global.GVA_MODEL
	Xeuhao  int    `json:"xuehao" form:"xuehao" gorm:"column:xuehao;comment:学号;type:int;size:1000;"`
	Xibie   string `json:"xiebi" form:"xiebi" gorm:"column:xiebi;comment:系别;type:varchar(10);"`
	Name    string `json:"name" form:"name" gorm:"column:name;comment:姓名;type:varchar(100);"`
	Xingbie string `json:"xingbie" form:"xingbie" gorm:"column:xingbie;comment:性别;type:varchar(100);"`
	Age     int    `json:"age" form:"age" gorm:"column:age;comment:年龄;type:int;size:100;"`
}

func (Stubl) TableName() string {
	return "stubl"
}
