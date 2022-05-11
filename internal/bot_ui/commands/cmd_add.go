package commands

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
)

func (m *Manager) cmdAdd(c *command) *string {
	var reply string
	if len(c.args) == 0 {
		reply = "👉 Hey! You forgot source URL"
		return &reply
	}
	sourceUrl := c.args[0]
	if len([]rune(sourceUrl)) > 2048 {
		reply = "🤯 Oh! Your URL is too looong"
		return &reply
	}
	sourceName, ok := getRssTitle(sourceUrl)
	if !ok {
		reply = "🤒 I had troubles parsing RSS from that URL, sorry"
		return &reply
	}

	err := m.backend.AddSource(c.userId, sourceName, sourceUrl)
	if err != nil {
		reply = "⚠ Oops. Internal Error. Please try again later."
	} else {
		reply = fmt.Sprintf("✅ %s", sourceName)
	}
	return &reply
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
