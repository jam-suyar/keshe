package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateStubljh
//@description: 创建Stubljh记录
//@param: stubljh model.Stubljh
//@return: err error

func CreateStubljh(stubljh model.Stubljh) (err error) {
	err = global.GVA_DB.Create(&stubljh).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStubljh
//@description: 删除Stubljh记录
//@param: stubljh model.Stubljh
//@return: err error

func DeleteStubljh(stubljh model.Stubljh) (err error) {
	err = global.GVA_DB.Delete(&stubljh).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStubljhByIds
//@description: 批量删除Stubljh记录
//@param: ids request.IdsReq
//@return: err error

func DeleteStubljhByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Stubljh{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateStubljh
//@description: 更新Stubljh记录
//@param: stubljh *model.Stubljh
//@return: err error

func UpdateStubljh(stubljh model.Stubljh) (err error) {
	err = global.GVA_DB.Save(&stubljh).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStubljh
//@description: 根据id获取Stubljh记录
//@param: id uint
//@return: err error, stubljh model.Stubljh

func GetStubljh(id uint) (err error, stubljh model.Stubljh) {
	err = global.GVA_DB.Where("id = ?", id).First(&stubljh).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStubljhInfoList
//@description: 分页获取Stubljh记录
//@param: info request.StubljhSearch
//@return: err error, list interface{}, total int64

func GetStubljhInfoList(info request.StubljhSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Stubljh{})
	var stubljhs []model.Stubljh
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Zhenduan != "" {
		db = db.Where("`zhenduan` LIKE ?", "%"+info.Zhenduan+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&stubljhs).Error
	return err, stubljhs, total
}
