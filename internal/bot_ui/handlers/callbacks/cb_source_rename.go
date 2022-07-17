package callbacks

import (
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/command"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/markup"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/user_state"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func cbSourceRenameReplyText(isActive bool, sourceName, url string) string {
	prevReply := cbSourceMenuReplyText(sourceName, isActive, url)
	return prevReply + "\n\n👇 Type new name:"
}

func (m *Manager) cbSourceRename(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	source, errReply := m.getSourceIdFromArg(c.UserId, c.Text)
	if errReply != "" {
		keyboard := markup.KeyboardBackToMenu()
		return errReply, &keyboard
	}

	c.UserState.State = user_state.SourceRename
	c.UserState.Value = source.Id
	reply := cbSourceRenameReplyText(source.IsActive, source.Name, source.Url)
	return reply, nil
}
