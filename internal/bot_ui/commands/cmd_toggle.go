package commands

import (
	"AggreBot/api"
	"AggreBot/internal/bot_ui/grpc_client"
	"context"
	"fmt"
)

func cmdToggle(c Command) *string {
	var reply string

	sourceToToggle, errReply := fetchSourceFromUserArg(c.userId, c.args)
	if errReply != nil {
		return errReply
	}

	source, err := grpc_client.Cl.UpdateSourceToggleActive(
		context.Background(),
		&api.SourceId{Id: sourceToToggle.Id},
	)
	if err != nil {
		reply = "⚠ Oops. Internal Error. Please try again later."
		return &reply
	}
	reply = fmt.Sprintf("%c %s", boolToEmoji(source.IsActive), source.Name)
	return &reply
}
