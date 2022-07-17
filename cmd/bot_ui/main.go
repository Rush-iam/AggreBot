package main

import (
	"context"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/config"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/exit_signal"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/grpc_client"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/pkg/tg_client"
	"log"
	"time"
)

var flags = map[string]string{
	"grpchost": "gRPC Server Host",
	"tgtoken":  "Telegram Bot Token",
}

func main() {
	cfg := config.FromFlags(flags)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	grpcClient := grpc_client.New(ctx, cfg["grpchost"])
	cancel()

	tgClient := tg_client.New(cfg["tgtoken"])

	bot := bot_ui.NewBot(tgClient, grpcClient)
	go bot.RunBotLoop()

	<-exit_signal.Wait()
	grpcClient.Close()
	log.Printf("gRPC client shutted down")
}
