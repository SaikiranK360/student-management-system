package mapper

import (
	"github.com/SaikiranK360/grpc-go-practice-2/proto-gen/student"
	"github.com/SaikiranK360/grpc-go-practice-2/service/student-service/model"
)

func ConvertStudentPOToStudentProto(studentPO *model.StudentPO) *student.Student {
	if studentPO == nil {
		return nil
	}

	return &student.Student{
		Id:   studentPO.Id,
		Name: studentPO.Name,
		City: studentPO.City,
	}
}
