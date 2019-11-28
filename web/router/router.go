package router

import (
	"github.com/gin-gonic/gin"
	"practice1/web/controller"
)

func InitRouter(router *gin.Engine) {
	// 测试
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	/**
	课程
	*/
	// 添加课程 course/add
	router.POST("/course/add", controller.AddCourse)
}
