package student

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	studentproto "github.com/SaikiranK360/grpc-go-practice-2/proto-gen/student"
	"github.com/SaikiranK360/grpc-go-practice-2/service/broker-service/api"
	"github.com/SaikiranK360/grpc-go-practice-2/service/broker-service/mapper"
	"github.com/SaikiranK360/grpc-go-practice-2/shared/config"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetStudent(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	studentConn, err := grpc.NewClient(config.Config.Service.Student.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}
	defer studentConn.Close()

	studentServiceClient := studentproto.NewStudentServiceClient(studentConn)
	studentCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	student, err := studentServiceClient.GetStudent(studentCtx, &studentproto.StudentRequest{Id: id})
	if err != nil {
		log.Fatalf("Error while fetching the student details: %+v", err)
	}

	studentResponseDTO := mapper.ConvertStudentProtoToStudentResponse(student)

	ctx.JSON(http.StatusOK, studentResponseDTO)
}

func CreateStudent(ctx *gin.Context) {
	var request api.StudentCreationRequestDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	studentConn, err := grpc.NewClient(config.Config.Service.Student.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}
	defer studentConn.Close()

	studentServiceClient := studentproto.NewStudentServiceClient(studentConn)
	studentCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	studentServiceClient.CreateStudent(studentCtx, mapper.ConvertStudentRequestToStudentRequestProto(&request))
}
