package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateStubl
//@description: 创建Stubl记录
//@param: stubl model.Stubl
//@return: err error

func CreateStubl(stubl model.Stubl) (err error) {
	err = global.GVA_DB.Create(&stubl).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStubl
//@description: 删除Stubl记录
//@param: stubl model.Stubl
//@return: err error

func DeleteStubl(stubl model.Stubl) (err error) {
	err = global.GVA_DB.Delete(&stubl).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStublByIds
//@description: 批量删除Stubl记录
//@param: ids request.IdsReq
//@return: err error

func DeleteStublByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Stubl{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateStubl
//@description: 更新Stubl记录
//@param: stubl *model.Stubl
//@return: err error

func UpdateStubl(stubl model.Stubl) (err error) {
	err = global.GVA_DB.Save(&stubl).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStubl
//@description: 根据id获取Stubl记录
//@param: id uint
//@return: err error, stubl model.Stubl

func GetStubl(id uint) (err error, stubl model.Stubl) {
	err = global.GVA_DB.Where("id = ?", id).First(&stubl).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStublInfoList
//@description: 分页获取Stubl记录
//@param: info request.StublSearch
//@return: err error, list interface{}, total int64

func GetStublInfoList(info request.StublSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Stubl{})
	var stubls []model.Stubl
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Xuehao != 0 {
		db = db.Where("`xuehao` = ?", info.Xuehao)
	}
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Xibie != "" {
		db = db.Where("`xibie` LIKE ?", "%"+info.Xibie+"%")
	}
	if info.Xingbie != "" {
		db = db.Where("`xingbie` LIKE ?", "%"+info.Xingbie+"%")
	}
	if info.Age != 0 {
		db = db.Where("`age` > ?", info.Age)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&stubls).Error
	return err, stubls, total
}
