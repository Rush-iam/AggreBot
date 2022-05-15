package main

import (
	"AggreBot/internal/backend"
	"AggreBot/internal/pkg/api"
	"AggreBot/internal/pkg/config"
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

var flags = map[string]string{
	"dbuser":   "Database Username",
	"dbpass":   "Database Password",
	"dbhost":   "Database Host",
	"dbname":   "Database Name",
	"grpchost": "gRPC Server Host",
	"resthost": "REST Proxy Host",
}

func main() {
	cfg := config.FromFlags(flags)

	db := db_client.NewClient(
		context.Background(), cfg["dbuser"], cfg["dbpass"], cfg["dbhost"], cfg["dbname"],
	)
	defer db.Close()

	go runGrpcServer(cfg["grpchost"], db)
	go runRestProxy(cfg["grpchost"], cfg["resthost"])

	<-exit_signal.Wait()
}

func runGrpcServer(grpcHost string, db *db_client.Client) {
	grpcServer := backend.NewServer(db)
	listener, err := net.Listen("tcp", grpcHost)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Start serving gRPC on %s...", grpcHost)
	err = grpcServer.Serve(listener)
	log.Fatal(err)
}

func runRestProxy(grpcHost string, restHost string) {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}
	err := api.RegisterNewsfeedConfiguratorHandlerFromEndpoint(
		ctx, mux, grpcHost, opts,
	)
	if err != nil {
		log.Fatalf("Can't connect to gRPC server: %v", err)
	}

	log.Printf("Start serving REST proxy on %s...", restHost)
	err = http.ListenAndServe(restHost, mux)
	log.Fatal(err)
}
