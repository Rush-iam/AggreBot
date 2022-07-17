package states

import (
	"fmt"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/command"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/errors"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/markup"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func stRenameReplyText(sourceIsActive bool, sourceName string) string {
	return fmt.Sprintf("New name:\n%s", markup.SourceString(sourceName, sourceIsActive))
}

func (m *Manager) stRenameReply(c *command.Command, sourceId int64) (string, bool) {
	sourceToRename, err := m.backend.GetSource(sourceId)
	if err != nil || sourceToRename.UserId != c.UserId {
		if status.Code(err) == codes.NotFound {
			return errors.ErrNoSuchSource, false
		} else {
			return errors.ErrInternalError, false
		}
	}
	newName := c.Text
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
