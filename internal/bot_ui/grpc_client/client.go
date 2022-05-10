package grpc_client

import (
	"AggreBot/api"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var connection *grpc.ClientConn
var Cl api.NewsfeedConfiguratorClient

func Init(grpcServerEndpoint string) {
	connection = connect(grpcServerEndpoint)
	Cl = api.NewNewsfeedConfiguratorClient(connection)
}

func Close() {
	_ = connection.Close()
	connection = nil
	Cl = nil
}

func connect(grpcServerEndpoint string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	grpcConnection, err := grpc.DialContext(
		ctx,
		grpcServerEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Can't connect to gRPC server: %v", err)
	} else {
		log.Printf("Connected to gRPC server: %s", grpcServerEndpoint)
	}
	return grpcConnection
}
