package main

import (
	"log"
	"net"
	"os"

	"github.com/sachintiptur/grpc-app/pkg/backend"
	"google.golang.org/grpc"

	pb "github.com/sachintiptur/grpc-app/proto"
)

// Backend gRPC server listens to gRPC requests and return environment variable's value
func main() {
	grpcServerAddress := os.Getenv("GRPC_SERVER")
	if grpcServerAddress == "" {
		grpcServerAddress = ":9000"
	}
	listener, err := net.Listen("tcp", grpcServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// register and start gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterGetterServer(grpcServer, &backend.Server{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
