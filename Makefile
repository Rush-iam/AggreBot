
PROTOGEN = --go_out . --go-grpc_out . --grpc-gateway_out . --openapiv2_out ./api

protogen:
	protoc -I./api $(PROTOGEN) ./api/api.proto

grpc:
	go run ./cmd/app

tg:
	go run ./cmd/tg_frontend