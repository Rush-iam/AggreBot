package messages

import (
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/pkg/api"
	"strconv"
)

func (m *Manager) getSourceFromUserArg(userId int64, args []string) (*api.Source, string) {
	sources, err := m.backend.GetUserSources(userId)
	if err != nil {
		return nil, errors.ErrInternalError
	}

	if len(sources) == 0 {
		return nil, errors.ErrNoAnySources
	}

	if len(args) == 0 {
		return nil, errors.ErrNoSourceIndex
	}

	indexString := args[0]
	index, err := strconv.Atoi(indexString)
	if err != nil {
		return nil, errors.ErrWrongSourceIndex
	}

	if index < 1 || len(sources) < index {
		return nil, errors.ErrNoSuchSource
	}

	return sources[index-1], ""
}
