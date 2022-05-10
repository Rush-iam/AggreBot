package commands

import (
	"AggreBot/api"
	"AggreBot/internal/bot_ui/grpc_client"
	"context"
	"fmt"
	"strconv"
)

func boolToEmoji(value bool) rune {
	if value == true {
		return '✅'
	} else {
		return '☑'
	}
}

func fetchSourceFromUserArg(userId int64, args []string) (*api.Source, *string) {
	var replyError string
	sources, ok := fetchUserSources(userId)
	if !ok {
		replyError = "⚠ Oops. Internal Error. Please try again later."
		return nil, &replyError
	}
	if len(sources) == 0 {
		replyError = "🤷 There is no any Sources"
		return nil, &replyError
	}
	if len(args) == 0 {
		replyError = "👉 Hey! You forgot # of Source"
		return nil, &replyError
	}
	indexString := args[0]
	index, err := strconv.Atoi(indexString)
	if err != nil {
		replyError = fmt.Sprintf(
			"👉 Hey! Index of Source should be a number, not '%s'", indexString,
		)
		return nil, &replyError
	}
	if index < 1 || len(sources) < index {
		replyError = fmt.Sprintf("👉 Hey! There is no Source with #%d", index)
		return nil, &replyError
	}
	return sources[index-1], nil
}

func fetchUserFilter(userId int64) (string, bool) {
	responseUser, err := grpc_client.Cl.GetUser(
		context.Background(),
		&api.UserId{Id: userId},
	)
	if err != nil {
		return "", false
	}
	return responseUser.Filter, true
}

func fetchUserSources(userId int64) ([]*api.Source, bool) {
	responseSources, err := grpc_client.Cl.GetUserSources(
		context.Background(),
		&api.UserId{Id: userId},
	)
	if err != nil {
		return nil, false
	}
	return responseSources.Sources, true
}
