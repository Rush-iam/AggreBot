package tg_client

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var Cl *tgbotapi.BotAPI

func Init(token string) {
	Cl = getTGClient(token)
}

func getTGClient(token string) *tgbotapi.BotAPI {
	tgClient, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Can't get TG client: %v", err)
	} else {
		log.Printf("Connected to Telegram API as @%s", tgClient.Self.UserName)
	}
	return tgClient
}
