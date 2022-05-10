package commands

import (
	"AggreBot/api"
	"AggreBot/internal/bot_ui/grpc_client"
	"context"
	"fmt"
	"regexp"
)

func cmdFilter(c Command) *string {
	var reply string
	var userFilter string
	if len(c.args) == 0 {
		userFilter = ""
	} else {
		userFilter = c.args[0]
	}
	if len([]rune(userFilter)) > 256 {
		reply = "🤯 Oh! Your Filter is too looong"
		return &reply
	}

	_, err := regexp.Compile(userFilter)
	if err != nil {
		reply = "🤒 I had troubles compiling that RegExp Filter, sorry"
		return &reply
	}

	_, err = grpc_client.Cl.UpdateUserFilter(
		context.Background(),
		&api.User{
			Id:     c.userId,
			Filter: userFilter,
		},
	)
	if err != nil {
		reply = "⚠ Oops. Internal Error. Please try again later."
		return &reply
	}
	if userFilter == "" {
		reply = "🔍 No Filter set"
	} else {
		reply = fmt.Sprintf("🔍 RegExp Filter: '%s'", userFilter)
	}
	return &reply
}
