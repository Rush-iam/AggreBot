package tg_frontend

import (
	"AggreBot/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Run(grpc *api.NewsfeedConfiguratorClient, tg *tgbotapi.BotAPI, uLast int) {
	uConfig := tgbotapi.NewUpdate(uLast)
	uConfig.Timeout = 60
	_ = grpc
	for {
		for u := range tg.GetUpdatesChan(uConfig) {
			if u.Message != nil {
				log.Printf("[%s] %s", u.Message.From.UserName, u.Message.Text)

				msg := tgbotapi.NewMessage(u.Message.Chat.ID, u.Message.Text)
				msg.ReplyToMessageID = u.Message.MessageID

				_, err := tg.Send(msg)
				if err != nil {
					log.Print(err)
				}
			} else {
				log.Print()
			}
		}
	}
}
