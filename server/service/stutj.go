package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateStutj
//@description: 创建Stutj记录
//@param: stutj model.Stutj
//@return: err error

func CreateStutj(stutj model.Stutj) (err error) {
	err = global.GVA_DB.Create(&stutj).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStutj
//@description: 删除Stutj记录
//@param: stutj model.Stutj
//@return: err error

func DeleteStutj(stutj model.Stutj) (err error) {
	err = global.GVA_DB.Delete(&stutj).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStutjByIds
//@description: 批量删除Stutj记录
//@param: ids request.IdsReq
//@return: err error

func DeleteStutjByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Stutj{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateStutj
//@description: 更新Stutj记录
//@param: stutj *model.Stutj
//@return: err error

func UpdateStutj(stutj model.Stutj) (err error) {
	err = global.GVA_DB.Save(&stutj).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStutj
//@description: 根据id获取Stutj记录
//@param: id uint
//@return: err error, stutj model.Stutj

func GetStutj(id uint) (err error, stutj model.Stutj) {
	err = global.GVA_DB.Where("id = ?", id).First(&stutj).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStutjInfoList
//@description: 分页获取Stutj记录
//@param: info request.StutjSearch
//@return: err error, list interface{}, total int64

func GetStutjInfoList(info request.StutjSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Stutj{})
	var stutjs []model.Stutj
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Xibie != "" {
		db = db.Where("`xibie` LIKE ?", "%"+info.Xibie+"%")
	}
	if info.Xuehao != 0 {
		db = db.Where("`xuehao` = ?", info.Xuehao)
	}
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Xingbie != "" {
		db = db.Where("`xingbie` LIKE ?", "%"+info.Xingbie+"%")
	}
	if info.Age != 0 {
		db = db.Where("`ahe` > ?", info.Age)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&stutjs).Error
	return err, stutjs, total
}
