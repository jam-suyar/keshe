package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateStutjts
//@description: 创建Stutjts记录
//@param: stutjts model.Stutjts
//@return: err error

func CreateStutjts(stutjts model.Stutjts) (err error) {
	err = global.GVA_DB.Create(&stutjts).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStutjts
//@description: 删除Stutjts记录
//@param: stutjts model.Stutjts
//@return: err error

func DeleteStutjts(stutjts model.Stutjts) (err error) {
	err = global.GVA_DB.Delete(&stutjts).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStutjtsByIds
//@description: 批量删除Stutjts记录
//@param: ids request.IdsReq
//@return: err error

func DeleteStutjtsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Stutjts{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateStutjts
//@description: 更新Stutjts记录
//@param: stutjts *model.Stutjts
//@return: err error

func UpdateStutjts(stutjts model.Stutjts) (err error) {
	err = global.GVA_DB.Save(&stutjts).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStutjts
//@description: 根据id获取Stutjts记录
//@param: id uint
//@return: err error, stutjts model.Stutjts

func GetStutjts(id uint) (err error, stutjts model.Stutjts) {
	err = global.GVA_DB.Where("id = ?", id).First(&stutjts).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStutjtsInfoList
//@description: 分页获取Stutjts记录
//@param: info request.StutjtsSearch
//@return: err error, list interface{}, total int64

func GetStutjtsInfoList(info request.StutjtsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Stutjts{})
	var stutjtss []model.Stutjts
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Height != 0 {
		db = db.Where("`height` = ?", info.Height)
	}
	if info.Weight != 0 {
		db = db.Where("`weight` > ?", info.Weight)
	}
	if info.Xunwei != 0 {
		db = db.Where("`xunwei` > ?", info.Xunwei)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&stutjtss).Error
	return err, stutjtss, total
}
