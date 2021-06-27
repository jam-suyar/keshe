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

// @Tags Stutjts
// @Summary 创建Stutjts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutjts true "创建Stutjts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stutjts/createStutjts [post]
func CreateStutjts(c *gin.Context) {
	var stutjts model.Stutjts
	_ = c.ShouldBindJSON(&stutjts)
	if err := service.CreateStutjts(stutjts); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Stutjts
// @Summary 删除Stutjts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutjts true "删除Stutjts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stutjts/deleteStutjts [delete]
func DeleteStutjts(c *gin.Context) {
	var stutjts model.Stutjts
	_ = c.ShouldBindJSON(&stutjts)
	if err := service.DeleteStutjts(stutjts); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Stutjts
// @Summary 批量删除Stutjts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Stutjts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /stutjts/deleteStutjtsByIds [delete]
func DeleteStutjtsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteStutjtsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Stutjts
// @Summary 更新Stutjts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutjts true "更新Stutjts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stutjts/updateStutjts [put]
func UpdateStutjts(c *gin.Context) {
	var stutjts model.Stutjts
	_ = c.ShouldBindJSON(&stutjts)
	if err := service.UpdateStutjts(stutjts); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Stutjts
// @Summary 用id查询Stutjts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stutjts true "用id查询Stutjts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stutjts/findStutjts [get]
func FindStutjts(c *gin.Context) {
	var stutjts model.Stutjts
	_ = c.ShouldBindQuery(&stutjts)
	if err, restutjts := service.GetStutjts(stutjts.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restutjts": restutjts}, c)
	}
}

// @Tags Stutjts
// @Summary 分页获取Stutjts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StutjtsSearch true "分页获取Stutjts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stutjts/getStutjtsList [get]
func GetStutjtsList(c *gin.Context) {
	var pageInfo request.StutjtsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetStutjtsInfoList(pageInfo); err != nil {
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
