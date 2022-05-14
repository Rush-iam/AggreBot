package tg_client

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync"
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

var threadLock sync.Mutex
var unlockTime time.Time

func SendMessage(tgClient *tgbotapi.BotAPI, userId int64, text string) error {
	if text == "" {
		return nil
	}

	threadLock.Lock()
	time.Sleep(unlockTime.Sub(time.Now()))
	unlockTime = time.Now().Add(time.Millisecond * sendSleepPeriod)
	threadLock.Unlock()

	reply := tgbotapi.NewMessage(userId, text)
	_, err := tgClient.Send(reply)
	return err
}
