import service from '@/utils/request'

// @Tags Stutj
// @Summary 创建Stutj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutj true "创建Stutj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stutj/createStutj [post]
export const createStutj = (data) => {
  return service({
    url: '/stutj/createStutj',
    method: 'post',
    data
  })
}

// @Tags Stutj
// @Summary 删除Stutj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutj true "删除Stutj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stutj/deleteStutj [delete]
export const deleteStutj = (data) => {
  return service({
    url: '/stutj/deleteStutj',
    method: 'delete',
    data
  })
}

// @Tags Stutj
// @Summary 删除Stutj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Stutj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stutj/deleteStutj [delete]
export const deleteStutjByIds = (data) => {
  return service({
    url: '/stutj/deleteStutjByIds',
    method: 'delete',
    data
  })
}

// @Tags Stutj
// @Summary 更新Stutj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutj true "更新Stutj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stutj/updateStutj [put]
export const updateStutj = (data) => {
  return service({
    url: '/stutj/updateStutj',
    method: 'put',
    data
  })
}

// @Tags Stutj
// @Summary 用id查询Stutj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutj true "用id查询Stutj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stutj/findStutj [get]
export const findStutj = (params) => {
  return service({
    url: '/stutj/findStutj',
    method: 'get',
    params
  })
}

// @Tags Stutj
// @Summary 分页获取Stutj列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Stutj列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stutj/getStutjList [get]
export const getStutjList = (params) => {
  return service({
    url: '/stutj/getStutjList',
    method: 'get',
    params
  })
}