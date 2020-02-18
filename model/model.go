package model

// 所有模型都必须继承此模型
type BaseModel struct {
	Id int			 `gorm:"primary_key" json:"id"`
	Create_at string `gorm:"default:current_time" json:"create_at"`
	Update_at string `gorm:"default:current_time" json:"update_at"`
}

