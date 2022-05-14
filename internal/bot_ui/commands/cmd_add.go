package commands

import (
	"AggreBot/internal/pkg/rss_feed"
	"fmt"
)

func cmdAddReply(sourceName string) string {
	return fmt.Sprintf("✅ %s", sourceName)
}

func (m *Manager) cmdAdd(c *command) string {
	if len(c.args) == 0 {
		return errAddNoUrl
	}

	sourceUrl := c.args[0]
	if len([]rune(sourceUrl)) > 2048 {
		return errAddUrlTooLong
	}

	sourceName, ok := rss_feed.GetTitle(sourceUrl)
	if !ok {
		return errAddRssParseError
	}

	err := m.backend.AddSource(c.userId, sourceName, sourceUrl)
	if err != nil {
		return errInternalError
	}

	return cmdAddReply(sourceName)
}
