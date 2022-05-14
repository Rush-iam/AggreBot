package command

import (
	"AggreBot/internal/bot_ui/user_state"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type Command struct {
	UserId    int64
	Cmd       string
	Args      []string
	UserState *user_state.UserState
}

func FromValues(userId int64, cmd string, userState *user_state.UserState) *Command {
	return &Command{
		UserId:    userId,
		Cmd:       cmd,
		UserState: userState,
	}
}

func FromMessage(msg *tgbotapi.Message, userState *user_state.UserState) *Command {
	for _, entity := range msg.Entities {
		if entity.IsCommand() && entity.Offset == 0 {
			command := []rune(msg.Text)[:entity.Length]
			return &Command{
				UserId:    msg.From.ID,
				Cmd:       string(command),
				Args:      strings.Fields(msg.Text)[1:],
				UserState: userState,
			}
		}
	}
	return nil
}

func FromCallbackQuery(query *tgbotapi.CallbackQuery, userState *user_state.UserState) *Command {
	commandWithArgs := strings.Fields(query.Data)
	return &Command{
		UserId:    query.From.ID,
		Cmd:       commandWithArgs[0],
		Args:      commandWithArgs[1:],
		UserState: userState,
	}
}
