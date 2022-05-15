package main

import (
	"AggreBot/internal/bot_ui"
	"AggreBot/internal/pkg/config"
	"AggreBot/internal/pkg/exit_signal"
	"AggreBot/internal/pkg/grpc_client"
	"AggreBot/internal/pkg/tg_client"
	"context"
)

var flags = map[string]string{
	"grpchost": "gRPC Server Host",
	"tgtoken":  "Telegram Bot Token",
}

func main() {
	cfg := config.FromFlags(flags)

	grpcClient := grpc_client.New(context.Background(), cfg["grpchost"])
	defer grpcClient.Close()
	tgClient := tg_client.New(cfg["tgtoken"])

	bot := bot_ui.NewBot(tgClient, grpcClient)
	go bot.RunBotLoop()

	<-exit_signal.Wait()
}
