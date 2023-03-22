package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

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
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

// handleRequest handles http request and renders html page with env data
func (data *EnvData) handleRequest(writer http.ResponseWriter, request *http.Request) {
	if err := data.getEnvData(); err != nil {
		log.Println(err.Error())
		http.Error(writer, fmt.Sprintf("Env variable %s is not set", data.EnvName), http.StatusBadRequest)
		return
	}
	parsedTemplate, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println("error parsing html file")
		log.Println(err.Error())
		return
	}
	err = parsedTemplate.Execute(writer, data)
	if err != nil {
		log.Println("error executing template :", err)
		return
	}
}

// frontend http server process
// server address and environment variable name can be passed as arguments
func main() {
	addr := flag.String("addr", ":8080", "Server address string")
	env := flag.String("env", "", "Environment variable name")
	flag.Parse()

	data := &EnvData{EnvName: *env}

	mux := http.NewServeMux()
	mux.HandleFunc("/getenv", data.handleRequest)
	log.Println("Server is listening")
	log.Fatal(http.ListenAndServe(*addr, mux))

}
