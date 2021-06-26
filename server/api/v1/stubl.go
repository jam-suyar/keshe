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

// @Tags Stubl
// @Summary 创建Stubl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubl true "创建Stubl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stubl/createStubl [post]
func CreateStubl(c *gin.Context) {
	var stubl model.Stubl
	_ = c.ShouldBindJSON(&stubl)
	if err := service.CreateStubl(stubl); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Stubl
// @Summary 删除Stubl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubl true "删除Stubl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stubl/deleteStubl [delete]
func DeleteStubl(c *gin.Context) {
	var stubl model.Stubl
	_ = c.ShouldBindJSON(&stubl)
	if err := service.DeleteStubl(stubl); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Stubl
// @Summary 批量删除Stubl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Stubl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /stubl/deleteStublByIds [delete]
func DeleteStublByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteStublByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Stubl
// @Summary 更新Stubl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubl true "更新Stubl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stubl/updateStubl [put]
func UpdateStubl(c *gin.Context) {
	var stubl model.Stubl
	_ = c.ShouldBindJSON(&stubl)
	if err := service.UpdateStubl(stubl); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Stubl
// @Summary 用id查询Stubl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stubl true "用id查询Stubl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stubl/findStubl [get]
func FindStubl(c *gin.Context) {
	var stubl model.Stubl
	_ = c.ShouldBindQuery(&stubl)
	if err, restubl := service.GetStubl(stubl.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restubl": restubl}, c)
	}
}

// @Tags Stubl
// @Summary 分页获取Stubl列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StublSearch true "分页获取Stubl列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stubl/getStublList [get]
func GetStublList(c *gin.Context) {
	var pageInfo request.StublSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetStublInfoList(pageInfo); err != nil {
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
