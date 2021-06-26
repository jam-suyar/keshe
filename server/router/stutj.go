package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitStutjRouter(Router *gin.RouterGroup) {
	StutjRouter := Router.Group("stutj").Use(middleware.OperationRecord())
	{
		StutjRouter.POST("createStutj", v1.CreateStutj)             // 新建Stutj
		StutjRouter.DELETE("deleteStutj", v1.DeleteStutj)           // 删除Stutj
		StutjRouter.DELETE("deleteStutjByIds", v1.DeleteStutjByIds) // 批量删除Stutj
		StutjRouter.PUT("updateStutj", v1.UpdateStutj)              // 更新Stutj
		StutjRouter.GET("findStutj", v1.FindStutj)                  // 根据ID获取Stutj
		StutjRouter.GET("getStutjList", v1.GetStutjList)            // 获取Stutj列表
	}
}
