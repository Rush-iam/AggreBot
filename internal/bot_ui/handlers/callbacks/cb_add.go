package callbacks

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/user_state"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func cbAddReplyText() string {
	return "👇 Type RSS/Atom Feed URL:"
}

func (m *Manager) cbAdd(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	c.UserState.State = user_state.SourceAdd
	return cbAddReplyText(), nil
}
