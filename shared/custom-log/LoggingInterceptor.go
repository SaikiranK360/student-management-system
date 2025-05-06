package customlog

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

var GRPCRequestLogger *log.Logger

func InitGRPCRequestLogger(pathName string) {
	logFile := GetLogFile(pathName)
	GRPCRequestLogger = log.New(logFile, "", log.LstdFlags|log.Lshortfile)
}

func LoggingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	GRPCRequestLogger.Printf("Incoming: %s", info.FullMethod)

	resp, err := handler(ctx, req)
	if err != nil {
		GRPCRequestLogger.Printf("Error: %v", err)
	} else {
		GRPCRequestLogger.Printf("Success: %s", info.FullMethod)
	}

	return resp, err
}
