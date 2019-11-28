package models

import "github.com/jinzhu/gorm"

// 学生模型
type Student struct {
	gorm.Model
	sno   string
	sname string
}
