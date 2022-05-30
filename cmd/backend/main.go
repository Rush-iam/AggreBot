package main

import (
	"AggreBot/internal/backend"
	"AggreBot/internal/pkg/api"
	"AggreBot/internal/pkg/config"
	"AggreBot/internal/pkg/db_client"
	"AggreBot/internal/pkg/exit_signal"
	"context"
	"errors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"time"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	db, err := db_client.NewClient(
		ctx, cfg["dbuser"], cfg["dbpass"], cfg["dbhost"], cfg["dbname"],
	)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	grpcServer := backend.NewServer(db)
	go runGrpcServer(cfg["grpchost"], grpcServer)

	restProxy := &http.Server{Addr: cfg["resthost"]}
	go runRestProxy(cfg["grpchost"], restProxy)

	<-exit_signal.Wait()
	shutdownServers(grpcServer, restProxy)
}

func runGrpcServer(grpcHost string, grpcServer *grpc.Server) {
	listener, err := net.Listen("tcp", grpcHost)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Start serving gRPC on %s...", grpcHost)
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func runRestProxy(grpcHost string, restProxy *http.Server) {
	mux := runtime.NewServeMux()
	restProxy.Handler = mux
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}
	err := api.RegisterNewsfeedConfiguratorHandlerFromEndpoint(
		context.Background(), mux, grpcHost, opts,
	)
	if err != nil {
		log.Fatalf("Can't connect to gRPC server: %v", err)
	}

	log.Printf("Start serving REST proxy on %s...", restProxy.Addr)
	err = restProxy.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func shutdownServers(grpcServer *grpc.Server, restProxy *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := restProxy.Shutdown(ctx); err != nil {
		log.Printf("REST proxy error while shutting down: %v", err)
	}
	log.Printf("REST proxy shutted down")
	grpcServer.GracefulStop()
	log.Printf("gRPC server shutted down")
}
