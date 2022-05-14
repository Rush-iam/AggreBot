package callbacks

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/bot_ui/markup"
	"AggreBot/internal/pkg/api"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func cbListReplyText(sourcesLen int) string {
	if sourcesLen == 0 {
		return "🗒 No sources. Try to add!"
	} else {
		return "📝 Your sources:"
	}
}

func cbListReplyButtons(sources []*api.Source) [][]tgbotapi.InlineKeyboardButton {
	buttons := make([][]tgbotapi.InlineKeyboardButton, len(sources)+1)
	for i, source := range sources {
		text := fmt.Sprintf("%c %s", markup.BoolToEmoji(source.IsActive), source.Name)
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
