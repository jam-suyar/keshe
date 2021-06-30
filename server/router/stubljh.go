package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitStubljhRouter(Router *gin.RouterGroup) {
	StubljhRouter := Router.Group("stubljh").Use(middleware.OperationRecord())
	{
		StubljhRouter.POST("createStubljh", v1.CreateStubljh)             // 新建Stubljh
		StubljhRouter.DELETE("deleteStubljh", v1.DeleteStubljh)           // 删除Stubljh
		StubljhRouter.DELETE("deleteStubljhByIds", v1.DeleteStubljhByIds) // 批量删除Stubljh
		StubljhRouter.PUT("updateStubljh", v1.UpdateStubljh)              // 更新Stubljh
		StubljhRouter.GET("findStubljh", v1.FindStubljh)                  // 根据ID获取Stubljh
		StubljhRouter.GET("getStubljhList", v1.GetStubljhList)            // 获取Stubljh列表
	}
}
