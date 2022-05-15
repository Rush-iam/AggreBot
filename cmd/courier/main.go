package main

import (
	"AggreBot/internal/courier"
	"AggreBot/internal/pkg/config"
	"AggreBot/internal/pkg/db_client"
	"AggreBot/internal/pkg/exit_signal"
	"context"
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

	db := db_client.NewClient(
		context.Background(), cfg["dbuser"], cfg["dbpass"], cfg["dbhost"], cfg["dbname"],
	)
	defer db.Close()

	worker := courier.NewCourier(db, cfg["tgtoken"])
	go worker.RunReader()
	go worker.RunSender()

	<-exit_signal.Wait()
}
