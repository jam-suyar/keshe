package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Stutj
// @Summary 创建Stutj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutj true "创建Stutj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stutj/createStutj [post]
func CreateStutj(c *gin.Context) {
	var stutj model.Stutj
	_ = c.ShouldBindJSON(&stutj)
	if err := service.CreateStutj(stutj); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Stutj
// @Summary 删除Stutj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutj true "删除Stutj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stutj/deleteStutj [delete]
func DeleteStutj(c *gin.Context) {
	var stutj model.Stutj
	_ = c.ShouldBindJSON(&stutj)
	if err := service.DeleteStutj(stutj); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Stutj
// @Summary 批量删除Stutj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Stutj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /stutj/deleteStutjByIds [delete]
func DeleteStutjByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteStutjByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Stutj
// @Summary 更新Stutj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutj true "更新Stutj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stutj/updateStutj [put]
func UpdateStutj(c *gin.Context) {
	var stutj model.Stutj
	_ = c.ShouldBindJSON(&stutj)
	if err := service.UpdateStutj(stutj); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Stutj
// @Summary 用id查询Stutj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutj true "用id查询Stutj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stutj/findStutj [get]
func FindStutj(c *gin.Context) {
	var stutj model.Stutj
	_ = c.ShouldBindQuery(&stutj)
	if err, restutj := service.GetStutj(stutj.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restutj": restutj}, c)
	}
}

// @Tags Stutj
// @Summary 分页获取Stutj列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StutjSearch true "分页获取Stutj列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stutj/getStutjList [get]
func GetStutjList(c *gin.Context) {
	var pageInfo request.StutjSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetStutjInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
