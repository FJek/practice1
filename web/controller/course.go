package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"practice1/model"
)


// 添加课程
func AddCourse(ctx *gin.Context) {
	var Course model.Course
	if err := ctx.ShouldBindJSON(&Course); err == nil {

		err := Course.AddCourse()
		// TODO 请求异常处理
		if err != nil {
			log.Print("course add err: ",err)
		}
	}
	// TODO 响应处理
}
