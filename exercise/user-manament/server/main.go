package main

import (
	"log"

	"github.com/kurojs/grabvn-golang-bootcamp/exercise/user-manament/server/grpc"
	"github.com/kurojs/grabvn-golang-bootcamp/exercise/user-manament/server/http"
)

func main() {
	go func() {
		log.Println("starting  gRPC server...")
		grpc.StartGRPCServer()
	}()

	log.Println("starting  HTTP server...")
	http.StartHTTPServer()
}
