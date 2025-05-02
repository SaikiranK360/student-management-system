package student

import (
	"github.com/SaikiranK360/grpc-go-practice-2/service/broker-service/controller/v1/student"
	"github.com/gin-gonic/gin"
)

func RegisterStudentRoutesV1(rg *gin.RouterGroup) {
	rg.GET("/:id", student.GetStudent)
}
