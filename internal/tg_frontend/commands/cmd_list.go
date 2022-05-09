package commands

import (
	"AggreBot/api"
	"AggreBot/internal/tg_frontend/grpc_client"
	"context"
	"fmt"
	"strings"
)

func list(c Command) *string {
	var reply string
	userFilter, sources, ok := fetchUserFilterAndSources(c.userId)
	if !ok {
		reply = "⚠ Oops. Internal Error. Please try again later."
		return &reply
	}

	var replyLines []string
	if userFilter != "" {
		replyLines = append(
			replyLines,
			fmt.Sprintf("🔍 RegExp Filter: '%s'", userFilter),
		)
	}
	if len(sources) == 0 {
		replyLines = append(replyLines, "🗒 No sources. Try to add!")
	} else {
		replyLines = append(replyLines, "📝 Your sources:")
		for i, source := range sources {
			var isActive rune
			if source.IsActive {
				isActive = '✅'
			} else {
				isActive = '☑'
			}
			replyLines = append(
				replyLines,
				fmt.Sprintf("%c %d. %s", isActive, i+1, source.Name),
			)
		}
	}
	reply = strings.Join(replyLines, "\n")
	return &reply
}

func fetchUserFilterAndSources(userId int64) (string, []*api.Source, bool) {
	responseSources, err1 := grpc_client.Cl.GetUserSources(
		context.Background(),
		&api.UserId{Id: userId},
	)
	responseUser, err2 := grpc_client.Cl.GetUser(
		context.Background(),
		&api.UserId{Id: userId},
	)
	if err1 != nil || err2 != nil {
		return "", nil, false
	}
	return responseUser.Filter, responseSources.Sources, true
}
