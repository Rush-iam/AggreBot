package commands

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
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

	sourceName, ok := getRssTitle(sourceUrl)
	if !ok {
		return errAddRssParseError
	}

	err := m.backend.AddSource(c.userId, sourceName, sourceUrl)
	if err != nil {
		return errInternalError
	}

	return cmdAddReply(sourceName)
}

func getRssTitle(rawReference string) (string, bool) {
	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(rawReference)
	if err != nil {
		log.Printf("getRssTitle: %v", err)
		return "", false
	}
	return feed.Title, true
}
