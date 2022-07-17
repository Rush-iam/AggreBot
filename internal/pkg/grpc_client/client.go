package grpc_client

import (
	"context"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client struct {
	connection *grpc.ClientConn
	api        api.NewsfeedConfiguratorClient
	ctx        context.Context
}

func New(connCtx context.Context, grpcServerEndpoint string) *Client {
	connection := connect(connCtx, grpcServerEndpoint)
	return &Client{
		connection: connection,
		api:        api.NewNewsfeedConfiguratorClient(connection),
		ctx:        context.Background(),
	}
}

func (c *Client) Close() {
	_ = c.connection.Close()
	c.connection = nil
	c.api = nil
}

func connect(ctx context.Context, grpcServerEndpoint string) *grpc.ClientConn {
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
