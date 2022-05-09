package commands

import (
	"AggreBot/api"
	"AggreBot/internal/tg_frontend/grpc_client"
	"context"
	"fmt"
	"strconv"
)

func delete(c Command) *string {
	var reply string
	sources, ok := fetchUserSources(c.userId)
	if !ok {
		reply = "⚠ Oops. Internal Error. Please try again later."
		return &reply
	}
	if len(sources) == 0 {
		reply = "🤷 There is no any Sources"
		return &reply
	}
	if len(c.args) == 0 {
		reply = "👉 Hey! You forgot # of Source"
		return &reply
	}
	indexString := c.args[0]
	index, err := strconv.Atoi(indexString)
	if err != nil {
		reply = fmt.Sprintf(
			"👉 Hey! Index of Source should be a number, not '%s'", indexString,
		)
		return &reply
	}
	if index < 1 || len(sources) < index {
		reply = fmt.Sprintf("👉 Hey! There is no Source with #%d", index)
		return &reply
	}
	sourceToDelete := sources[index-1]
	_, err = grpc_client.Cl.DeleteSource(
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
