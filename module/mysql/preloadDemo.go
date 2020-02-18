package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
	"time"
)
type BaseModel struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Student struct {
	BaseModel
	Sex string `gorm:sex`
	Sname string `gorm:sname`
	Sage int `gorm:sage`
}

type Course struct {
	BaseModel
	StudentId int
	Student Student
	Cname string `gorm:cname`
}

func  main() {
	db2, err := gorm.Open("mysql", "root:123456@/demo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Info().Msg("数据库连接成功")
	if db2 == nil{
		return
	}
	db2.LogMode(true)
	//=========================//
	//queryStudents(db2)
	//queryCourses(db2)
	preloadDemo(db2)
	//=========================//
	db2.Close()
}

func queryStudents(db *gorm.DB) {
	var stu Student
	err := db.Find(&stu).Error
	if err != nil {
		log.Error().Msg(err.Error())
	}
	fmt.Printf("%+v\n",stu)
}
func queryCourses(db *gorm.DB) {
	var course Course
	err := db.Find(&course).Error
	if err != nil {
		log.Error().Msg(err.Error())
	}
	fmt.Printf("%+v\n",course)
}
// preload 查询
func preloadDemo(db *gorm.DB) {
	var course Course
	err := db.Preload("Student").
		Find(&course).Error
	if err != nil {
		log.Error().Msg(err.Error())
	}
	fmt.Printf("%+v\n",course)
}
