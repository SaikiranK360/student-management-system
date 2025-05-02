package webapi

import (
	"github.com/SaikiranK360/grpc-go-practice-2/service/broker-service/route/web-api/v1/student"
	"github.com/gin-gonic/gin"
)

func RegisterWebAPIRoutes(rg *gin.RouterGroup) {
	studentRoutes := rg.Group("/student")
	student.RegisterStudentRoutesV1(studentRoutes)
}
