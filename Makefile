#     _                        ____        _
#    / \   __ _  __ _ _ __ ___| __ )  ___ | |_
#   / _ \ / _` |/ _` | '__/ _ \  _ \ / _ \| __|
#  / ___ \ (_| | (_| | | |  __/ |_) | (_) | |_
# /_/   \_\__, |\__, |_|  \___|____/ \___/ \__|
#         |___/ |___/      by nGragas/Rush-iam

include .env_local

PROTOGEN = --go_out . --go-grpc_out . --grpc-gateway_out . --openapiv2_out ./api

GOOSE_DBSTRING="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

DB_ARG =-dbhost $(DB_HOST):$(DB_PORT)	\
		-dbname $(DB_NAME)				\
		-dbuser $(DB_USER)				\
		-dbpass $(DB_PASSWORD)

GRPC_HOST_ARG	= -grpchost $(BACKEND_HOST):$(GRPC_PORT)
REST_HOST_ARG	= -resthost $(BACKEND_HOST):$(REST_PORT)
TG_TOKEN_ARG	= -tgtoken $(TG_TOKEN)

protogen:
	protoc -I./api $(PROTOGEN) ./api/api.proto

goose:
	goose -dir ./migrations postgres $(GOOSE_DBSTRING) up

run_backend:
	go run ./cmd/backend $(DB_ARG) $(GRPC_HOST_ARG) $(REST_HOST_ARG)

run_bot_ui:
	go run ./cmd/bot_ui $(GRPC_HOST_ARG) $(TG_TOKEN_ARG)

run_courier:
	go run ./cmd/courier $(DB_ARG) $(TG_TOKEN_ARG)