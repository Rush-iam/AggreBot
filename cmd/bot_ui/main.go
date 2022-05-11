package main

import (
	"AggreBot/internal/bot_ui"
	"AggreBot/internal/pkg/exit_signal"
	"AggreBot/internal/pkg/grpc_client"
	"AggreBot/internal/pkg/tg_client"
	"context"
)

const grpcServerEndpoint = "localhost:8080"
const tgToken = "5336663940:AAHwU2dP2TLSVde7EYLeVVJAsr5goVuVkz4"

func main() {
	grpcClient := grpc_client.New(context.Background(), grpcServerEndpoint)
	defer grpcClient.Close()
	tgClient := tg_client.New(tgToken)

	bot := bot_ui.NewBot(tgClient, grpcClient)
	go bot.RunBotLoop()

	<-exit_signal.Wait()
}
