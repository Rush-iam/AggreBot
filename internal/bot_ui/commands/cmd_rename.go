package commands

import (
	"AggreBot/api"
	"AggreBot/internal/bot_ui/grpc_client"
	"context"
	"fmt"
	"strings"
)

func cmdRename(c Command) *string {
	var reply string
	sourceToRename, errReply := fetchSourceFromUserArg(c.userId, c.args)
	if errReply != nil {
		return errReply
	}
	if len(c.args) == 1 {
		reply = "👉 Hey! You forgot to provide new name"
		return &reply
	}
	newName := strings.Join(c.args[1:], " ")
	_, err := grpc_client.Cl.UpdateSourceName(
		context.Background(),
		&api.UpdateSourceNameRequest{
			Id:   sourceToRename.Id,
			Name: newName,
		},
	)
	if err != nil {
		reply = "⚠ Oops. Internal Error. Please try again later."
		return &reply
	}
	reply = fmt.Sprintf("%c %s", boolToEmoji(sourceToRename.IsActive), newName)
	return &reply
}
