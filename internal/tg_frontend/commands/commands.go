package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type Command struct {
	userId int64
	cmd    string
	args   []string
}

type commandHandler func(Command) *string

var commandsMap = map[string]commandHandler{
	"/start": start,
}

func (c Command) Execute() *string {
	cmdHandler, ok := commandsMap[c.cmd]
	if ok {
		return cmdHandler(c)
	} else {
		return nil
	}
}

func ParseFromMessage(msg *tgbotapi.Message) *Command {
	for _, entity := range msg.Entities {
		if entity.IsCommand() && entity.Offset == 0 {
			command := []rune(msg.Text)[:entity.Length]
			return &Command{
				userId: msg.From.ID,
				cmd:    string(command),
				args:   strings.Fields(msg.Text)[1:],
			}
		}
	}
	return nil
}
