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

// @Tags Stubljh
// @Summary 创建Stubljh
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubljh true "创建Stubljh"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stubljh/createStubljh [post]
func CreateStubljh(c *gin.Context) {
	var stubljh model.Stubljh
	_ = c.ShouldBindJSON(&stubljh)
	if err := service.CreateStubljh(stubljh); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Stubljh
// @Summary 删除Stubljh
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubljh true "删除Stubljh"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stubljh/deleteStubljh [delete]
func DeleteStubljh(c *gin.Context) {
	var stubljh model.Stubljh
	_ = c.ShouldBindJSON(&stubljh)
	if err := service.DeleteStubljh(stubljh); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Stubljh
// @Summary 批量删除Stubljh
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Stubljh"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /stubljh/deleteStubljhByIds [delete]
func DeleteStubljhByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteStubljhByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Stubljh
// @Summary 更新Stubljh
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubljh true "更新Stubljh"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stubljh/updateStubljh [put]
func UpdateStubljh(c *gin.Context) {
	var stubljh model.Stubljh
	_ = c.ShouldBindJSON(&stubljh)
	if err := service.UpdateStubljh(stubljh); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Stubljh
// @Summary 用id查询Stubljh
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubljh true "用id查询Stubljh"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stubljh/findStubljh [get]
func FindStubljh(c *gin.Context) {
	var stubljh model.Stubljh
	_ = c.ShouldBindQuery(&stubljh)
	if err, restubljh := service.GetStubljh(stubljh.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restubljh": restubljh}, c)
	}
}

// @Tags Stubljh
// @Summary 分页获取Stubljh列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StubljhSearch true "分页获取Stubljh列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stubljh/getStubljhList [get]
func GetStubljhList(c *gin.Context) {
	var pageInfo request.StubljhSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetStubljhInfoList(pageInfo); err != nil {
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
