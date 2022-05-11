package commands

import (
	"AggreBot/internal/pkg/api"
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

func (m *Manager) getSourceFromUserArg(userId int64, args []string) (*api.Source, *string) {
	var replyError string
	sources, err := m.backend.GetUserSources(userId)
	if err != nil {
		replyError = "⚠ Oops. Internal Error. Please try again later."
		return nil, &replyError
	}
	if len(sources) == 0 {
		replyError = "🤷 You didn't add any Sources"
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
