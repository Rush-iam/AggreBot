package grpc_client

import (
	"AggreBot/internal/pkg/api"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type Client struct {
	connection *grpc.ClientConn
	api        api.NewsfeedConfiguratorClient
	ctx        context.Context
}

func New(ctx context.Context, grpcServerEndpoint string) *Client {
	connection := connect(grpcServerEndpoint)
	return &Client{
		connection: connection,
		api:        api.NewNewsfeedConfiguratorClient(connection),
		ctx:        ctx,
	}
}

func (c *Client) Close() {
	_ = c.connection.Close()
	c.connection = nil
	c.api = nil
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
		log.Fatalf("Can't connect to gRPC server %v", err)
	} else {
		log.Printf("Connected to gRPC server: %s", grpcServerEndpoint)
	}
	return grpcConnection
}
