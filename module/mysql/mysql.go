package main

import (
	"github.com/jinzhu/gorm"
	"log"
)

// store 存储 mysql 连接
type Store struct {
	Db *gorm.DB
}
var store Store

// 初始化
func Init() error {
	db, err := Connect()
	if err != nil {
		return err
	}
	store.Db = db
	return nil

}

// 连接数据库
func Connect() (*gorm.DB,error){
	// TODO 读取配置文件
	db, err := gorm.Open("mysql", "root:123456@/demo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("db open err: ",err)
		return nil,err
	}
	// TODO 日志
	log.Println("【数据库连接成功】")
	// TODO 连接池
	return db,nil
}

// 获得连接
func GetDb() *gorm.DB {
	return store.Db
}



