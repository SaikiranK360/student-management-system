package handler

import (
	"context"

	student "github.com/SaikiranK360/grpc-go-practice-2/proto-gen/student"
	"github.com/SaikiranK360/grpc-go-practice-2/service/student-service/mapper"
	"github.com/SaikiranK360/grpc-go-practice-2/service/student-service/repository"
)

type StudentHandler struct {
	student.UnimplementedStudentServiceServer
}

func (studentHandler *StudentHandler) GetStudent(ctx context.Context, request *student.StudentRequest) (*student.Student, error) {
	student := repository.GetStudentById(request.GetId())
	return mapper.ConvertStudentPOToStudentProto(student), nil
}
