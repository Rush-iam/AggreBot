
PROTOGEN = --go_out . --go-grpc_out . --grpc-gateway_out . --openapiv2_out ./api

protogen:
	protoc -I./api $(PROTOGEN) ./api/api.proto

run_backend:
	go run ./cmd/backend

run_bot_ui:
	go run ./cmd/bot_ui

run_courier:
	go run ./cmd/courier