package main

import (
	"AggreBot/internal/tg_frontend"
	"AggreBot/internal/tg_frontend/grpc_client"
	"AggreBot/internal/tg_frontend/tg_client"
)

const grpcServerEndpoint = "localhost:8080"
const tgToken = "5336663940:AAHwU2dP2TLSVde7EYLeVVJAsr5goVuVkz4"

func main() {
	grpc_client.Init(grpcServerEndpoint)
	defer grpc_client.Close()

	tg_client.Init(tgToken)

	tg_frontend.RunBotLoop()
}
