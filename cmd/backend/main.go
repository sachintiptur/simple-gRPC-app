package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/sachintiptur/grpc-app/proto"
)

type server struct {
	pb.UnimplementedGetterServer
}

// GetEnvVariable gets the environment variable value
func (s *server) GetEnvVariable(ctx context.Context, req *pb.EnvRequest) (*pb.EnvResponse, error) {
	env := req.GetEnvName()
	if _, ok := os.LookupEnv(env); !ok {
		return &pb.EnvResponse{}, fmt.Errorf("env %s is not set", env)
	}
	val := os.Getenv(env)
	return &pb.EnvResponse{EnvValue: val}, nil
}

// backend gRPC server
// listens to gRPC requests and return environment variable's value
func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// register and start gRPC server
	srv := grpc.NewServer()
	pb.RegisterGetterServer(srv, &server{})
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
