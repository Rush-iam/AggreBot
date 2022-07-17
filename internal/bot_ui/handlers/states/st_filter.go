package states

import (
	"fmt"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/command"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/errors"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/markup"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"regexp"
)

func stFilterReplyText(userFilter string) string {
	if userFilter == "" {
		return "🔍 No Filter set"
	} else {
		return fmt.Sprintf("🔍 Filter set: \"%s\"", userFilter)
	}
}

func (m *Manager) stFilterReply(c *command.Command) string {
	userFilter := c.Text
	if len([]rune(userFilter)) > 256 {
		return errors.ErrFilterTooLong
	}

	_, err := regexp.Compile(userFilter)
	if err != nil {
		return errors.ErrFilterRegExp
	}

	err = m.backend.UpdateUserFilter(c.UserId, userFilter)
	if err != nil {
		return errors.ErrInternalError
	}

	return stFilterReplyText(userFilter)
}

func (m *Manager) stFilter(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	reply := m.stFilterReply(c)
	keyboard := markup.KeyboardBackToMenu()
	return reply, &keyboard
}
