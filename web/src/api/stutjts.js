import service from '@/utils/request'

// @Tags Stutjts
// @Summary 创建Stutjts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutjts true "创建Stutjts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stutjts/createStutjts [post]
export const createStutjts = (data) => {
  return service({
    url: '/stutjts/createStutjts',
    method: 'post',
    data
  })
}

// @Tags Stutjts
// @Summary 删除Stutjts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutjts true "删除Stutjts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stutjts/deleteStutjts [delete]
export const deleteStutjts = (data) => {
  return service({
    url: '/stutjts/deleteStutjts',
    method: 'delete',
    data
  })
}

// @Tags Stutjts
// @Summary 删除Stutjts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Stutjts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stutjts/deleteStutjts [delete]
export const deleteStutjtsByIds = (data) => {
  return service({
    url: '/stutjts/deleteStutjtsByIds',
    method: 'delete',
    data
  })
}

// @Tags Stutjts
// @Summary 更新Stutjts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutjts true "更新Stutjts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stutjts/updateStutjts [put]
export const updateStutjts = (data) => {
  return service({
    url: '/stutjts/updateStutjts',
    method: 'put',
    data
  })
}

// @Tags Stutjts
// @Summary 用id查询Stutjts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutjts true "用id查询Stutjts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stutjts/findStutjts [get]
export const findStutjts = (params) => {
  return service({
    url: '/stutjts/findStutjts',
    method: 'get',
    params
  })
}

// @Tags Stutjts
// @Summary 分页获取Stutjts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Stutjts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stutjts/getStutjtsList [get]
export const getStutjtsList = (params) => {
  return service({
    url: '/stutjts/getStutjtsList',
    method: 'get',
    params
  })
}