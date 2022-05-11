package commands

import (
	"AggreBot/internal/pkg/api"
	"strconv"
)

func boolToEmoji(value bool) rune {
	if value == true {
		return '✅'
	} else {
		return '☑'
	}
}

func (m *Manager) getSourceFromUserArg(userId int64, args []string) (*api.Source, string) {
	sources, err := m.backend.GetUserSources(userId)
	if err != nil {
		return nil, errInternalError
	}

	if len(sources) == 0 {
		return nil, errNoAnySources
	}

	if len(args) == 0 {
		return nil, errNoSourceIndex
	}

	indexString := args[0]
	index, err := strconv.Atoi(indexString)
	if err != nil {
		return nil, errWrongSourceIndex
	}

	if index < 1 || len(sources) < index {
		return nil, errNoSuchSource
	}

	return sources[index-1], ""
}
