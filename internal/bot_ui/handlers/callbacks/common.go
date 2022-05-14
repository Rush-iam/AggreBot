package callbacks

import (
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/pkg/api"
	"strconv"
)

func (m *Manager) getSourceFromArg(userId int64, args []string) (*api.Source, string) {
	if len(args) == 0 {
		return nil, errors.ErrInternalError
	}

	sourceId, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.ErrInternalError
	}

	source, err := m.backend.GetSource(int64(sourceId))
	if err != nil || source.UserId != userId {
		return nil, errors.ErrNoSuchSource
	}

	return source, ""
}
