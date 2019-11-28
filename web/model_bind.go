package xxx

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

/**
8 模型绑定与验证
*/
type User struct {
	Uname    string `form:"uname" json:"uname" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	/**
	日志写到文件里
	*/
	gin.DisableConsoleColor()
	file, _ := os.Create("practice1/prac.log")
	gin.DefaultWriter = io.MultiWriter(file)
	router := gin.Default()

	// 从json 解析
	router.POST("/login", func(context *gin.Context) {
		var user User
		// 验证账号密码
		if e := context.ShouldBind(&user); e == nil {
			if user.Uname == "feng" && user.Password == "123456" {
				// 登录成功
				context.JSON(http.StatusOK, gin.H{
					"message": "login successfully",
				})
			} else {
				// 登录失败
				context.JSON(http.StatusOK, gin.H{
					"message": "username or password err",
				})
			}
		} else {
			context.String(500, "server error")
		}
	})

	router.Run()
}
