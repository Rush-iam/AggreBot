package main

import (
	"AggreBot/internal/backend"
	"AggreBot/internal/pkg/api"
	"AggreBot/internal/pkg/db_client"
	"AggreBot/internal/pkg/exit_signal"
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
	db := db_client.NewClient(
		context.Background(), dbUser, dbPassword, dbHost, dbPort, dbName,
	)
	defer db.Close()

	go runGrpcServer(db)
	go runRestProxy()

	<-exit_signal.Wait()
}

func runGrpcServer(db *db_client.Client) {
	grpcServer := backend.NewServer(db)
	listener, err := net.Listen("tcp", grpcServerEndpoint)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Start serving gRPC on %s...", grpcServerEndpoint)
	err = grpcServer.Serve(listener)
	log.Fatal(err)
}

func runRestProxy() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}
	err := api.RegisterNewsfeedConfiguratorHandlerFromEndpoint(
		ctx, mux, grpcServerEndpoint, opts,
	)
	if err != nil {
		log.Fatalf("Can't connect to gRPC server: %v", err)
	}

	log.Printf("Start serving REST proxy on %s...", restProxyEndpoint)
	err = http.ListenAndServe(restProxyEndpoint, mux)
	log.Fatal(err)
}
