
PROTOGEN = --go_out . --go-grpc_out . --grpc-gateway_out . --openapiv2_out ./api

protoc:
	protoc -I./api $(PROTOGEN) ./api/api.proto
