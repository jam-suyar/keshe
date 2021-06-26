package request

import "gin-vue-admin/model"

type StublSearch struct {
	model.Stubl
	PageInfo
}
