package callbacks

import (
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/pkg/api"
	"strconv"
	"strings"
)

func (m *Manager) getSourceIdFromArg(userId int64, text string) (*api.Source, string) {
	words := strings.Fields(text)
	if len(words) <= 1 {
		return nil, errors.ErrInternalError
	}

	sourceIdStr := words[1]
	sourceId, err := strconv.Atoi(sourceIdStr)
	if err != nil {
		return nil, errors.ErrInternalError
	}

	source, err := m.backend.GetSource(int64(sourceId))
	if err != nil || source.UserId != userId {
		return nil, errors.ErrNoSuchSource
	}

	return source, ""
}
