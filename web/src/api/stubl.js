import service from '@/utils/request'

// @Tags Stubl
// @Summary 创建Stubl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubl true "创建Stubl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stubl/createStubl [post]
export const createStubl = (data) => {
  return service({
    url: '/stubl/createStubl',
    method: 'post',
    data
  })
}

// @Tags Stubl
// @Summary 删除Stubl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubl true "删除Stubl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stubl/deleteStubl [delete]
export const deleteStubl = (data) => {
  return service({
    url: '/stubl/deleteStubl',
    method: 'delete',
    data
  })
}

// @Tags Stubl
// @Summary 删除Stubl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Stubl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stubl/deleteStubl [delete]
export const deleteStublByIds = (data) => {
  return service({
    url: '/stubl/deleteStublByIds',
    method: 'delete',
    data
  })
}

// @Tags Stubl
// @Summary 更新Stubl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubl true "更新Stubl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stubl/updateStubl [put]
export const updateStubl = (data) => {
  return service({
    url: '/stubl/updateStubl',
    method: 'put',
    data
  })
}

// @Tags Stubl
// @Summary 用id查询Stubl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubl true "用id查询Stubl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stubl/findStubl [get]
export const findStubl = (params) => {
  return service({
    url: '/stubl/findStubl',
    method: 'get',
    params
  })
}

// @Tags Stubl
// @Summary 分页获取Stubl列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Stubl列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stubl/getStublList [get]
export const getStublList = (params) => {
  return service({
    url: '/stubl/getStublList',
    method: 'get',
    params
  })
}