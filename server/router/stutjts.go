package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitStutjtsRouter(Router *gin.RouterGroup) {
	StutjtsRouter := Router.Group("stutjts").Use(middleware.OperationRecord())
	{
		StutjtsRouter.POST("createStutjts", v1.CreateStutjts)             // 新建Stutjts
		StutjtsRouter.DELETE("deleteStutjts", v1.DeleteStutjts)           // 删除Stutjts
		StutjtsRouter.DELETE("deleteStutjtsByIds", v1.DeleteStutjtsByIds) // 批量删除Stutjts
		StutjtsRouter.PUT("updateStutjts", v1.UpdateStutjts)              // 更新Stutjts
		StutjtsRouter.GET("findStutjts", v1.FindStutjts)                  // 根据ID获取Stutjts
		StutjtsRouter.GET("getStutjtsList", v1.GetStutjtsList)            // 获取Stutjts列表
	}
}
