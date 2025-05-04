package webapi

import (
	v1 "github.com/SaikiranK360/grpc-go-practice-2/service/broker-service/route/web-api/v1"
	"github.com/gin-gonic/gin"
)

func RegisterWebAPIRoutes(router *gin.RouterGroup) {
	webAPIRoutes := router.Group("/web-api")
	v1.RegisterV1Routes(webAPIRoutes)
}
