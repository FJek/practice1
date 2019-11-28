package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"practice1/web/router"
)

func main() {
	engine := gin.Default()
	// 初始化路由
	router.InitRouter(engine)
	// 连接数据库
	db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("db conn err :", err)
	}
	defer db.Close()

	engine.Run()

}
