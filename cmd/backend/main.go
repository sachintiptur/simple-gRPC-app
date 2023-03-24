package main

import (
	"log"
	"net"

	"github.com/sachintiptur/grpc-app/pkg/backend"
	"google.golang.org/grpc"

	pb "github.com/sachintiptur/grpc-app/proto"
)

// Backend gRPC server listens to gRPC requests and return environment variable's value
func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// register and start gRPC server
	srv := grpc.NewServer()
	pb.RegisterGetterServer(srv, &backend.Server{})
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
