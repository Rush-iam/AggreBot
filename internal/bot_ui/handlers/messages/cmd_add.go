package messages

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/pkg/rss_feed"
	"fmt"
	"strings"
)

func cmdAddReply(sourceName string) string {
	return fmt.Sprintf("✅ %s", sourceName)
}

func (m *Manager) cmdAdd(c *command.Command) string {
	words := strings.Fields(c.Text)
	if len(words) == 0 {
		return errors.ErrAddNoUrl
	}

	sourceUrl := words[1]
	if len([]rune(sourceUrl)) > 2048 {
		return errors.ErrAddUrlTooLong
	}

	sourceName, ok := rss_feed.GetTitle(sourceUrl)
	if !ok {
		return errors.ErrAddRssParseError
	}

	err := m.backend.AddSource(c.UserId, sourceName, sourceUrl)
	if err != nil {
		return errors.ErrInternalError
	}

	return cmdAddReply(sourceName)
}
