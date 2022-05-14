package callbacks

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/bot_ui/markup"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func cbMenuReplyText() string {
	return "🤖 What do you want?"
}

func cbMenuReplyButtons(userFilter string) [][]tgbotapi.InlineKeyboardButton {
	var filterText, filterData string
	if userFilter == "" {
		filterText = "🔍 Set Filter (not set)"
		filterData = "filter"
	} else {
		filterText = fmt.Sprintf("🔍 Filter: \"%s\"", userFilter)
		filterData = "filter_menu"
	}

	buttons := [][]tgbotapi.InlineKeyboardButton{
		markup.ButtonRow(markup.Button("📝 My Source Feeds", "list")),
		markup.ButtonRow(markup.Button("➕ Add new Source", "add")),
		markup.ButtonRow(markup.Button(filterText, filterData)),
	}
	return buttons
}

func (m *Manager) cbMenu(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	userFilter, err := m.backend.GetUserFilter(c.UserId)
	if err != nil {
		return errors.ErrInternalError, nil
	}

	reply := cbMenuReplyText()
	keyboard := markup.Keyboard(cbMenuReplyButtons(*userFilter))
	return reply, &keyboard
}
