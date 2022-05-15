package callbacks

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/bot_ui/markup"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func cbFilterRemoveReplyText() string {
	return "🔍 Filter has been removed"
}

func (m *Manager) cbFilterRemove(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	keyboard := markup.KeyboardBackToMenu()
	err := m.backend.UpdateUserFilter(c.UserId, "")
	if err != nil {
		return errors.ErrInternalError, &keyboard
	}
	reply := cbFilterRemoveReplyText()
	return reply, &keyboard
}
