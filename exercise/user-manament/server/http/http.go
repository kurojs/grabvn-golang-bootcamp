package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/kurojs/grabvn-golang-bootcamp/exercise/user-manament/proto"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	grpcPort = ":50051"
	httpPort = ":8080"
)

func StartHTTPServer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := proto.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway %v", err)
	}

	srv := http.Server{
		Addr:    httpPort,
		Handler: mux,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down HTTP server...")
			if err := srv.Shutdown(context.Background()); err != nil {
				log.Fatalf("failed to shutdown HTTP server: %v", err)
			}
		}
	}()
	srv.ListenAndServe()
}
