package main

import (
	"context"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/courier"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/config"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/db_client"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/exit_signal"
	"log"
	"time"
)

var flags = map[string]string{
	"dbuser":  "Database Username",
	"dbpass":  "Database Password",
	"dbhost":  "Database Host",
	"dbname":  "Database Name",
	"tgtoken": "Telegram Bot Token",
}

func main() {
	cfg := config.FromFlags(flags)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	db, err := db_client.NewClient(
		ctx, cfg["dbuser"], cfg["dbpass"], cfg["dbhost"], cfg["dbname"],
	)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	worker := courier.NewCourier(db, cfg["tgtoken"])
	go worker.RunReader()
	go worker.RunSender()

	<-exit_signal.Wait()
}
