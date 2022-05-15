package states

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/bot_ui/markup"
	"AggreBot/internal/pkg/rss_feed"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func stAddReplyText(sourceName string) string {
	return fmt.Sprintf("New Source:\n%s", markup.SourceString(sourceName, true))
}

func (m *Manager) stAddReply(c *command.Command) string {
	sourceUrl := c.Text
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

	return stAddReplyText(sourceName)
}

func (m *Manager) stAdd(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	reply := m.stAddReply(c)
	keyboard := markup.KeyboardBackToMenu()
	return reply, &keyboard
}
