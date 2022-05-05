package main

import (
	"AggreBot/api"
	"AggreBot/internal/app/handlers"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

const (
	grpcServerEndpoint = "localhost:8080"
	restProxyEndpoint  = "localhost:8081"
)

func main() {
	lis, err := net.Listen("tcp", grpcServerEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterNewsfeedConfiguratorServer(grpcServer, handlers.Server{})

	go restProxyRun()
	log.Printf("Start serving gRPC on %s...", grpcServerEndpoint)
	log.Fatal(grpcServer.Serve(lis))
}

func restProxyRun() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := api.RegisterNewsfeedConfiguratorHandlerFromEndpoint(
		ctx, mux, grpcServerEndpoint, opts,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start serving REST proxy on %s...", restProxyEndpoint)
	log.Fatal(http.ListenAndServe(restProxyEndpoint, mux))
}
