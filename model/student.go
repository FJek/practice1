package model

// 学生模型
type Student struct {
	BaseModel
	Sno   string `json:"sno"`
	Sname string `json:"sname"`
}
