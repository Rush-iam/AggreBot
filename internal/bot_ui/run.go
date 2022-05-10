package bot_ui

import (
	"AggreBot/internal/bot_ui/commands"
	"AggreBot/internal/bot_ui/tg_client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func RunBotLoop() {
	uConfig := tgbotapi.NewUpdate(0)
	uConfig.Timeout = 60
	for {
		for u := range tg_client.Cl.GetUpdatesChan(uConfig) {
			if u.Message != nil && u.Message.Chat.IsPrivate() {
				replyText := handleMessage(u.Message)
				if replyText != nil {
					sendMessage(tg_client.Cl, u.Message.From.ID, replyText)
				}
			}
		}
	}
}

func handleMessage(msg *tgbotapi.Message) *string {
	log.Printf("[%s] %s", msg.From.UserName, msg.Text)

	if len(msg.Entities) > 0 {
		cmd := commands.ParseFromMessage(msg)
		if cmd != nil {
			return cmd.Execute()
		}
	}
	help := "👇 Use commands from Menu"
	return &help
}

func sendMessage(tg *tgbotapi.BotAPI, userId int64, text *string) {
	reply := tgbotapi.NewMessage(userId, *text)
	_, err := tg.Send(reply)
	if err != nil {
		log.Print(err)
	}
	log.Printf("[bot] %s", *text)
}
