package main

import (
	"log"
	"net"

	studentproto "github.com/SaikiranK360/grpc-go-practice-2/proto-gen/student"
	"github.com/SaikiranK360/grpc-go-practice-2/service/student-service/db"
	"github.com/SaikiranK360/grpc-go-practice-2/service/student-service/handler"
	"github.com/SaikiranK360/grpc-go-practice-2/shared/config"
	"google.golang.org/grpc"
)

var TCP = "tcp"

func main() {
	// Load the contents from the config.json file
	config.LoadConfig()

	// Initiate the DB connection
	db.InitiateMySQLDBConnection()
	sqlDB, _ := db.MYSQLDB.DB()
	defer sqlDB.Close()

	lis, err := net.Listen(TCP, config.Config.Service.Student.Port)
	if err != nil {
		panic("Failed to listen at port " + lis.Addr().String())
	}
	log.Println("Listening at port ", config.Config.Service.Student.Port)

	server := grpc.NewServer()
	studentproto.RegisterStudentServiceServer(server, &handler.StudentHandler{})

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %+v", err)
	}
}
