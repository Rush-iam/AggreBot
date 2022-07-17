package callbacks

import (
	"fmt"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/command"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/errors"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/markup"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func cbMenuReplyText() string {
	return "🤖"
}

func cbMenuReplyButtons(sourcesCount int, userFilter string) [][]tgbotapi.InlineKeyboardButton {
	var filterText, filterData string
	if userFilter == "" {
		filterText = "🔍 Set Filter (not set)"
		filterData = "filter"
	} else {
		filterText = fmt.Sprintf("🔍 Filter: \"%s\"", userFilter)
		filterData = "filter_menu"
	}
	var buttons [][]tgbotapi.InlineKeyboardButton
	if sourcesCount > 0 {
		buttons = append(buttons, markup.ButtonRow(markup.Button(fmt.Sprintf("📝 My Sources (%d)", sourcesCount), "list")))
	}
	buttons = append(buttons, markup.ButtonRow(markup.Button("➕ Add new Source", "source_add")))
	buttons = append(buttons, markup.ButtonRow(markup.Button(filterText, filterData)))
	return buttons
}

func (m *Manager) cbMenu(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	var userFilter string
	user, err := m.backend.GetUser(c.UserId)
	if err != nil {
		if status.Code(err) != codes.NotFound || m.registerUser(c.UserId) == false {
			return errors.ErrInternalError, nil
		}
	} else {
		userFilter = user.Filter
	}

	sources, err := m.backend.GetUserSources(c.UserId)
	if err != nil {
		return errors.ErrInternalError, nil
	}

	reply := cbMenuReplyText()
	keyboard := markup.Keyboard(cbMenuReplyButtons(len(sources), userFilter))
	return reply, &keyboard
}
