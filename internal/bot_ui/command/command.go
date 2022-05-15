package command

import (
	"AggreBot/internal/bot_ui/user_state"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Command struct {
	UserId    int64
	Text      string
	UserState *user_state.UserState
}

func FromMessage(msg *tgbotapi.Message, userState *user_state.UserState) *Command {
	return &Command{
		UserId:    msg.From.ID,
		Text:      msg.Text,
		UserState: userState,
	}
}

func FromCallbackQuery(query *tgbotapi.CallbackQuery, userState *user_state.UserState) *Command {
	return &Command{
		UserId:    query.From.ID,
		Text:      query.Data,
		UserState: userState,
	}
}
