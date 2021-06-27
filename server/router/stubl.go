package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitStublRouter(Router *gin.RouterGroup) {
	StublRouter := Router.Group("stubl").Use(middleware.OperationRecord())
	{
		StublRouter.POST("createStubl", v1.CreateStubl)             // 新建Stubl
		StublRouter.DELETE("deleteStubl", v1.DeleteStubl)           // 删除Stubl
		StublRouter.DELETE("deleteStublByIds", v1.DeleteStublByIds) // 批量删除Stubl
		StublRouter.PUT("updateStubl", v1.UpdateStubl)              // 更新Stubl
		StublRouter.GET("findStubl", v1.FindStubl)                  // 根据ID获取Stubl
		StublRouter.GET("getStublList", v1.GetStublList)            // 获取Stubl列表
	}
}
