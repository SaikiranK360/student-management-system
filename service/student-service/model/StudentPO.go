package model

type StudentPO struct {
	Id   int64 `gorm:"primarykey"`
	Name string
	City string
}

func (StudentPO) TableName() string {
	return "student"
}
