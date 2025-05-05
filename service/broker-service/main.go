package main

import (
	"log"
	"os/user"

	webapi "github.com/SaikiranK360/grpc-go-practice-2/service/broker-service/route/web-api"
	"github.com/SaikiranK360/grpc-go-practice-2/shared/config"
	customlog "github.com/SaikiranK360/grpc-go-practice-2/shared/custom-log"
	"github.com/gin-gonic/gin"
)

func main() {
	// Lod the configuration from the config.json file
	config.LoadConfig()

	gin.SetMode(gin.ReleaseMode)

	// Set the output log files
	customlog.CreateLogDirectories()
	currentUser, _ := user.Current()
	requestLogFile := customlog.GetLogFile(currentUser.HomeDir + "/requestlogs/broker-service.log")
	systemLogFile := customlog.GetLogFile(currentUser.HomeDir + "/studentprojectlogs/broker-service.log")

	gin.DefaultWriter = requestLogFile
	gin.DefaultErrorWriter = systemLogFile
	log.SetOutput(systemLogFile)

	router := gin.Default()

	router.SetTrustedProxies([]string{config.Config.Service.Broker.IP})

	webapi.RegisterWebAPIRoutes(&router.RouterGroup)

	log.Printf("Listening at port %+v", config.Config.Service.Broker.Port)
	err := router.Run(config.Config.Service.Broker.Port)

	if err != nil {
		log.Fatalf("%+v", err.Error())
	}
}
