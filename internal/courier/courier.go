package courier

import (
	"AggreBot/internal/pkg/db_client"
	"AggreBot/internal/pkg/tg_client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mmcdole/gofeed"
	"sync"
)

type courier struct {
	db            *db_client.Client
	tg            *tgbotapi.BotAPI
	sourcesInWork sync.Map
	jobQueue      chan job
}

type job struct {
	source       *db_client.CourierSource
	entries      []*gofeed.Item
	wasReadError bool
}

func NewCourier(db *db_client.Client, tgToken string) *courier {
	return &courier{
		db:            db,
		tg:            tg_client.New(tgToken),
		sourcesInWork: sync.Map{},
		jobQueue:      make(chan job, 1),
	}
}
