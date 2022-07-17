package callbacks

import (
	"fmt"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/command"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/markup"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func cbSourceMenuReplyText(sourceName string, isActive bool, url string) string {
	return fmt.Sprintf("%s\n%s", markup.SourceString(sourceName, isActive), url)
}

func cbSourceMenuReplyButtons(isActive bool, sourceId int64) [][]tgbotapi.InlineKeyboardButton {
	var toggleText, toggleData string
	if isActive {
		toggleText = fmt.Sprintf("%c Disable", markup.BoolToEmoji(!isActive))
		toggleData = fmt.Sprintf("source_active_disable %d", sourceId)
	} else {
		toggleText = fmt.Sprintf("%c Enable", markup.BoolToEmoji(!isActive))
		toggleData = fmt.Sprintf("source_active_enable %d", sourceId)
	}
	buttons := [][]tgbotapi.InlineKeyboardButton{
		markup.ButtonRow(markup.Button("✏ Rename", fmt.Sprintf("source_rename %d", sourceId))),
		markup.ButtonRow(markup.Button(toggleText, toggleData)),
		markup.ButtonRow(markup.Button("🗑 Remove", fmt.Sprintf("source_remove %d", sourceId))),
		markup.ButtonBackToList(),
	}
	return buttons
}

func (m *Manager) cbSourceMenu(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	source, errReply := m.getSourceIdFromArg(c.UserId, c.Text)
	if errReply != "" {
		keyboard := markup.KeyboardBackToMenu()
		return errReply, &keyboard
	}

	reply := cbSourceMenuReplyText(source.Name, source.IsActive, source.Url)
	keyboard := markup.Keyboard(cbSourceMenuReplyButtons(source.IsActive, source.Id))
	return reply, &keyboard
}
