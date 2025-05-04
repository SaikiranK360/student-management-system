package main

import (
	webapi "github.com/SaikiranK360/grpc-go-practice-2/service/broker-service/route/web-api"
	"github.com/SaikiranK360/grpc-go-practice-2/shared/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.SetTrustedProxies([]string{config.Config.Service.Broker.IP})

	webapi.RegisterWebAPIRoutes(&router.RouterGroup)

	router.Run(config.Config.Service.Broker.Port)
}
