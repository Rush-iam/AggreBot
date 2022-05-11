package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type command struct {
	userId  int64
	cmdName string
	args    []string
}

func ParseFromMessage(msg *tgbotapi.Message) *command {
	for _, entity := range msg.Entities {
		if entity.IsCommand() && entity.Offset == 0 {
			newCommand := []rune(msg.Text)[:entity.Length]
			return &command{
				userId:  msg.From.ID,
				cmdName: string(newCommand),
				args:    strings.Fields(msg.Text)[1:],
			}
		}
	}
	return nil
}
