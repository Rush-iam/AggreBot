package callbacks

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/bot_ui/markup"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func cbFilterMenuReplyText(userFilter string) string {
	if userFilter == "" {
		return "🔍 No Filter set"
	} else {
		return fmt.Sprintf("🔍 Filter: \"%s\"", userFilter)
	}
}

func cbFilterMenuReplyButtons() [][]tgbotapi.InlineKeyboardButton {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		markup.ButtonRow(markup.Button("✏ Edit", "filter")),
		markup.ButtonRow(markup.Button("🗑 Remove", "filter_remove")),
		markup.ButtonBackToMenu(),
	}
	return buttons
}

func (m *Manager) cbFilterMenu(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	userFilter, err := m.backend.GetUserFilter(c.UserId)
	if err != nil {
		keyboard := markup.KeyboardBackToMenu()
		return errors.ErrInternalError, &keyboard
	}

	reply := cbFilterMenuReplyText(*userFilter)
	keyboard := markup.Keyboard(cbFilterMenuReplyButtons())
	return reply, &keyboard
}
