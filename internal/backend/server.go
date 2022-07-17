package backend

import (
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/api"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/db_client"
	"google.golang.org/grpc"
)

type Server struct {
	api.UnimplementedNewsfeedConfiguratorServer
	db *db_client.Client
}

func NewServer(db *db_client.Client) *grpc.Server {
	server := grpc.NewServer()
	serverStruct := Server{
		db: db,
	}
	api.RegisterNewsfeedConfiguratorServer(server, serverStruct)
	return server
}
