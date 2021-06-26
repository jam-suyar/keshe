package request

import "gin-vue-admin/model"

type StutjSearch struct {
	model.Stutj
	PageInfo
}
