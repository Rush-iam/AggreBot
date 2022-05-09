package commands

import (
	"AggreBot/api"
	"AggreBot/internal/tg_frontend/grpc_client"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func start(c Command) *string {
	var reply string
	_, err := grpc_client.Cl.AddUser(
		context.Background(),
		&api.UserId{Id: c.userId},
	)
	if err != nil && status.Code(err) != codes.AlreadyExists {
		reply = "⚠ Oops. Internal Error. Please try again later."
	} else {
		reply = "🤖"
	}
	return &reply
}
