package bot_ui

import (
	"AggreBot/internal/bot_ui/commands"
	"AggreBot/internal/pkg/grpc_client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	tgClient       *tgbotapi.BotAPI
	commandManager *commands.Manager
}

func NewBot(tgClient *tgbotapi.BotAPI, grpcClient *grpc_client.Client) *Bot {
	return &Bot{
		tgClient:       tgClient,
		commandManager: commands.NewManager(grpcClient),
	}
}

func (bot *Bot) RunBotLoop() {
	uConfig := tgbotapi.NewUpdate(0)
	uConfig.Timeout = 60
	for {
		for u := range bot.tgClient.GetUpdatesChan(uConfig) {
			if u.Message != nil && u.Message.Chat.IsPrivate() {
				replyText := bot.handleMessage(u.Message)
				if replyText != nil {
					bot.sendMessage(u.Message.From.ID, replyText)
				}
			}
		}
	}
}

func (bot *Bot) handleMessage(msg *tgbotapi.Message) *string {
	log.Printf("[%s] %s", msg.From.UserName, msg.Text)

	if len(msg.Entities) > 0 {
		command := commands.ParseFromMessage(msg)
		if command != nil {
			return bot.commandManager.Execute(command)
		}
	}
	help := "👇 Use commands from Menu"
	return &help
}

func (bot *Bot) sendMessage(userId int64, text *string) {
	reply := tgbotapi.NewMessage(userId, *text)
	_, err := bot.tgClient.Send(reply)
	if err != nil {
		log.Print(err)
	}
	log.Printf("[bot] %s", *text)
}
