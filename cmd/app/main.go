package main

import (
	"AggreBot/api"
	"AggreBot/internal/app/db"
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
	dbUser     = "postgres"
	dbPassword = "j3qq4"
	dbHost     = "localhost"
	dbPort     = "5432"
	dbName     = "aggrebot"

	grpcServerEndpoint = "localhost:8080"
	restProxyEndpoint  = "localhost:8081"
)

func main() {
	db.Init(dbUser, dbPassword, dbHost, dbPort, dbName)
	defer db.Close()

	go runRestProxy()
	runGrpcServer()
}

func runGrpcServer() {
	lis, err := net.Listen("tcp", grpcServerEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterNewsfeedConfiguratorServer(grpcServer, handlers.Server{})

	log.Printf("Start serving gRPC on %s...", grpcServerEndpoint)
	err = grpcServer.Serve(lis)
	log.Fatal(err)
}

func runRestProxy() {
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
	err = http.ListenAndServe(restProxyEndpoint, mux)
	log.Fatal(err)
}
