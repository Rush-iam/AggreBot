include .env_local

PROTOGEN = --go_out . --go-grpc_out . --grpc-gateway_out . --openapiv2_out ./api

DB_ARG =-dbhost $(DB_HOST):$(DB_PORT)	\
		-dbname $(DB_NAME)				\
		-dbuser $(DB_USER)				\
		-dbpass $(DB_PASSWORD)

GRPC_HOST_ARG	= -grpchost $(BACKEND_HOST):$(GRPC_PORT)
REST_HOST_ARG	= -resthost $(BACKEND_HOST):$(REST_PORT)
TG_TOKEN_ARG	= -tgtoken $(TG_TOKEN)

protogen:
	protoc -I./api $(PROTOGEN) ./api/api.proto

run_backend:
	go run ./cmd/backend $(DB_ARG) $(GRPC_HOST_ARG) $(REST_HOST_ARG)

run_bot_ui:
	go run ./cmd/bot_ui $(GRPC_HOST_ARG) $(TG_TOKEN_ARG)

run_courier:
	go run ./cmd/courier $(DB_ARG) $(TG_TOKEN_ARG)