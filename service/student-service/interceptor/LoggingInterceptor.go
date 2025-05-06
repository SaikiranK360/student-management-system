package interceptor

import (
	"context"
	"log"

	customlog "github.com/SaikiranK360/grpc-go-practice-2/shared/custom-log"
	"google.golang.org/grpc"
)

var GRPCRequestLogger *log.Logger

/*
I am not putting this logic in shared since grpc dependency doesn't have to be there in shared.
*/
func InitGRPCRequestLogger(pathName string) {
	logFile := customlog.GetLogFile(pathName)
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
