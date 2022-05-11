package main

import (
	"AggreBot/internal/courier"
	"AggreBot/internal/pkg/db_client"
	"AggreBot/internal/pkg/exit_signal"
	"context"
)

const (
	dbUser     = "postgres"
	dbPassword = "j3qq4"
	dbHost     = "localhost"
	dbPort     = "5432"
	dbName     = "aggrebot"
)
const tgToken = "5336663940:AAHwU2dP2TLSVde7EYLeVVJAsr5goVuVkz4"

func main() {
	db := db_client.NewClient(
		context.Background(), dbUser, dbPassword, dbHost, dbPort, dbName,
	)
	defer db.Close()

	worker := courier.NewCourier(db, tgToken)
	go worker.RunReader()
	go worker.RunSender()

	<-exit_signal.Wait()
}
