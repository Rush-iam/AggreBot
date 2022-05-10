package commands

import (
	"AggreBot/api"
	"AggreBot/internal/bot_ui/grpc_client"
	"context"
	"fmt"
)

func cmdDelete(c Command) *string {
	var reply string
	sourceToDelete, errReply := fetchSourceFromUserArg(c.userId, c.args)
	if errReply != nil {
		return errReply
	}
	_, err := grpc_client.Cl.DeleteSource(
		context.Background(),
		&api.SourceId{Id: sourceToDelete.Id},
	)
	if err != nil {
		reply = "⚠ Oops. Internal Error. Please try again later."
		return &reply
	}
	reply = fmt.Sprintf("🗑 %s", sourceToDelete.Name)
	return &reply
}
