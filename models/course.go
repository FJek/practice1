package models

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model

	cno   string
	cname string
	tno   string
}

// 添加课程
func AddCourse(course *Course) {

}
