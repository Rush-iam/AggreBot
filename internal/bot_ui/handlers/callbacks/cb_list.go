package callbacks

import (
	"fmt"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/command"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/errors"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/markup"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func cbListReplyText(sourcesLen int) string {
	if sourcesLen == 0 {
		return "🗒 Nothing here. Try to add!"
	} else {
		return "📝 Your sources:"
	}
}

func cbListReplyButtons(sources []*api.Source) [][]tgbotapi.InlineKeyboardButton {
	buttons := make([][]tgbotapi.InlineKeyboardButton, len(sources)+1)
	for i, source := range sources {
		text := markup.SourceString(source.Name, source.IsActive)
		data := fmt.Sprintf("source_menu %d", source.Id)
		buttons[i] = markup.ButtonRow(markup.Button(text, data))
	}
	buttons[len(sources)] = markup.ButtonBackToMenu()
	return buttons
}

func (m *Manager) cbList(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	sources, err := m.backend.GetUserSources(c.UserId)
	if err != nil {
		keyboard := markup.KeyboardBackToMenu()
		return errors.ErrInternalError, &keyboard
	}

	reply := cbListReplyText(len(sources))
	keyboard := markup.Keyboard(cbListReplyButtons(sources))
	return reply, &keyboard
}
