package repository

import (
	"github.com/SaikiranK360/grpc-go-practice-2/service/student-service/db"
	"github.com/SaikiranK360/grpc-go-practice-2/service/student-service/model"
)

func GetStudentById(id int64) *model.StudentPO {
	var student model.StudentPO
	db.MYSQLDB.First(&student, id)
	return &student
}
