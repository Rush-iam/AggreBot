package main

import (
	"AggreBot/api"
	"AggreBot/internal/tg_frontend"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const grpcServerEndpoint = "localhost:8080"
const tgToken = "5336663940:AAHwU2dP2TLSVde7EYLeVVJAsr5goVuVkz4"

var lastProcessedUpdate int

func main() {
	grpcConnection := connectGRPC(grpcServerEndpoint)
	defer closeGRPC(grpcConnection)
	grpcClient := api.NewNewsfeedConfiguratorClient(grpcConnection)
	tgClient := getTGClient(tgToken)

	tg_frontend.Run(&grpcClient, tgClient, lastProcessedUpdate)
}

func connectGRPC(grpcServerEndpoint string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	grpcConnection, err := grpc.DialContext(
		ctx,
		grpcServerEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Can't connect to gRPC server: %v", err)
	} else {
		log.Printf("Connected to gRPC server: %s", grpcServerEndpoint)
	}
	return grpcConnection
}

func closeGRPC(grpcConnection *grpc.ClientConn) {
	_ = grpcConnection.Close()
}

func getTGClient(token string) *tgbotapi.BotAPI {
	tgClient, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Can't get TG client: %v", err)
	} else {
		log.Printf("Connected to Telegram API as %s", tgClient.Self.UserName)
	}
	return tgClient
}
