package frontend

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	pb "github.com/sachintiptur/grpc-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// EnvData struct to store environment variable name and value
type EnvData struct {
	EnvName  string
	EnvValue string
}

// getEnvData makes gRPC call to get environment variable information
func (data *EnvData) getEnvData() (err error) {
	grpcServer := os.Getenv("GRPC_SERVER")
	if grpcServer == "" {
		grpcServer = ":9000"
	}
	conn, err := grpc.Dial(grpcServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewGetterClient(conn)
	res, err := client.GetEnvVariable(context.Background(), &pb.EnvRequest{EnvName: data.EnvName})
	if err != nil {
		return err
	}
	data.EnvValue = res.GetEnvValue()

	return nil
}

// HandleRequest handles http request and renders html page with env data
func (data *EnvData) HandleRequest(writer http.ResponseWriter, request *http.Request) {
	// read env variable name from http query
	data.EnvName = request.URL.Query().Get("env")
	if data.EnvName == "" {
		http.Error(writer, "env query parameter is missing in the request", http.StatusBadRequest)
		return
	}
	if err := data.getEnvData(); err != nil {
		log.Println(err.Error())
		http.Error(writer, fmt.Sprintf("Env variable %s is not set", data.EnvName), http.StatusBadRequest)
		return
	}
	parsedTemplate, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = parsedTemplate.Execute(writer, data)
	if err != nil {
		log.Println("error executing template :", err)
		return
	}
}
