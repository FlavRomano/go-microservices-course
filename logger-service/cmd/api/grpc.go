package main

import (
	"context"
	"fmt"
	"log"
	"logger-service/cmd/data"
	"logger-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer // required for every service on grpc, it's to ensure backwards compatibility
	Models                             data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{
			Result: "failed to log",
		}
		return res, err
	}

	res := &logs.LogResponse{
		Result: "logged",
	}
	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC due %v", err)
	}

	svr := grpc.NewServer()

	logs.RegisterLogServiceServer(svr, &LogServer{Models: app.Models})
	log.Println("gRPC server started on port", gRpcPort)

	if err := svr.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC due %v", err)
	}
}
