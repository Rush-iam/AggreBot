package callbacks

import (
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/command"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/user_state"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func cbSourceAddReplyText() string {
	return "👇 Enter RSS/Atom Feed URL:"
}

func (m *Manager) cbSourceAdd(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	c.UserState.State = user_state.SourceAdd
	return cbSourceAddReplyText(), nil
}
