package callbacks

import (
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/command"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/errors"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/markup"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (m *Manager) cbSourceActive(c *command.Command, newIsActive bool) (string, *tgbotapi.InlineKeyboardMarkup) {
	source, errReply := m.getSourceIdFromArg(c.UserId, c.Text)
	if errReply != "" {
		keyboard := markup.KeyboardBackToMenu()
		return errReply, &keyboard
	}
	sourceNew, err := m.backend.UpdateSourceIsActive(source.Id, newIsActive)
	if err != nil {
		keyboard := markup.KeyboardBackToMenu()
		return errors.ErrInternalError, &keyboard
	}

	reply := cbSourceMenuReplyText(sourceNew.Name, sourceNew.IsActive, sourceNew.Url)
	keyboard := markup.Keyboard(cbSourceMenuReplyButtons(sourceNew.IsActive, source.Id))
	return reply, &keyboard
}

func (m *Manager) cbSourceActiveEnable(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	return m.cbSourceActive(c, true)
}

func (m *Manager) cbSourceActiveDisable(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	return m.cbSourceActive(c, false)
}
