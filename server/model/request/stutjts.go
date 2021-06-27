package request

import "gin-vue-admin/model"

type StutjtsSearch struct {
	model.Stutjts
	PageInfo
}
