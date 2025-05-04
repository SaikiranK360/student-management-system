package model

import "github.com/jinzhu/gorm"

type StudentPO struct {
	gorm.Model
	Id   int64
	Name string
	City string
}

func (StudentPO) TableName() string {
	return "student"
}
