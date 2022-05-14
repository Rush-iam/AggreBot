package states

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/bot_ui/markup"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func stRenameReplyText(sourceToRenameIsActive bool, newName string) string {
	return fmt.Sprintf("%c New name: %s", markup.BoolToEmoji(sourceToRenameIsActive), newName)
}

func (m *Manager) stRenameReply(c *command.Command, sourceId int64) (string, bool) {
	sourceToRename, err := m.backend.GetSource(sourceId)
	if err != nil || sourceToRename.UserId != c.UserId {
		return errors.ErrInternalError, false
	}

	newName := c.Cmd
	err = m.backend.UpdateSourceName(sourceToRename.Id, newName)
	if err != nil {
		return errors.ErrInternalError, false
	}

	return stRenameReplyText(sourceToRename.IsActive, newName), true
}

func (m *Manager) stRename(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	sourceId := c.UserState.Value
	reply, ok := m.stRenameReply(c, sourceId)

	var keyboard tgbotapi.InlineKeyboardMarkup
	if !ok {
		keyboard = markup.KeyboardBackToMenu()
	} else {
		keyboard = tgbotapi.NewInlineKeyboardMarkup(markup.ButtonBackToSource(sourceId))
	}
	return reply, &keyboard
}
