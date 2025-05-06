package main

import (
	"log"
	"net"
	"os/user"

	studentproto "github.com/SaikiranK360/grpc-go-practice-2/proto-gen/student"
	"github.com/SaikiranK360/grpc-go-practice-2/service/student-service/db"
	"github.com/SaikiranK360/grpc-go-practice-2/service/student-service/handler"
	"github.com/SaikiranK360/grpc-go-practice-2/service/student-service/interceptor"
	"github.com/SaikiranK360/grpc-go-practice-2/shared/config"
	customlog "github.com/SaikiranK360/grpc-go-practice-2/shared/custom-log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var TCP = "tcp"

func main() {
	// Load the contents from the config.json file
	config.LoadConfig()

	// Set the log files
	customlog.CreateLogDirectories()
	currentUser, _ := user.Current()
	systemLogFile := customlog.GetLogFile(currentUser.HomeDir + "/studentprojectlogs/student-service.log")
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(systemLogFile, systemLogFile, systemLogFile))
	log.SetOutput(systemLogFile)

	requestLogFilePath := currentUser.HomeDir + "/requestlogs/student-service.log"
	interceptor.InitGRPCRequestLogger(requestLogFilePath)

	// Initiate the DB connection
	db.InitiateMySQLDBConnection()
	sqlDB, _ := db.MYSQLDB.DB()
	defer sqlDB.Close()

	lis, err := net.Listen(TCP, config.Config.Service.Student.Port)
	if err != nil {
		panic("Failed to listen at port " + lis.Addr().String())
	}
	log.Println("Listening at port ", config.Config.Service.Student.Port)

	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.LoggingInterceptor))
	studentproto.RegisterStudentServiceServer(server, &handler.StudentHandler{})

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %+v", err)
	}
}
