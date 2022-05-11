package courier

import (
	"AggreBot/internal/pkg/db_client"
	"AggreBot/internal/pkg/safe_set"
	"AggreBot/internal/pkg/tg_client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mmcdole/gofeed"
)

type courier struct {
	db            *db_client.Client
	sourcesInWork *safe_set.Int64
	jobQueue      chan job
	tgClient      *tgbotapi.BotAPI
}

type job struct {
	ok      bool
	source  *db_client.CourierSource
	entries []*gofeed.Item
}

func NewCourier(db *db_client.Client, tgToken string) *courier {
	return &courier{
		db:            db,
		sourcesInWork: safe_set.NewInt64(),
		jobQueue:      make(chan job, 1),
		tgClient:      tg_client.New(tgToken),
	}
}
