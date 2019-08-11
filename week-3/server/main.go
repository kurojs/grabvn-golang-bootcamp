package main

import (
	"log"

	"github.com/kurojs/grabvn-golang-bootcamp/week-3/server/grpc"
	"github.com/kurojs/grabvn-golang-bootcamp/week-3/server/http"
)

func main() {
	go func() {
		log.Println("starting  gRPC server...")
		grpc.StartGRPCServer()
	}()

	log.Println("starting  HTTP server...")
	http.StartHTTPServer()
}
