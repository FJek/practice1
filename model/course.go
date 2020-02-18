package model

type Course struct {
	//gorm.Model
	// 数据绑定
	Cno   string `form:"cno" json:"cno" binding:"required"`
	Cname string `form:"Cname" json:"Cname" binding:"required"`
	Tno   string `form:"Tno" json:"Tno" binding:"required"`
}

