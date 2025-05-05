package repository

import (
	"log"

	"github.com/SaikiranK360/grpc-go-practice-2/service/student-service/db"
	"github.com/SaikiranK360/grpc-go-practice-2/service/student-service/model"
)

func GetStudentById(id int64) *model.StudentPO {
	var student model.StudentPO
	db.MYSQLDB.First(&student, id)
	return &student
}

func CreateStudent(studentPO *model.StudentPO) *model.StudentPO {
	result := db.MYSQLDB.Create(studentPO)
	if result.Error != nil {
		log.Fatalf("%+v", result.Error)
	}
	return studentPO
}
