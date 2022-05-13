package tg_client

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

func New(token string) *tgbotapi.BotAPI {
	tgClient, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Can't get TG client: %v", err)
	} else {
		log.Printf("Connected to Telegram API as @%s", tgClient.Self.UserName)
	}
	return tgClient
}

const sendRpsLimit = 30
const sendSleepPeriod = 1000 / sendRpsLimit

var unlockTime time.Time

func SendMessage(tgClient *tgbotapi.BotAPI, userId int64, text string) error {
	if text == "" {
		return nil
	}

	curTime := time.Now()
	time.Sleep(unlockTime.Sub(curTime))
	unlockTime = curTime.Add(time.Millisecond * sendSleepPeriod)

	reply := tgbotapi.NewMessage(userId, text)
	_, err := tgClient.Send(reply)
	return err
}
