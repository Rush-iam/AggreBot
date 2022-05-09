package commands

import (
	"AggreBot/api"
	"AggreBot/internal/tg_frontend/grpc_client"
	"context"
	"fmt"
	"strconv"
	"strings"
)

func rename(c Command) *string {
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
	if len(c.args) == 1 {
		reply = "👉 Hey! You forgot to provide new name"
		return &reply
	}

	sourceToRename := sources[index-1]
	newName := strings.Join(c.args[1:], " ")
	_, err = grpc_client.Cl.UpdateSourceName(
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

	var isActive rune
	if sourceToRename.IsActive {
		isActive = '✅'
	} else {
		isActive = '☑'
	}
	reply = fmt.Sprintf("%c %d. %s", isActive, index, newName)
	return &reply
}
