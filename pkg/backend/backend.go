package backend

import (
	"context"
	"fmt"
	"os"

	pb "github.com/sachintiptur/grpc-app/proto"
)

// Server struct to implement the GetterServer interface
type Server struct {
	pb.UnimplementedGetterServer
}

// GetEnvVariable gets the environment variable value
func (s *Server) GetEnvVariable(ctx context.Context, req *pb.EnvRequest) (*pb.EnvResponse, error) {
	env := req.GetEnvName()
	if _, ok := os.LookupEnv(env); !ok {
		return &pb.EnvResponse{}, fmt.Errorf("env %s is not set", env)
	}
	val := os.Getenv(env)
	return &pb.EnvResponse{EnvValue: val}, nil
}
