import service from '@/utils/request'

// @Tags Stubljh
// @Summary 创建Stubljh
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubljh true "创建Stubljh"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stubljh/createStubljh [post]
export const createStubljh = (data) => {
  return service({
    url: '/stubljh/createStubljh',
    method: 'post',
    data
  })
}

// @Tags Stubljh
// @Summary 删除Stubljh
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubljh true "删除Stubljh"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stubljh/deleteStubljh [delete]
export const deleteStubljh = (data) => {
  return service({
    url: '/stubljh/deleteStubljh',
    method: 'delete',
    data
  })
}

// @Tags Stubljh
// @Summary 删除Stubljh
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Stubljh"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stubljh/deleteStubljh [delete]
export const deleteStubljhByIds = (data) => {
  return service({
    url: '/stubljh/deleteStubljhByIds',
    method: 'delete',
    data
  })
}

// @Tags Stubljh
// @Summary 更新Stubljh
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubljh true "更新Stubljh"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stubljh/updateStubljh [put]
export const updateStubljh = (data) => {
  return service({
    url: '/stubljh/updateStubljh',
    method: 'put',
    data
  })
}

// @Tags Stubljh
// @Summary 用id查询Stubljh
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubljh true "用id查询Stubljh"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stubljh/findStubljh [get]
export const findStubljh = (params) => {
  return service({
    url: '/stubljh/findStubljh',
    method: 'get',
    params
  })
}

// @Tags Stubljh
// @Summary 分页获取Stubljh列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Stubljh列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stubljh/getStubljhList [get]
export const getStubljhList = (params) => {
  return service({
    url: '/stubljh/getStubljhList',
    method: 'get',
    params
  })
}