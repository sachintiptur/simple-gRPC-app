package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/sachintiptur/grpc-app/pkg/frontend"
)

// frontend http server process
// server address can be passed as an argument
func main() {
	addr := flag.String("addr", ":8080", "Server address string")
	flag.Parse()

	data := &frontend.EnvData{}

	mux := http.NewServeMux()
	mux.HandleFunc("/getenv", data.HandleRequest)
	log.Println("Server is listening")
	log.Fatal(http.ListenAndServe(*addr, mux))
}
