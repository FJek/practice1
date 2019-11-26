package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
)

func main() {
	/**
	7 写入日志文件
	*/
	// 禁用控制台颜色
	gin.DisableConsoleColor()
	// 创建写入日志的文件
	file, _ := os.Create("practice1/prac.log")
	// 默认写入到这个文件
	gin.DefaultWriter = io.MultiWriter(file)
	// 1 使用默认中间件创建一个路由：
	// 日志与恢复中间件
	r := gin.Default()
	//2 获取用户个人信息
	r.GET("/personal/:name", personal)
	r.POST("register", register)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 3 这个处理器可以匹配 /user/john ， 但是它不会匹配 /user
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	// 4 但是，这个可以匹配 /user/john 和 /user/john/send
	// 如果没有其他的路由匹配 /user/john ， 它将重定向到 /user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	// 5 查询字符串参数使用现有的底层 request 对象解析。
	// 请求响应匹配的 URL： /welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		// 这个是 c.Request.URL.Query().Get("lastname") 的快捷方式。
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	/**
	6 上传文件
	*/
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Print(file.Filename + ", file size:" + string(file.Size))
		// 上传到指定的dst
		//c.SaveUploadedFile(file,"/Users/Document/test.t")
		c.String(http.StatusOK, "%s uploaded ", file.Filename)
	})

	// 在 0.0.0.0:8080 上监听并服务
	r.Run()
}

func register(context *gin.Context) {
	if true {
		context.JSON(200, gin.H{
			"message": "注册成功",
		})
	} else {
		context.JSON(400, gin.H{
			"message": "注册失败",
		})
	}
}

func personal(context *gin.Context) {
	// 2 获取path中的name参数
	name := context.Param("name")
	context.JSON(200, gin.H{
		"name":  name,
		"email": "fengzhiwen@forchange.tech",
	})
}
