package main

import "github.com/gin-gonic/gin"

func main() {
	// 使用默认中间件创建一个路由：
	// 日志与恢复中间件
	r := gin.Default()
	//1 获取用户个人信息
	r.GET("/personal/:name",personal)
	r.POST("register",register)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 在 0.0.0.0:8080 上监听并服务
}

func register(context *gin.Context) {
	if true {
		context.JSON(200,gin.H{
			"message": "注册成功",
		})
	} else {
		context.JSON(400,gin.H{
			"message": "注册失败",
		})
	}
}

func personal(context *gin.Context) {
	// 2 获取path中的name参数
	name := context.Param("name")
	context.JSON(200,gin.H{
		"name": name,
		"email": "fengzhiwen@forchange.tech",
	})
}