package mapper

import (
	studentproto "github.com/SaikiranK360/grpc-go-practice-2/proto-gen/student"
	"github.com/SaikiranK360/grpc-go-practice-2/service/broker-service/api"
)

func ConvertStudentProtoToStudentResponse(s *studentproto.Student) *api.StudentResponseDTO {
	return &api.StudentResponseDTO{
		Id:   s.GetId(),
		Name: s.GetName(),
		City: s.GetCity(),
	}
}

func ConvertStudentRequestToStudentRequestProto(s *api.StudentCreationRequestDTO) *studentproto.StudentCreationRequest {
	return &studentproto.StudentCreationRequest{
		Name: s.Name,
		City: s.City,
	}
}
