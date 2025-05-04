package v1

import (
	"github.com/SaikiranK360/grpc-go-practice-2/service/broker-service/route/web-api/v1/student"
	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	student.RegisterStudentRoutesV1(v1)
}
